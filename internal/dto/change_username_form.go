package dto

// ChangeUsernameForm is a struct that defines the data transfer object for updating the username.
type ChangeUsernameForm struct {
	// Username is the new username.
	Username string `json:"username"`
}
