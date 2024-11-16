package api

// Run is the entry point of the API
//
// Returns void
func Run() {
	// Initialize the database and server

	initDatabase()

	initServer()
}
