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
	rand := utils.GetRandInt(1, 20)
	idx := 0
	log.Printf("Start CRON task. Creating " + strconv.Itoa(rand) + " temporary table")

	db := utils.GetCon()

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
