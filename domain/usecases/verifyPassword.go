package usecases

import "golang.org/x/crypto/bcrypt"

// VerifyPassword is a function that verifies the password.
//
// password: The password to verify
// hashedPwd: The hashed password to compare with
//
// Retuns true if the password is correct, false otherwise
func VerifyPassword(password string, hashedPwd string) bool {
	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))

	return err == nil
}
