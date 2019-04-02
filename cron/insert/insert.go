package main

import (
	"log"
	"sync"

	"github.com/brianvoe/gofakeit"

	"github.com/LuluDansMaRue/kubernetes-docker-lab-cron/utils"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	db := utils.GetCon(false)

	// Launch a goroutine (this is a closure that is execute asynchronously in a seperate light thread)
	go func() {
		stmt, err := db.Prepare("INSERT INTO bobba SET name = ?, price = ?, shop = ?, flavor = ?, calory = ?")

		if err != nil {
			log.Fatal(err.Error())
			wg.Done()
			return
		}

		_, execErr := stmt.Exec(
			gofakeit.BeerName(),
			10.0,
			gofakeit.Name(),
			gofakeit.NameSuffix(),
			100.0,
		)

		if execErr != nil {
			log.Fatal(execErr.Error())
			wg.Done()
			return
		}

		wg.Done()
		return
	}()

	wg.Wait()

	log.Printf("Finishing inserting a bobba")
	err := db.Close()

	if err != nil {
		log.Fatal("Unable to close the Database " + err.Error())
	}
}
