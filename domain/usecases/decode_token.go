package usecases

import (
	"main/domain/entities"

	"github.com/golang-jwt/jwt/v5"
)

// DecodeToken is a function that decodes a JWT token
//
// token: the token to decode
//
// Returns the decoded token.
func DecodeToken(token *jwt.Token) *entities.JWTClaims {
	return token.Claims.(*entities.JWTClaims)
}
