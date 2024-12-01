package dto

import "main/data/models"

// CurrentVendorOrdersResponseDTO is a data transfer object that represents the response
type CurrentVendorOrderDTO struct {
	// ID is the unique identifier of the order
	ID uint `json:"id"`

	// Date is the date of the order
	Date string `json:"date"`

	// User is the user who made the order
	User *PublicUserDTO `json:"user"`

	// Court is the ordered court
	Court *CourtDTO `json:"court"`

	// BookStartTime is the start time of the booking
	BookStartTime string `json:"book_start_time"`

	// BookEndTime is the end time of the booking
	BookEndTime string `json:"book_end_time"`

	// Price is the price of the booking
	Price float64 `json:"price"`
}

// CurrentVendorOrdersResponseDTO is a data transfer object that represents the response
//
// m: The booking model
//
// Returns a CurrentVendorOrdersResponseDTO
func (c CurrentVendorOrderDTO) FromModel(m *models.Booking) *CurrentVendorOrderDTO {
	// Get the date
	date, _ := m.Date.Value()

	// Get the start and end time
	bookStartTime, _ := m.BookStartTime.Value()

	bookEndTime, _ := m.BookEndTime.Value()

	return &CurrentVendorOrderDTO{
		ID:            m.Order.ID,
		Date:          date.(string),
		User:          PublicUserDTO{}.FromModel(&m.User),
		Court:         CourtDTO{}.FromModel(&m.Court),
		BookStartTime: bookStartTime.(string),
		BookEndTime:   bookEndTime.(string),
		Price:         m.Order.Price,
	}
}
