package dto

// ChangeUsernameFormDTO is a struct that defines the data transfer object for updating the username.
type ChangeUsernameFormDTO struct {
	// NewUsername is the new user's username.
	NewUsername string `json:"new_username"`
}
