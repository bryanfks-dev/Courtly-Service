package dto

import (
	"main/data/models"
)

// CurrentUserBookingDTO is a data transfer object that represents the current user booking.
type CurrentUserBookingDTO struct {
	// ID is the ID of the booking
	ID uint `json:"id"`

	// Court is the court of the booking
	Court *UserCourtDTO `json:"court"`

	// BookStartTime is the start time of the booking
	BookStartTime string `json:"book_start_time"`

	// BookEndTime is the end time of the booking
	BookEndTime string `json:"book_end_time"`
}

// FromModel is a function that converts a booking model to a current user booking DTO.
//
// m: The booking model.
//
// Returns the booking DTO.
func (b CurrentUserBookingDTO) FromModel(m *models.Booking) *CurrentUserBookingDTO {
	// Get the start time
	startTime, _ := m.BookStartTime.Value()

	// Get the end time
	endTime, _ := m.BookEndTime.Value()

	return &CurrentUserBookingDTO{
		ID:            m.ID,
		Court:         UserCourtDTO{}.FromModel(&m.Court),
		BookStartTime: startTime.(string),
		BookEndTime:   endTime.(string),
	}
}
