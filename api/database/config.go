package database

import (
	"os"
)

// Config structure
// Config structure represent the configuration of the mysql database
type Config struct {
	username string
	password string
	host     string
	port     string
	database string
}

func parseConfig() Config {
	config := Config{
		username: os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		host:     os.Getenv("MYSQL_HOST"),
		port:     os.Getenv("MYSQL_PORT"),
		database: os.Getenv("MYSQL_DATABASE"),
	}

	return config
}
