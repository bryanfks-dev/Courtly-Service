package dto

// Response is a struct that represents the response that is sent by the server.
// The response contains the status code, message, and data.
type Response struct {
	// Success is the status of the response.
	Success bool `json:"success"`

	// Message is the message of the response.
	Message any `json:"message"`

	// Data is the data of the response.
	Data any `json:"data"`
}
