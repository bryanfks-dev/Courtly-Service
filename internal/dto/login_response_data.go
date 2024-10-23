package dto

// LoginResponseData is a type that represents the response data of the login response.
type LoginResponseData struct {
	// User is the current user.
	User CurrentUser `json:"user"`

	// Token is the JWT token.
	Token string `json:"token"`
}
