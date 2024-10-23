package api

import (
	"main/internal/providers/mysql"
)

// initDatabase is a helper function that initialize the database.
//
// Returns void
func initDatabase() {
	// Start the database connection
	err := mysql.Connect()

	// Check if there is an error connecting to the database
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	// Run migrations
	err = mysql.Migrate()

	// Check if there is an error migrating the database
	if err != nil {
		panic("Error migrating the database: " + err.Error())
	}
}
