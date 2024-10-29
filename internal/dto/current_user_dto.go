package dto

import (
	"fmt"
	"main/core/config"
	"main/data/models"
	"main/delivery/http/router"
	"main/pkg/utils"
)

// CurrentUser is a struct that represents the current user dto.
type CurrentUser struct {
	// ID is the primary key of the user.
	ID uint `json:"id"`

	// Username is the username of the user.
	Username string `json:"username"`

	// PhoneNumber is the phone number of the user.
	PhoneNumber string `json:"phone_number"`

	// ProfilePictureUrl is the profile picture of the user.
	ProfilePictureUrl string `json:"profile_picture_url"`
}

// FromModel creates a CurrentUser DTO from a User model.
func (c CurrentUser) FromModel(m *models.User) CurrentUser {
	// If the profile picture is blank, return the CurrentUser DTO without the profile picture.
	if utils.IsBlank(m.ProfilePicture) {
		return CurrentUser{
			ID:                m.ID,
			Username:          m.Username,
			PhoneNumber:       m.PhoneNumber,
			ProfilePictureUrl: m.ProfilePicture,
		}
	}

	// profilePicturePath is the path to the profile picture.
	profilePicturePath := fmt.Sprintf("%s:%d%s/%s", config.ServerConfig.Host, config.ServerConfig.Port, router.UserProfiles, m.ProfilePicture)

	return CurrentUser{
		ID:                m.ID,
		Username:          m.Username,
		PhoneNumber:       m.PhoneNumber,
		ProfilePictureUrl: profilePicturePath,
	}
}
