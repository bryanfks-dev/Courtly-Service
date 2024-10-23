package entities

import "github.com/golang-jwt/jwt/v5"

// JWTClaims is a struct that represents the JWT claims.
type JWTClaims struct {
	// Id is the id of the user.
	Id uint `json:"id"`

	jwt.RegisteredClaims
}
