package dto

import "main/data/models"

// ResponseDTO is a struct that represents the response data transfer object.
type BookingDTO struct {
	// ID is the ID of the booking
	ID uint `json:"id"`

	// Order is the order of the booking
	Order *OrderDTO `json:"order"`

	// CourtType is the type of the court
	CourtType string `json:"court_type"`

	// Vendor is the name of the vendor
	Vendor *VendorDTO `json:"vendor_name"`

	// Date is the date of the booking
	Date string `json:"date"`

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
	return &BookingDTO{
		ID:            m.ID,
		Order:         OrderDTO{}.FromModel(&m.Order),
		CourtType:     m.CourtType.Type,
		Vendor:        VendorDTO{}.FromModel(&m.Vendor),
		Date:          m.Date.String(),
		BookStartTime: m.BookStartTime.String(),
		BookEndTime:   m.BookEndTime.String(),
	}
}
