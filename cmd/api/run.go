package api

import "sync"

// Run is the entry point of the API
//
// Returns void
func Run() {
	// Create a wait group
	var wg sync.WaitGroup

	// Add 1 to the wait group
	wg.Add(1)

	// Initialize the database and server
	go func() {
		initDatabase()
		wg.Done()
	}()

	initServer()

	// Wait for the wait group to finish
	wg.Wait()
}
