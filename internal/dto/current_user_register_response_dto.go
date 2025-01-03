package dto

// CurrentUserRegisterResponseDTO is a struct that represents the response dto for
// the current user register.
type CurrentUserRegisterResponseDTO struct {
	// User is the current user.
	User *CurrentUserDTO `json:"user"`
}
