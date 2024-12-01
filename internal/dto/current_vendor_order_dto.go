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
	Court *CurrentVendorCourtDTO `json:"court"`

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
	return &CurrentVendorOrderDTO{
		ID:            m.Order.ID,
		Date:          m.Date.String(),
		User:          PublicUserDTO{}.FromModel(&m.User),
		Court:         CurrentVendorCourtDTO{}.FromModel(&m.Court),
		BookStartTime: m.BookStartTime.String(),
		BookEndTime:   m.BookEndTime.String(),
		Price:         m.Order.Price,
	}
}

// FromModels is a function that converts a slice of booking models to
// a slice of booking DTOs.
//
// m: The slice of booking models
//
// Returns a pointer to a slice of booking DTOs
func (c CurrentVendorOrderDTO) FromModels(m *[]models.Booking) *[]CurrentVendorOrderDTO {
	// Create a slice of booking DTOs
	orders := []CurrentVendorOrderDTO{}

	// Iterate over the slice of booking models
	for _, order := range *m {
		orders = append(orders, *c.FromModel(&order))
	}

	return &orders
}
