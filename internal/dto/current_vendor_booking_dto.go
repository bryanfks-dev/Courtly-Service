package dto

import "main/data/models"

// CurrentVendorBookingDTO is a struct that represents the current vendor booking
// data transfer object.
type CurrentVendorBookingDTO struct {
	// ID is the ID of the booking
	ID uint `json:"id"`

	// Court is the court of the booking
	Court *CurrentVendorCourtDTO `json:"court"`

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
func (b CurrentVendorBookingDTO) FromModel(m *models.Booking) *CurrentVendorBookingDTO {
	// Get the start time
	startTime, _ := m.BookStartTime.Value()

	// Get the end time
	endTime, _ := m.BookEndTime.Value()

	return &CurrentVendorBookingDTO{
		ID:            m.ID,
		Court:         CurrentVendorCourtDTO{}.FromModel(&m.Court),
		BookStartTime: startTime.(string),
		BookEndTime:   endTime.(string),
	}
}
