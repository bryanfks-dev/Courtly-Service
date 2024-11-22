package dto

import "main/data/models"

// CourtTypeResponseDTO is a struct that defines the court type response DTO.
type CourtTypeResponseDTO struct {
	Types *[]string `json:"types"`
}

// FromModels is a function that converts court type models to court type response DTOs.
//
// courtTypes: The court types.
//
// Returns a new instance of the CourtTypeResponse DTO.
func (c CourtTypeResponseDTO) FromModels(courtTypes *[]models.CourtType) *CourtTypeResponseDTO {
	// Create a new court type response DTO
	courtTypesString := []string{}

	// Loop through the court types
	for _, courtType := range *courtTypes {
		// Append the court type to the court types string 
		courtTypesString = append(courtTypesString, courtType.Type)
	}

	return &CourtTypeResponseDTO{
		Types: &courtTypesString,
	}
}
