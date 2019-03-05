// This cron file will try to connect to the database with wrong credentials for a random number of access
// The goal is to see if there's any significient pick catch by the monitoring tool
package main

import (
	"log"
	"strconv"
	"sync"

	"github.com/LuluDansMaRue/kubernetes-docker-lab-cron/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Main function
func main() {
	// set the logger output
	f := utils.SetLogOutput()
	defer f.Close()
	log.SetOutput(f)

	rand := utils.GetRandInt(10, 30)
	idx := 0

	// Create a waiting group that will defer the execution of the rest of the code
	var wg sync.WaitGroup
	wg.Add(rand)

	log.Printf("Start CRON task. Testing " + strconv.Itoa(rand) + " connection to the database")

	// Execute a goroutine foreach index
	for idx <= rand {
		go func(idx int) {
			db := utils.GetCon(true)

			// We ping the database in order to know if we're able to connect
			// Note: the ping will appear as an administrator command on the metrics
			err := db.Ping()

			if err == nil {
				log.Fatal("User has been able to connect with wrong credentials")
				// Close db (not handling error)
				db.Close()
				wg.Done()
			}

			log.Printf("Connection testing " + strconv.Itoa(idx) + " pass")
			// Close db (not handling error)
			db.Close()

			wg.Done()
		}(idx)

		idx++
	}

	wg.Wait()
	log.Printf("Finishing bruteforce connection testing")
}
