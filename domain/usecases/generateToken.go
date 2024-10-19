package usecases

import (
	"main/core/config"
	"main/domain/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken is a function that generates a JWT token.
//
// id: the id of the user
//
// Returns a string containing the token and an error if there is any
func GenerateToken(id uint) (string, error) {
	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.JWTClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	})

	// Sign the token with the secret
	return token.SignedString([]byte(config.JWTConfig.Secret))
}
