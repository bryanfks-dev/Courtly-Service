package dto

// ChangeUsernameForm is a struct that defines the data transfer object for updating the username.
type ChangeUsernameForm struct {
	// NewUsername is the new user's username.
	NewUsername string `json:"new_username"`
}
