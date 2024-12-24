package api

// Run is the entry point of the API
//
// Returns void
func Run() {
	// Initialize the database
	initDatabase()

	// Initialize Midtrans
	initMidtrans()

	// Initialize the server
	initServer()
}
