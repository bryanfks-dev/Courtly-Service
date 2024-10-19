package repository

import (
	"main/data/models"
	"main/internal/providers/database"
	"time"
)

// ClearBlacklistToken is a function that deletes all the expired tokens from the blacklist
// table in the database
//
// Returns an error if the operation was not successful
func ClearBlacklistToken() error {
	res := database.Conn.Delete(models.BlacklistedToken{}, "expires_at > ?", time.Now())

	return res.Error
}
