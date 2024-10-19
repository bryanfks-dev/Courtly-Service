package utils

import "strings"

// IsBlank is a function that checks if a string is blank.
//
// s: The string to check.
//
// Returns true if the string is blank, false otherwise.
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
