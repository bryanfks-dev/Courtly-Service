package api

import (
	"main/core/config"
	"main/internal/providers/database"
)

// initDatabase is a helper function that initialize the database.
//
// Returns void
func initDatabase() {
	// Start the database connection
	err := database.Connect(config.DBConfig)

	// Check if there is an error connecting to the database
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	// Run migrations
	err = database.Migrate()

	// Check if there is an error migrating the database
	if err != nil {
		panic("Error migrating the database: " + err.Error())
	}
}
