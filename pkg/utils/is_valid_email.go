package utils

import "regexp"

// IsValidEmail is a function that checks if an email is valid.
//
// s: The email.
//
// Returns a boolean.
func IsValidEmail(s string) bool {
	// Regular expression for email validation
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return regex.MatchString(s)
}
