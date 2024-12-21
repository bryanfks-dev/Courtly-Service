package dto

import "main/data/models"

// CurrentUserOrdersDTO is a struct that represents the current user orders
// data transfer object.
type CurrentUserOrdersDTO struct {
	Orders *[]OrderDTO `json:"orders"`
}

// FromModels is a function that converts a slice of order models to a current user orders DTO.
//
// m: The slice of order models.
//
// Returns a pointer to the current user orders DTO.
func (c CurrentUserOrdersDTO) FromModels(m *[]models.Order) *CurrentUserOrdersDTO {
	// Create a slice of order DTOs
	dto := []OrderDTO{}

	// Convert the orders to order DTOs
	for _, order := range *m {
		dto = append(dto, *OrderDTO{}.FromModel(&order, nil))
	}

	return &CurrentUserOrdersDTO{
		Orders: &dto,
	}
}
