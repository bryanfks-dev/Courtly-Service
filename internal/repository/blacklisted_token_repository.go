package repository

import (
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
	var count int64

	err := mysql.Conn.Model(&models.BlacklistedToken{}).Where("token = ?", token).Limit(1).Count(&count).Error

	return count > 0, err
}

// Clear is a function that deletes all the expired tokens from the blacklist
// table in the database
//
// Returns an error if the operation was not successful
func (*BlacklistedTokenRepository) Clear() error {
	return mysql.Conn.Delete(&models.BlacklistedToken{}, "expires_at > ?", time.Now()).Error
}

// Remove is a function that deletes a token from the blacklist table in the database
//
// token: The token to be deleted
//
// Returns an error if the operation was not successful
func (*BlacklistedTokenRepository) Remove(token string) error {
	return mysql.Conn.Delete(&models.BlacklistedToken{}, "token = ?", token).Error
}
