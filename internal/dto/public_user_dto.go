package dto

import (
	"fmt"
	"main/core/config"
	"main/data/models"
	"main/delivery/http/router"
	"main/pkg/utils"
)

// PublicUser is a struct that represents the public user dto.
type PublicUser struct {
	// ID is the primary key of the user.
	ID uint `json:"id"`

	// Username is the username of the user.
	Username string `json:"username"`

	// ProfilePictureUrl is the url of ther user profile picture.
	ProfilePictureUrl string `json:"profile_picture_url"`
}

// FromModel creates a PublicUser DTO from a User model.
//
// m: The User model.
//
// Returns a PublicUser DTO.
func (p PublicUser) FromModel(m models.User) PublicUser {
	// If the profile picture is blank, return the PublicUser DTO without the profile picture.
	if utils.IsBlank(m.ProfilePicture) {
		return PublicUser{
			ID:                m.ID,
			Username:          m.Username,
			ProfilePictureUrl: m.ProfilePicture,
		}
	}

	// profilePicturePath is the path to the profile picture.
	profilePicturePath := fmt.Sprintf("%s:%d%s/%s", config.ServerConfig.Host, config.ServerConfig.Port, router.UserProfiles, m.ProfilePicture)

	return PublicUser{
		ID:                m.ID,
		Username:          m.Username,
		ProfilePictureUrl: profilePicturePath,
	}
}
