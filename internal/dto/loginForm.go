package dto

// RegisterForm is a struct that represents the login form
// that is sent by the client.
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
