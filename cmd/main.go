package main

import (
	"log"
	"main/cmd/api"
	"main/cmd/routines"
	"main/internal/providers/mysql"

	"github.com/joho/godotenv"
)

// Main is the entry point of the application
func main() {
	// Load the environment variables
	err := godotenv.Load()

	// Check if there is an error loading the environment variables
	if err != nil {
		panic("Error loading environment variables: " + err.Error())
	}

	// Start the routines
	routines.Run()

	// Start the API
	api.Run()

	// Close the database connection
	defer func() {
		err := mysql.CloseConnection()

		// Check if there is an error closing the database connection
		if err != nil {
			log.Fatal("Error closing the database connection: " + err.Error())
		}
	}()
}
