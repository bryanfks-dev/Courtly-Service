package dto

import "main/data/models"

// CurrentUserBookingsResponseDTO is a struct that represents
// the current user bookings response data transfer object.
type CurrentUserBookingsResponseDTO struct {
	// Bookings is the list of bookings
	Bookings *[]BookingDTO `json:"bookings"`
}

// FromModels is a function that converts a list of booking models
// to a list of booking DTOs.
//
// m: The list of booking models.
//
// Returns the list of booking DTOs.
func (c CurrentUserBookingsResponseDTO) FromModels(m *[]models.Booking) *CurrentUserBookingsResponseDTO {
	// bookings is a placeholder for the bookings
	bookings := []BookingDTO{}

	// Convert the booking models to booking DTOs
	for _, model := range *m {
		bookings = append(bookings, *BookingDTO{}.FromModel(&model))
	}

	return &CurrentUserBookingsResponseDTO{
		Bookings: &bookings,
	}
}
