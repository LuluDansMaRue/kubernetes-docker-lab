package utils

import (
	"database/sql"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driver  string = "mysql"
	logpath string = "/var/log/bobba-cron.log"
)

// BuildDBEndpoint build a connection to the database
// Param string username
// Param string pwd
// Return string
func BuildDBEndpoint(username string, pwd string) string {
	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DATABASE")

	return username + ":" + pwd + "@tcp(" + host + ")/" + db
}

// GetCon get the connection instance
// Get the connection instance to the database
// Param bool generated
// Return database *sql.DB
func GetCon(generated bool) *sql.DB {
	var username string
	var pwd string

	if generated {
		username = gofakeit.Username()
		pwd = gofakeit.Password(true, true, true, true, true, 32)
	} else {
		username = os.Getenv("MYSQL_USER")
		pwd = os.Getenv("MYSQL_PASSWORD")
	}

	endpoint := BuildDBEndpoint(username, pwd)
	db, err := sql.Open(driver, endpoint)

	if err != nil {
		panic(err)
	}

	return db
}

// GetRandInt return a random int between 2 numbers
// Generate a random number
// Return int
func GetRandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// SetLogOutput set the log file output
// See this stack: https://stackoverflow.com/questions/19965795/go-golang-write-log-to-file
// Return * os.File
func SetLogOutput() *os.File {
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
