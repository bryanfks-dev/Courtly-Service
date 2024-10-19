package usecases

import "golang.org/x/crypto/bcrypt"

// HashPassword is a function that hashes a password.
//
// password: the password to hash
//
// Returns the hashed password and an error if there is one
func HashPassword(password string) (string, error) {
	// Generate a hashed password from the password string
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}
