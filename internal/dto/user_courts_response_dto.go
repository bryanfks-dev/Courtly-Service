package dto

import (
	"main/core/types"
	"main/data/models"
)

// UserCourtsResponseDTO is a struct that defines the court response data transfer object
// for user client type.
type UserCourtsResponseDTO struct {
	// Courts is the list of courts.
	Courts *[]UserCourtDTO `json:"courts"`
}

// FromCourtMaps is a function that converts court maps to user court response DTOs.
//
// m: The court maps.
//
// Returns a new instance of the UserCourtsResponseDTO.
func (c UserCourtsResponseDTO) FromCourtMaps(m *[]types.CourtMap) *UserCourtsResponseDTO {
	// Create a new court response DTO
	courts := []UserCourtDTO{}

	// Loop through the court maps
	for _, courtMap := range *m {
		courts = append(courts, *UserCourtDTO{}.FromCourtMap(&courtMap))
	}

	return &UserCourtsResponseDTO{
		Courts: &courts,
	}
}

// FromModels is a function that converts court models to user court response DTOs.
//
// models: The court models.
//
// Returns a new instance of the UserCourtsResponseDTO.
func (c UserCourtsResponseDTO) FromCourtModels(models *[]models.Court) *UserCourtsResponseDTO {
	// Create a new court response DTO
	courts := []UserCourtDTO{}

	// Loop through the court models
	for _, model := range *models {
		courts = append(courts, *UserCourtDTO{}.FromModel(&model))
	}

	return &UserCourtsResponseDTO{
		Courts: &courts,
	}
}
