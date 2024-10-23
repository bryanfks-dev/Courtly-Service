package usecases

import (
	"main/data/models"
	"main/internal/providers/mysql"

	"github.com/golang-jwt/jwt/v5"
)

// BlacklistToken is a usecase that invalidates a token
// by removing it from the blacklisted_tokens table
//
// token: the token to invalidate
//
// Returns an error if the token could not be invalidated
func BlacklistToken(token *jwt.Token) error {
	// Get the token claims
	claims := DecodeToken(token)

	// Insert the token into the blacklisted_tokens table
	res := mysql.Conn.Create(&models.BlacklistedToken{
		Token:     token.Raw,
		ExpiresAt: claims.ExpiresAt.Time,
	})

	return res.Error
}
