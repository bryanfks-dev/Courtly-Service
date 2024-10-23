package config

import (
	"log"
	"main/pkg/utils"
	"strconv"
)

// Database is a struct that contains the mysql database
// configuration for the application.
type Database struct {
	// Hostname is the hostname of the database server.
	Hostname string

	// Port is the port number of the database server.
	Port int

	// Username is the username used to connect to the database.
	Username string

	// Password is the password used to connect to the database.
	Password string

	// DatabaseName is the name of the database to connect to.
	DatabaseName string
}

var DBConfig = Database{}

// LoadData is a method that loads the database configuration from the environment variables.
func (d Database) LoadData() {
	d.Hostname = utils.GetEnv("DB_HOSTNAME", "localhost")

	port, err := strconv.Atoi(utils.GetEnv("DB_PORT", "3306"))

	// Check if the port number is valid
	if err != nil {
		log.Fatal("Invalid port number")
	}

	d.Port = port

	d.Username = utils.GetEnv("DB_USERNAME", "root")

	d.Password = utils.GetEnv("DB_PASSWORD", "")

	d.DatabaseName = utils.GetEnv("DB_NAME", "courtly_db")

	DBConfig = d
}
