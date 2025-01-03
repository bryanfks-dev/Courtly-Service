package dto

import "main/data/models"

// CurrentUserOrdersResponseDTO is a struct that represents the current user orders
// data transfer object.
type CurrentUserOrdersResponseDTO struct {
	Orders *[]CurrentUserOrderDTO `json:"orders"`
}

// FromModels is a function that converts a slice of order models to a 
// current user orders response DTO.
//
// m: The slice of order models.
//
// Returns a pointer to the current user orders DTO.
func (c CurrentUserOrdersResponseDTO) FromModels(m *[]models.Order) *CurrentUserOrdersResponseDTO {
	// Create a slice of order DTOs
	dto := []CurrentUserOrderDTO{}

	// Convert the orders to order DTOs
	for _, order := range *m {
		dto = append(dto, *CurrentUserOrderDTO{}.FromModel(&order, nil))
	}

	return &CurrentUserOrdersResponseDTO{
		Orders: &dto,
	}
}
