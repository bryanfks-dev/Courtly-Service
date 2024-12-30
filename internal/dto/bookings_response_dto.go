package dto

import "main/data/models"

// BookingDTO is a struct that represents the booking data transfer object.
type BookingsResponseDTO struct {
	// Bookings is the bookings
	Bookings *[]BookingDTO `json:"bookings"`
}

// FromModels is a function that converts booking models to booking response DTOs.
//
// m: The booking models.
//
// Returns a new instance of the BookingsResponseDTO.
func (BookingsResponseDTO) FromModels(m *[]models.Booking) *BookingsResponseDTO {
	// Create a new booking response DTO
	bookings := []BookingDTO{}

	// Loop through the booking models
	for _, model := range *m {
		bookings = append(bookings, *BookingDTO{}.FromModel(&model))
	}

	return &BookingsResponseDTO{
		Bookings: &bookings,
	}
}
