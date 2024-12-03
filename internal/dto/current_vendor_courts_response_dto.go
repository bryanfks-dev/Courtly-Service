package dto

import "main/data/models"

// CurrentVendorCourtsDTO is a struct that defines the current vendor court data transfer object.
type CurrentVendorCourtsResponseDTO struct {
	Courts *[]CurrentVendorCourtDTO `json:"courts"`
}

// FromModels is a function that converts a slice of court models to a current vendor courts response DTO.
//
// m: The slice of court models.
//
// Returns the current vendor courts response DTO.
func (c CurrentVendorCourtsResponseDTO) FromModels(m *[]models.Court) *CurrentVendorCourtsResponseDTO {
	// courts is a placeholder for the courts
	courts := []CurrentVendorCourtDTO{}

	// Convert the court models to court DTOs
	for _, model := range *m {
		courts = append(courts, *CurrentVendorCourtDTO{}.FromModel(&model))
	}

	return &CurrentVendorCourtsResponseDTO{
		Courts: &courts,
	}
}
