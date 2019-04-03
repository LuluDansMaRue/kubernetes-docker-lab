package database

import (
	"github.com/LuluDansMaRue/kubernetes-docker-lab/bobba"
)

// AddBobba add a bubble tea into the database
func (db *BobbaCon) AddBobba(b bobba.Bobba) (int64, error) {
	stmt, err := db.connection.Prepare("INSERT INTO bobba SET name = ?, price = ?, shop = ?, flavor = ?, calory = ?")

	if err != nil {
		return 0, err
	}

	row, errInsert := stmt.Exec(
		b.Name,
		b.Price,
		b.Shop,
		b.Flavor,
		b.Calory,
	)

	if errInsert != nil {
		return 0, errInsert
	}

	insertedID, _ := row.LastInsertId()
	db.Close()

	return insertedID, nil
}

// GetAllBobba get all the bubble tea
// Return an array of bobba
// Return error
func (db *BobbaCon) GetAllBobba() ([]bobba.Bobba, error) {
	var bobbas []bobba.Bobba
	rows, err := db.connection.Query("SELECT * FROM bobba")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var name string
		var price float32
		var shop string
		var flavor string
		var calory float32

		err = rows.Scan(&id, &name, &price, &shop, &flavor, &calory)
		if err != nil {
			panic(err)
		}

		b := bobba.Bobba{
			ID:     id,
			Name:   name,
			Price:  price,
			Shop:   shop,
			Flavor: flavor,
			Calory: calory,
		}

		bobbas = append(bobbas, b)
	}

	db.Close()

	return bobbas, nil
}

// RemoveBobba remove a bubble tea
// Param string id
// Return error
func (db *BobbaCon) RemoveBobba(id int) error {
	stmt, err := db.connection.Prepare("DELETE FROM bobba WHERE id = ?")
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(id)
	if execErr != nil {
		return execErr
	}

	db.Close()

	return nil
}

// GetSingleBobba get a single bubble tea
// Param string id
// Return bobba
// Return error
func (db *BobbaCon) GetSingleBobba(id int) (bobba.Bobba, error) {
	var b bobba.Bobba
	stmt, err := db.connection.Prepare("SELECT * FROM bobba WHERE id = ?")
	if err != nil {
		return b, err
	}

	row := stmt.QueryRow(id)
	var dbID int
	var name string
	var price float32
	var shop string
	var flavor string
	var calory float32

	err = row.Scan(&dbID, &name, &price, &shop, &flavor, &calory)
	if err != nil {
		return b, err
	}

	b.ID = dbID
	b.Name = name
	b.Price = price
	b.Shop = shop
	b.Flavor = flavor
	b.Calory = calory

	db.Close()

	return b, nil
}
