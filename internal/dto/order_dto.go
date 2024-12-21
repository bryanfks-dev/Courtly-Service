package dto

import (
	"main/core/enums"
	"main/data/models"
)

// OrderDTO is a struct that represents the order data transfer object.
type OrderDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

	// Bookings is the list of bookings
	Date string `json:"date"`

	// Vendor is the vendor of the order
	Vendor *PublicVendorDTO `json:"vendor"`

	// CourtType is the court type of the order
	CourtType string `json:"court_type"`

	// PaymentMethod is the payment method of the order
	PaymentMethod string `json:"payment_method"`

	// Price is the price of the order
	Price float64 `json:"price"`

	// AppFee is the application fee of the order
	AppFee float64 `json:"app_fee"`

	// Status is the status of the order
	Status string `json:"status"`

	// Reviewed is the review status of the order
	Reviewed *bool `json:"reviewed,omitempty"`
}

// FromModel is a function that converts an order model to an order DTO.
//
// m: The order model.
// reviewed: The review status of the order.
//
// Returns the order DTO.
func (o OrderDTO) FromModel(m *models.Order, reviewed *bool) *OrderDTO {
	// Check if the order status is pending
	if m.Status == enums.Pending.Label() {
		reviewed = nil
	}

	return &OrderDTO{
		ID:            m.ID,
		Date:          m.CreatedAt.Format("2006-01-02"),
		Vendor:        PublicVendorDTO{}.FromModel(&m.Bookings[0].Vendor),
		CourtType:     m.Bookings[0].Court.CourtType.Type,
		PaymentMethod: m.PaymentMethod.Method,
		Price:         m.Price,
		AppFee:        m.AppFee,
		Status:        m.Status,
		Reviewed:      reviewed,
	}
}
