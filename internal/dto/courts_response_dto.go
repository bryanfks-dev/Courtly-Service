package dto

import "main/data/models"

// CourtsResponseDTO is a struct that defines the court response data transfer object.
type CourtsResponseDTO struct {
	// Courts is the list of courts.
	Courts *[]CourtDTO `json:"courts"`
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
