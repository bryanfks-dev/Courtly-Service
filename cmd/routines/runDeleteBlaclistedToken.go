package routines

import (
	"log"
	"main/internal/providers/database"
	"main/internal/repository"
)

// runClearBlacklistedToken is a helper function that runs the clear blacklist token routine.
// This routine will delete the blacklist token every 24 hours.
//
// Returns void
func runClearBlacklistedToken() {
	// Check for database connection
	err := database.Ping()

	// Check if there is an error with the database connection
	if err != nil {
		log.Fatal("Error connecting to the database: " + err.Error())
	}

	// Delete the blacklist token
	err = repository.ClearBlacklistToken()

	// Check if there is an error deleting the blacklist token
	if err != nil {
		log.Fatal("Error deleting blacklist token: " + err.Error())
	} else {
		// Log the success of the cleanup
		log.Println("Blacklist token cleaned up")
	}
}
