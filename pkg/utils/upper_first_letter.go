package utils

import "strings"

// UpperFirstLetter is a function that uppercases the first letter of a string.
//
// s: The string.
//
// Returns the string with the first letter uppercased.
func UpperFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
