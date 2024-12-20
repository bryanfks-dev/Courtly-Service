package dto

import (
	"main/data/models"
)

// ResponseDTO is a struct that represents the response data transfer object.
type BookingDTO struct {
	// ID is the ID of the booking
	ID uint `json:"id"`

	// Court is the court of the booking
	Court *CourtDTO `json:"court"`

	// BookStartTime is the start time of the booking
	BookStartTime string `json:"book_start_time"`

	// BookEndTime is the end time of the booking
	BookEndTime string `json:"book_end_time"`
}

// FromModel is a function that converts a booking model to a booking DTO.
//
// m: The booking model.
//
// Returns the booking DTO.
func (b BookingDTO) FromModel(m *models.Booking) *BookingDTO {
	// Get the start time
	startTime, _ := m.BookStartTime.Value()

	// Get the end time
	endTime, _ := m.BookEndTime.Value()

	return &BookingDTO{
		ID:            m.ID,
		Court:         CourtDTO{}.FromModel(&m.Court),
		BookStartTime: startTime.(string),
		BookEndTime:   endTime.(string),
	}
}
