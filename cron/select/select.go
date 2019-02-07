// Select.go CRON task
// This cron file will try to do a random select on the database
// The goal is to see if there's any pick on the database monitoring system
package main

import (
	"log"
	"strconv"
	"sync"

	"../utils"
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

	log.Printf("Start CRON Task. Making " + strconv.Itoa(rand) + " Select on bobba database")
	db := utils.GetCon(false)

	for idx <= rand {
		go func(id int) {
			stmt, err := db.Prepare("SELECT * FROM bobba WHERE ID = ?")
			if err != nil {
				log.Fatal(err.Error())
				wg.Done()
				return
			}

			_, execErr := stmt.Exec(idx)
			if execErr != nil {
				log.Fatal("Unable to execute the query. " + execErr.Error())
				wg.Done()
				return
			}

			log.Println("Finishing testing select for ID " + strconv.Itoa(id))

			wg.Done()
			return
		}(idx)

		idx++
	}

	wg.Wait()

	log.Printf("Finishing to test the select connection")

	err := db.Close()
	if err != nil {
		log.Fatal("Unable to close the Database " + err.Error())
	}
}
