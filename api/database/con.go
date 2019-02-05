package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// BobbaCon representing the SQL Driver
type BobbaCon struct {
	connection *sql.DB
}

const (
	driver string = "mysql"

	// errors
	instanceErr string = "Unable to create a database instance"
)

// Create create the database connection instance
// Return BobbaCon*
// Return error
func Create() (*BobbaCon, error) {
	config := parseConfig()
	db, err := sql.Open(driver, config.username+":"+config.password+"@/tcp("+config.host+")/"+config.database)

	if err != nil {
		return nil, errors.New(instanceErr)
	}

	instance := BobbaCon{
		connection: db,
	}

	return &instance, nil
}
