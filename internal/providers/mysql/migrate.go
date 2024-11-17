package mysql

import "main/data/models"

// Migrate is a function that migrates the database
// It creates the tables if they do not exist
//
// Returns an error if any
func Migrate() error {
	// Migrate the database
	return Conn.AutoMigrate(
		&models.User{},
		&models.BlacklistedToken{},
		&models.Vendor{},
		&models.CourtType{},
		&models.Court{},
		&models.PaymentMethod{},
		&models.Review{},
		&models.Book{},
		&models.Order{})
}
