package dto

// UserLoginForm is a struct that represents the login form
// that is sent by the user.
type UserLoginForm struct {
	// Username is the username of the user.
	Username string `json:"username"`

	// Password is the password of the user.
	Password string `json:"password"`
}
