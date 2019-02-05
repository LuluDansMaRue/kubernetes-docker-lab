package database

import (
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := Config{
		username: os.Getenv("username"),
		password: os.Getenv("password"),
		host:     os.Getenv("host"),
		port:     os.Getenv("port"),
		database: os.Getenv("database"),
	}

	return config
}
