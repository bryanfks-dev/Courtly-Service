package dto

import (
	"main/core/types"
	"main/data/models"
)

// CourtsResponseDTO is a struct that defines the court response data transfer object.
type CourtsResponseDTO struct {
	// Courts is the list of courts.
	Courts *[]CourtDTO `json:"courts"`
}

// FromCourtMaps is a function that converts court maps to court response DTOs.
//
// m: The court maps.
//
// Returns a new instance of the CourtsResponseDTO.
func (c CourtsResponseDTO) FromCourtMaps(m *[]types.CourtMap) *CourtsResponseDTO {
	// Create a new court response DTO
	courts := []CourtDTO{}

	// Loop through the court maps
	for _, courtMap := range *m {
		courts = append(courts, *CourtDTO{}.FromCourtMap(&courtMap))
	}

	return &CourtsResponseDTO{
		Courts: &courts,
	}
}

// FromModels is a function that converts court models to court response DTOs.
//
// models: The court models.
//
// Returns a new instance of the CourtsResponseDTO.
func (c CourtsResponseDTO) FromCourtModels(models *[]models.Court) *CourtsResponseDTO {
	// Create a new court response DTO
	courts := []CourtDTO{}

	// Loop through the court models
	for _, model := range *models {
		courts = append(courts, *CourtDTO{}.FromModel(&model))
	}

	return &CourtsResponseDTO{
		Courts: &courts,
	}
}
