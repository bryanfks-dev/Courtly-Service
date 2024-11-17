package entities

import (
	"main/core/enums"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims is a struct that represents the JWT claims.
type JWTClaims struct {
	// Id is the id of the user.
	Id uint `json:"id"`

	// ClientType is the client type of the user.
	ClientType enums.ClientType `json:"client_type"`

	// RegisteredClaims is the registered claims of the JWT.
	jwt.RegisteredClaims
}
