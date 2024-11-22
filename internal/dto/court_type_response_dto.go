package dto

import "main/data/models"

// CourtTypeResponseDTO is a struct that defines the court type response DTO.
type CourtTypeResponseDTO struct {
	Types []string `json:"types"`
}

// FromModels is a function that converts court type models to court type response DTOs.
//
// courtTypes: The court types.
//
// Returns a new instance of the CourtTypeResponseDTO.
func (c CourtTypeResponseDTO) FromModels(courtTypes []models.CourtType) CourtTypeResponseDTO {
	// Create a new court type response DTO
	courtTypeResponseDTO := CourtTypeResponseDTO{}

	// Loop through the court types
	for _, courtType := range courtTypes {
		// Append the court type to the types
		courtTypeResponseDTO.Types = append(courtTypeResponseDTO.Types, courtType.Type)
	}

	return courtTypeResponseDTO
}
