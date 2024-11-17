package dto

// CurrentUserResponseData is a struct that represents the response data for the current user.
type CurrentUserResponseData struct {
	// User is the current user.
	User *CurrentUser `json:"user"`
}
