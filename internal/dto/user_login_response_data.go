package dto

// UserLoginResponseData is a struct that represents the response data for the user login.
type UserLoginResponseData struct {
	// User is the current user.
	User *CurrentUser `json:"user"`

	// Token is the JWT token.
	Token string `json:"token"`
}
