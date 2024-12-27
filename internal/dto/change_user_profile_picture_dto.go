package dto

// ChangeUserProfilePictureDTO is a struct that represents the 
// data transfer object for changing the user profile picture.
type ChangeUserProfilePictureDTO struct {
	// ProfilePicture is the profile picture of the user.
	Image string `json:"image"`
}
