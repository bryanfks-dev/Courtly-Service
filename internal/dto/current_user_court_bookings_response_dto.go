package dto

import "main/data/models"

// CurrentUserCourtBookingsResponseDTO is a data transfer object that
// represents the current user court bookings response.
type CurrentUserCourtBookingsResponseDTO struct {
	// Bookings is the bookings of the current user.
	Bookings *[]CurrentUserBookingDTO `json:"bookings"`
}

// FromModels is a function that converts booking models to a current 
// user court booking response DTOs.
//
// m: The booking models.
//
// Returns a new instance of the CurrentUserCourtBookingsResponseDTO.
func (b CurrentUserCourtBookingsResponseDTO) FromModels(m *[]models.Booking) *CurrentUserCourtBookingsResponseDTO {
	// Create a new booking response DTO
	bookings := []CurrentUserBookingDTO{}

	// Loop through the booking models
	for _, model := range *m {
		bookings = append(bookings, *CurrentUserBookingDTO{}.FromModel(&model))
	}

	return &CurrentUserCourtBookingsResponseDTO{
		Bookings: &bookings,
	}
}
