package utils

// GetDigits is a function that returns the number of digits of a number.
//
// n: The number.
//
// Returns the number of digits.
//
// Example:
// GetDigits(123) -> 3
func GetDigits(n int) int {
	// Get the digits of the number
	digits := 0

	// Handle negative numbers
	if n < 0 {
		n = -n
	}

	// Count the digits
	for n > 0 {
		// Increment the digits
		digits++

		// Divide the number by 10
		n /= 10
	}

	return digits
}
