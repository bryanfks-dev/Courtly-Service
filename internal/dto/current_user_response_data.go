package dto

import "main/data/models"

// CurrentUserResponseData is a struct that represents the response data for the current user.
type CurrentUserResponseData struct {
	User models.User `json:"user"`
}
