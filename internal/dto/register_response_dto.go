package dto

// RegisterResponseDTO is a struct that represents the response dto for the register endpoint.
type RegisterResponseDTO struct {
	// User is the current user.
	User *CurrentUserDTO `json:"user"`
}
