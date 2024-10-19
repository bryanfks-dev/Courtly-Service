package config

import "main/pkg/utils"

// JWT is a struct that contains the JWT configuration.
type JWT struct {
	// Secret is the secret used to sign the JWT token.
	Secret string
}

// LoadData is a method that loads the data for the JWT configuration.
func (j JWT) LoadData() JWT {
	j.Secret = utils.GetEnv("JWT_SECRET", "my_secret")

	return j
}

// JWTConfig is the global variable that holds the JWT configuration.
var JWTConfig = JWT{}.LoadData()
