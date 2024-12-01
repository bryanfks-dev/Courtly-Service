package dto

import "main/data/models"

// CurrentVendorOrdersResponseDTO is a data transfer object that represents the response
type CurrentVendorOrdersResponseDTO struct {
	Orders *[]CurrentVendorOrderDTO `json:"orders"`
}

// FromModels is a function that converts a slice of booking models to a current vendor orders response DTO.
//
// m: The slice of booking models.
//
// Returns the current vendor orders response DTO.
func (c CurrentVendorOrdersResponseDTO) FromModels(m *[]models.Booking) *CurrentVendorOrdersResponseDTO {
	// orders is a slice of current user response DTOs
	orders := []CurrentVendorOrderDTO{}

	// Iterate over the booking models
	for _, booking := range *m {
		orders = append(orders, *CurrentVendorOrderDTO{}.FromModel(&booking))
	}

	return &CurrentVendorOrdersResponseDTO{
		Orders: &orders,
	}
}
