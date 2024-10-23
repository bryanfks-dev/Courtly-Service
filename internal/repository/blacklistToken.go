package repository

import (
	"main/data/models"
	"main/internal/providers/mysql"
	"time"
)

// ClearBlacklistToken is a function that deletes all the expired tokens from the blacklist
// table in the database
//
// Returns an error if the operation was not successful
func ClearBlacklistToken() error {
	res := mysql.Conn.Delete(models.BlacklistedToken{}, "expires_at > ?", time.Now())

	return res.Error
}
