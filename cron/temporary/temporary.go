// Temporary.go is a CRON task
// This cron file will create a random number of Temporary table
// The goal is to see if there's any significient by a monitoring tool
package main

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	driver string = "mysql"
)

// Build DB Endpoint
// Return string
func buildDBEndpoint() string {
	username := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
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

// GetRandInt
// Generate a random number from 1 to 15
// Return int
func getRandInt() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 15

	return rand.Intn(max-min) + min
}

// Main function
func main() {
	rand := getRandInt()
	idx := 0
	log.Printf("Start CRON task. Creating " + strconv.Itoa(rand) + " temporary table")

	db := getCon()

	var wg sync.WaitGroup
	wg.Add(rand)

	for idx <= rand {
		go func(idx int) {
			stmt, err := db.Prepare("CREATE TEMPORARY TABLE ? SELECT * FROM bobba LIMIT 0")
			if err != nil {
				log.Fatal(err.Error())
			}

			_, execErr := stmt.Exec(
				"bobba_temp" + strconv.Itoa(idx),
			)

			if execErr != nil {
				log.Fatal(execErr.Error())
				wg.Done()
			}

			log.Printf("Finishing creating temporary table bobba_temp" + strconv.Itoa(idx))
			wg.Done()
		}(idx)

		idx++
	}

	wg.Wait()

	log.Printf("Finishing creating temporary table. Connection will close & clean the temporary table")
	err := db.Close()

	if err != nil {
		log.Fatal("Unable to close the Database " + err.Error())
	}

	return
}
