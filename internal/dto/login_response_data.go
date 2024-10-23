package dto

import "main/data/models"

// LoginResponseData is a type that represents the response data of the login response.
type LoginResponseData struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
