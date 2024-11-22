package dto

// CurrentUserResponseDTO is a struct that represents the response dto for the current user.
type CurrentUserResponseDTO struct {
	// User is the current user.
	User *CurrentUserDTO `json:"user"`
}
