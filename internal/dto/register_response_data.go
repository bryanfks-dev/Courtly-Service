package dto

import "main/data/models"

// RegisterResponseData is a struct that represents the response data for the register endpoint.
type RegisterResponseData struct {
	User models.User `json:"user"`
}
