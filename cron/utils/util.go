package utils

import (
	"database/sql"
	"math/rand"
	"os"
	"time"
)

const (
	driver string = "mysql"
)

// BuildDBEndpoint build a connection to the database
// Return string
func BuildDBEndpoint() string {
	username := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DATABASE")

	return username + ":" + pwd + "@tcp(" + host + ")/" + db
}

// GetCon get the connection instance
// Get the connection instance to the database
// Return database *sql.DB
func GetCon() *sql.DB {
	endpoint := BuildDBEndpoint()
	db, err := sql.Open(driver, endpoint)

	if err != nil {
		panic(err)
	}

	return db
}

// GetRandInt return a random int between 2 numbers
// Generate a random number
// Return int
func GetRandInt(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
