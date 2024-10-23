package dto

import (
	"fmt"
	"main/core/config"
	"main/data/models"
	"main/delivery/http/router"
)

// PublicUser is a struct that represents the public user dto.
type PublicUser struct {
	// ID is the primary key of the user.
	Username string `json:"username"`

	// PhoneNumber is the phone number of the user.
	ProfilePicture string `json:"profile_picture"`
}

// FromModel creates a PublicUser DTO from a User model.
func (p PublicUser) FromModel(m models.User) PublicUser {
	// profilePicturePath is the path to the profile picture.
	profilePicturePath := fmt.Sprintf("%s:%d%s/%s", config.ServerConfig.Host, config.ServerConfig.Port, router.UserProfiles, m.ProfilePicture)

	return PublicUser{
		Username:       m.Username,
		ProfilePicture: profilePicturePath,
	}
}
