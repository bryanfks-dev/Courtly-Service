package mysql

// Ping is a helper function that checks for database connection
//
// Returns an error if the database connection is not successful
func Ping() error {
	// Check for database connection
	err := DB.Ping()

	return err
}
