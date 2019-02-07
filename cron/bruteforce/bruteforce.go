// Bruteforce.go is a CRON task
// This cron file will try to connect to the database with wrong credentials for a random number of access
// The goal is to see if there's any significient pick catch by the monitoring tool
package main

import (
	"log"
	"strconv"
	"sync"

	"../utils"
	_ "github.com/go-sql-driver/mysql"
)

// Main function
func main() {
	f := utils.SetLogOutput()
	defer f.Close()
	log.SetOutput(f)

	rand := utils.GetRandInt(10, 30)
	idx := 0
	var wg sync.WaitGroup
	wg.Add(rand)

	log.Printf("Start CRON task. Testing " + strconv.Itoa(rand) + " connection to the database")

	for idx <= rand {
		go func(idx int) {
			db := utils.GetCon(true)

			err := db.Ping()

			if err == nil {
				log.Fatal("User has been able to connect with wrong credentials")
				// Close db (not handling error)
				db.Close()
				wg.Done()
			}

			log.Printf("Connection testing" + strconv.Itoa(idx) + " pass")
			// Close db (not handling error)
			db.Close()

			wg.Done()
		}(idx)

		idx++
	}

	wg.Wait()
	log.Printf("Finishing bruteforce connection testing")
}
