package usecases

import (
	"errors"
	"main/core/config"

	"github.com/golang-jwt/jwt/v5"
)

// VerifyToken is a function that verifies a JWT token.
//
// tokenString: the token to verify
//
// Returns a boolean indicating if the token is valid and a jwt.Token object
func VerifyToken(tokenString string) (*jwt.Token, bool) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		// Check if the token is signed with the correct signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.JWTConfig.Secret), nil
	})

	// Check if there is an error parsing the token
	if err != nil {
		return nil, false
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, false
	}

	return token, true
}
