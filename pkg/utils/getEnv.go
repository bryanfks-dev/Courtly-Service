package utils

import "os"

// GetEnv is a helper function that returns the value of an environment variable
// or a default value if the environment variable is not set.
//
// key: The name of the environment variable.
// defaultValue: The default value to return if the environment variable is not set.
//
// Returns the value of the environment variable or the default value if the environment variable is not set.
func GetEnv(key string, defaultValue string) string {
	// LookupEnv retrieves the value of the environment variable named by the key.
	val, ok := os.LookupEnv(key)

	// If the environment variable is not set, return the default value.
	if !ok {
		return defaultValue
	}

	return val
}
