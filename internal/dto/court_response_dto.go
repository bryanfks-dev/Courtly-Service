package dto

import "main/data/models"

// CourtResponseDTO is a struct that defines the court response data transfer object.
type CourtResponseDTO struct {
	// Courts is the list of courts.
	Courts *[]CourtDTO `json:"courts"`
}

// FromModels is a function that converts court models to court response DTOs.
//
// models: The court models.
//
// Returns a new instance of the CourtResponseDTO.
func (c CourtResponseDTO) FromCourtModels(models *[]models.Court) *CourtResponseDTO {
	// Create a new court response DTO
	var courts []CourtDTO

	// Loop through the court models
	for _, model := range *models {
		courts = append(courts, *CourtDTO{}.FromModel(&model))
	}

	return &CourtResponseDTO{
		Courts: &courts,
	}
}
