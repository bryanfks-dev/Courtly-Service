package dto

import "main/data/models"

// CurrentVendorOrdersResponseDTO is a data transfer object that represents the response
type CurrentVendorOrderDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

	// Date is the date of the order was created
	Date string `json:"date"`

	// User is the user of the order
	User *PublicUserDTO `json:"user"`

	// CourtType is the court type of the order
	CourtType string `json:"court_type"`

	// Price is the price of the order
	Price float64 `json:"price"`

	// AppFee is the application fee of the order
	AppFee float64 `json:"app_fee"`
}

// CurrentVendorOrdersResponseDTO is a data transfer object that represents the response
//
// m: The booking model
//
// Returns a CurrentVendorOrdersResponseDTO
func (c CurrentVendorOrderDTO) FromModel(m *models.Order) *CurrentVendorOrderDTO {
	return &CurrentVendorOrderDTO{
		ID:        m.ID,
		Date:      m.CreatedAt.Format("2006-01-02"),
		User:      PublicUserDTO{}.FromModel(&m.Bookings[0].User),
		CourtType: m.Bookings[0].Court.CourtType.Type,
		Price:     m.Price,
		AppFee:    m.AppFee,
	}
}

// FromModels is a function that converts a slice of order models to
// a slice of order DTOs.
//
// m: The slice of booking models
//
// Returns a pointer to a slice of booking DTOs
func (c CurrentVendorOrderDTO) FromModels(m *[]models.Order) *[]CurrentVendorOrderDTO {
	// Create a slice of booking DTOs
	orders := []CurrentVendorOrderDTO{}

	// Iterate over the slice of booking models
	for _, order := range *m {
		orders = append(orders, *c.FromModel(&order))
	}

	return &orders
}
