// Temporary.go is a CRON task
// This cron file will create a random number of Temporary table
// The goal is to see if there's any significient by a monitoring tool
package main

import (
	"log"
	"strconv"
	"sync"

	"../utils"
)

// Main function
func main() {
	// Set the log output
	// f := utils.SetLogOutput()
	// defer f.Close()
	// log.SetOutput(f)

	// Choose the number of iteration that this teste will run
	rand := utils.GetRandInt(1, 30)
	idx := 0
	log.Printf("Start CRON task. Creating " + strconv.Itoa(rand) + " temporary table")

	db := utils.GetCon(false)

	var wg sync.WaitGroup
	wg.Add(rand)

	// For each indexes we will run a goroutine (light thread) that will execute our SQL queries
	// We use a waitGroup in order for our program to wait for this light thread to finish
	for idx <= rand {
		go func(idx int) {
			_, err := db.Exec("CREATE TEMPORARY TABLE bobba_temp_" + strconv.Itoa(idx) + " SELECT * FROM bobba LIMIT 0")
			if err != nil {
				log.Fatal(err.Error())
			}

			log.Printf("Finishing creating temporary table bobba_temp" + strconv.Itoa(idx))
			wg.Done()
		}(idx)

		idx++
	}

	wg.Wait()

	log.Printf("Finishing creating temporary table. Connection will close & clean the temporary table")
	// Close the db connection
	err := db.Close()

	if err != nil {
		log.Fatal("Unable to close the Database " + err.Error())
	}

	return
}
