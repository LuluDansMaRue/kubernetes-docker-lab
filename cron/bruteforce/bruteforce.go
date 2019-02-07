// Bruteforce.go is a CRON task
// This cron file will try to connect to the database with wrong credentials for a random number of access
// The goal is to see if there's any significient pick catch by the monitoring tool
package main

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	driver string = "mysql"
	min    int    = 100
	max    int    = 300
)

// Build DB Endpoint
// Param string username
// Param string password
// Return string
func buildDBEndpoint(user string, password string) string {
	username := user
	pwd := password
	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DATABASE")

	return username + ":" + pwd + "@tcp(" + host + ")/" + db
}

// GetCon
// Get the connection instance to the database
// Return database *sql.DB
func getCon() *sql.DB {
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, true, true, 32)

	endpoint := buildDBEndpoint(username, password)
	db, err := sql.Open(driver, endpoint)

	if err != nil {
		panic(err)
	}

	return db
}

// GetRandInt
// Generate a random number from 1 to 15
// Return int
func getRandInt() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

// Main function
func main() {
	rand := getRandInt()
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
