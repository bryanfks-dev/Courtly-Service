package usecases

import "github.com/golang-jwt/jwt/v5"

// DecodeToken is a function that decodes a JWT token
//
// token: the token to decode
//
// Returns the claims of the token
func DecodeToken(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}
