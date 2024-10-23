package dto

// RegisterResponseData is a struct that represents the response data for the register endpoint.
type RegisterResponseData struct {
	// User is the current user.
	User CurrentUser `json:"user"`
}
