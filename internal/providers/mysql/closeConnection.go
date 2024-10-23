package mysql

// CloseConnection is a function that closes the database connection
//
// Returns an error if the operation was not successful
func CloseConnection() error {
	// Close the database connection
	err := DB.Close()

	return err
}
