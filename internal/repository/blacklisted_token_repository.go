package repository

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"
	"time"
)

// BlacklistedTokenRepository is an interface that defines the blacklisted token repository
type BlacklistedTokenRepository struct{}

// NewBlacklistedTokenRepository is a factory function that returns a new instance of the BlacklistedTokenRepository
//
// Returns a new instance of the BlacklistedTokenRepository
func NewBlacklistedTokenRepository() *BlacklistedTokenRepository {
	return &BlacklistedTokenRepository{}
}

// Create is a function that adds a token to the blacklist table in the database
//
// token: The token to be added to the blacklist
//
// Returns an error if the operation was not successful
func (*BlacklistedTokenRepository) Create(token *models.BlacklistedToken) error {
	return mysql.Conn.Create(token).Error
}

// IsBlacklisted is a function that checks if a token is blacklisted
//
// token: The token to be checked
//
// Returns a boolean value
func (*BlacklistedTokenRepository) IsBlacklisted(token string) (bool, error) {
	// count is a variable that holds the number of blacklisted tokens
	var count int64

	// Check if the token is blacklisted
	err := mysql.Conn.Model(&models.BlacklistedToken{}).Where("token = ?", token).Limit(1).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking if token is blacklisted: ", err)

		return false, err
	}

	return count > 0, err
}

// Clear is a function that deletes all the expired tokens from the blacklist
// table in the database
//
// Returns an error if the operation was not successful
func (*BlacklistedTokenRepository) Clear() error {
	// Delete all the expired tokens
	err := mysql.Conn.Delete(&models.BlacklistedToken{}, "expires_at > ?", time.Now()).Error

	// Return an error if any
	if err != nil {
		log.Println("Error clearing blacklisted tokens: ", err)

		return err
	}

	return nil
}

// Delete is a function that deletes a token from the blacklist table in the database
//
// token: The token to be deleted
//
// Returns an error if the operation was not successful
func (*BlacklistedTokenRepository) Delete(token string) error {
	// Delete the token
	err := mysql.Conn.Delete(&models.BlacklistedToken{}, "token = ?", token).Error

	// Return an error if any
	if err != nil {
		log.Println("Error deleting blacklisted token: ", err)

		return err
	}

	return nil
}
