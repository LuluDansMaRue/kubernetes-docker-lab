package main

import (
	"sync"

	"./bobba"
)

func runAddRoutine(b bobba.Bobba) (int64, error) {
	var wg sync.WaitGroup
	var errs error
	var id int64
	wg.Add(1)

	// run the call to the database
	go func() {
		inserted, err := db.AddBobba(b)
		if err != nil {
			errs = err
			wg.Done()
			return
		}

		id = inserted
		wg.Done()
		return
	}()

	wg.Wait()

	return id, errs
}
