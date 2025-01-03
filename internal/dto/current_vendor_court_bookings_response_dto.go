package dto

import "main/data/models"

// CurrentVendorCourtBookingsResponseDTO is a data transfer object that
// represents the current vendor court bookings response.
type CurrentVendorCourtBookingsResponseDTO struct {
	// Bookings is the bookings of the current vendor.
	Bookings []CurrentVendorBookingDTO `json:"bookings"`
}

// FromModels is a function that converts a list of booking models to a list of
// current vendor booking DTOs.
//
// bookings: the list of booking models
//
// Returns a list of current vendor booking DTOs.
func (c CurrentVendorCourtBookingsResponseDTO) FromModels(m *[]models.Booking) *CurrentVendorCourtBookingsResponseDTO {
	// Create a list of current vendor booking DTOs
	var bookings []CurrentVendorBookingDTO

	// Iterate through the list of booking models
	for _, booking := range *m {
		bookings = append(bookings, *CurrentVendorBookingDTO{}.FromModel(&booking))
	}

	return &CurrentVendorCourtBookingsResponseDTO{
		Bookings: bookings,
	}
}
