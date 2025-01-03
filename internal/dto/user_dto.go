package dto

import (
	"fmt"
	"main/data/models"
	"main/delivery/http/router"
	"main/pkg/utils"
)

// UserDTO is a struct that represents the user dto.
type UserDTO struct {
	// ID is the primary key of the user.
	ID uint `json:"id"`

	// Username is the username of the user.
	Username string `json:"username"`

	// ProfilePictureUrl is the url of ther user profile picture.
	ProfilePictureUrl string `json:"profile_picture_url"`
}

// FromModel creates a User DTO from a User model.
//
// m: The User model.
//
// Returns a User DTO.
func (p UserDTO) FromModel(m *models.User) *UserDTO {
	// If the profile picture is blank, return the UserDTO without the profile picture.
	if utils.IsBlank(m.ProfilePicture) {
		return &UserDTO{
			ID:                m.ID,
			Username:          m.Username,
			ProfilePictureUrl: m.ProfilePicture,
		}
	}

	// profilePicturePath is the path to the profile picture.
	profilePicturePath := fmt.Sprintf("%s/%s", router.UserProfiles, m.ProfilePicture)

	return &UserDTO{
		ID:                m.ID,
		Username:          m.Username,
		ProfilePictureUrl: profilePicturePath,
	}
}
