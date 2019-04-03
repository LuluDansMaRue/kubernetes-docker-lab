package main

import (
	"sync"

	"github.com/LuluDansMaRue/kubernetes-docker-lab/bobba"
	"github.com/LuluDansMaRue/kubernetes-docker-lab/database"
)

// Run Add Routine
// Run a goroutine to add a bubble tea
func runAddRoutine(b bobba.Bobba) (int64, error) {
	var wg sync.WaitGroup
	var errs error
	var id int64
	wg.Add(1)

	// run the call to the database
	go func() {
		db, err := database.Create()
		if err != nil {
			errs = err
			wg.Done()
			return
		}

		inserted, opErr := db.AddBobba(b)
		if opErr != nil {
			errs = opErr
			wg.Done()
			return
		}

		id = inserted
		db.Close()
		wg.Done()
		return
	}()

	wg.Wait()

	return id, errs
}

// RunRemoveRoutine run a routine for removing a data in the database
// Param int id
// Return error
func runRemoveRoutine(id int) error {
	var wg sync.WaitGroup
	var errs error
	wg.Add(1)

	go func() {
		db, conErr := database.Create()
		if conErr != nil {
			errs = conErr
			wg.Done()
			return
		}

		err := db.RemoveBobba(id)
		if err != nil {
			errs = err
		}

		db.Close()
		wg.Done()
		return
	}()

	wg.Wait()

	return errs
}

// RunGetAllRoutine
// Run a goroutine to retrieve all the bobbas
// Return Bobba array
// Return error
func runGetAllRoutine() ([]bobba.Bobba, error) {
	var wg sync.WaitGroup
	var errs error
	var teas []bobba.Bobba
	wg.Add(1)

	go func() {
		db, conErr := database.Create()
		if conErr != nil {
			errs = conErr
			wg.Done()
			return
		}

		bobbas, err := db.GetAllBobba()
		if err != nil {
			errs = err
			wg.Done()
			return
		}

		teas = bobbas

		db.Close()
		wg.Done()
		return
	}()

	wg.Wait()

	return teas, errs
}

// RunGetDetailedRoutine
// Run a routine to get a single bobba based on the ID
// Param int id
// Return bobba Bobba
// Return error
func runGetDetailedRoutine(id int) (bobba.Bobba, error) {
	var wg sync.WaitGroup
	var errs error
	var tea bobba.Bobba
	wg.Add(1)

	go func() {
		db, conErr := database.Create()
		if conErr != nil {
			errs = conErr
			wg.Done()
			return
		}

		bobba, err := db.GetSingleBobba(id)
		if err != nil {
			errs = err
			wg.Done()
			return
		}

		tea = bobba

		db.Close()
		wg.Done()
		return
	}()

	wg.Wait()

	return tea, errs
}
