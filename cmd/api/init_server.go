package api

import (
	"main/internal/server"
)

// initServer is a helper function that initialize the server.
//
// Returns void
func initServer() {
	// Create a new server
	server, err := server.NewServer()

	server.Logger.Fatal(err)
}
