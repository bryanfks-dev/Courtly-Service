package entities

// ProcessError is a struct that represents the process error response.
type ProcessError struct {
	// Message is the error message.
	Message any

	// IsClientError is a flag that indicates if the error is a client error.
	ClientError bool
}
