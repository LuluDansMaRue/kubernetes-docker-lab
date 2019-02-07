// Bruteforce.go is a CRON task
// This cron file will try to connect to the database with wrong credentials for a random number of access
// The goal is to see if there's any significient pick catch by the monitoring tool
package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"sync"

	"../utils"
	"github.com/brianvoe/gofakeit"
)

const (
	driver string = "mysql"
)

// Build DB Endpoint
// Return string
func buildDBEndpoint() string {
	username := gofakeit.Username()
	pwd := gofakeit.Password(true, true, true, true, true, 32)

	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DATABASE")

	return username + ":" + pwd + "@tcp(" + host + ")/" + db
}

// GetCon
// Get the connection instance to the database
// Return database *sql.DB
func getCon() *sql.DB {
	endpoint := buildDBEndpoint()
	db, err := sql.Open(driver, endpoint)

	if err != nil {
		panic(err)
	}

	return db
}

// Main function
func main() {
	rand := utils.GetRandInt(100, 300)
	idx := 0
	var wg sync.WaitGroup
	wg.Add(rand)

	log.Printf("Start CRON task. Testing " + strconv.Itoa(rand) + " connection to the database")

	for idx <= rand {
		go func(idx int) {
			db := getCon()

			err := db.Ping()

			if err == nil {
				log.Fatal("User has been able to connect with wrong credentials")
				wg.Done()
			}

			log.Printf("Connection testing" + strconv.Itoa(idx) + " pass")
			wg.Done()
		}(idx)

		idx++
	}

	wg.Wait()
	log.Printf("Finishing bruteforce connection testing")
}
