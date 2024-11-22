package dto

import "main/data/models"

// VendorCourtResponseDTO is a struct that defines the VendorCourtTypeResponse DTO.
type VendorCourtResponseDTO struct {
	// Courts is a list of Court DTO.
	Courts *[]VendorCourtDTO `json:"courts"`
}

// FromCourtModels is a function that creates a VendorCourtResponse DTO from a list of Court models.
//
// m: The list of Court models.
//
// Returns a VendorCourtResponse DTO.
func (v VendorCourtResponseDTO) FromCourtModels(m *[]models.Court) *VendorCourtResponseDTO {
	// Create a new list of VendorCourtDTO
	courts := []VendorCourtDTO{}

	// Iterate through the courts
	for _, court := range *m {
		courts = append(courts, VendorCourtDTO{
			ID:    court.ID,
			Name:  court.Name,
			Price: court.Price,
		})
	}

	return &VendorCourtResponseDTO{
		Courts: &courts,
	}
}
