package dto

// PublicUserResponseDTO is a struct that represents the public user response dto.
type PublicUserResponseDTO struct {
	// Users is a field that represents the public user.
	User *PublicUserDTO `json:"user"`
}
