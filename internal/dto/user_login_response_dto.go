package dto

// UserLoginResponseDTO is a struct that represents the response dto for the user login.
type UserLoginResponseDTO struct {
	// User is the current user.
	User *CurrentUserDTO `json:"user"`

	// Token is the JWT token.
	Token string `json:"token"`
}
