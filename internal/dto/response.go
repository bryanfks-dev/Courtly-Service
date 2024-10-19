package dto

// Response is a struct that represents the response that is sent by the server.
// The response contains the status code, message, and data.
type Response struct {
	// StatusCode is the status code of the response.
	StatusCode int `json:"status_code"`

	// Message is the message of the response.
	Message any `json:"message"`

	// Data is the data of the response.
	Data any `json:"data"`
}
