package utils

import "strings"

// ToUpperFirst converts the first character of a string to uppercase.
//
// s: The string to convert.
//
// Return the converted string with the first character in uppercase.
func ToUpperFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
