package dto

import (
	"main/core/enums"
	"main/data/models"
)

// CurrentUserOrderDTO is a struct that represents the order data transfer object
// for user client type.
type CurrentUserOrderDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

	// Date is the date of the order was created
	Date string `json:"date"`

	// Vendor is the vendor of the order
	Vendor *VendorDTO `json:"vendor"`

	// CourtType is the court type of the order
	CourtType string `json:"court_type"`

	// Price is the price of the order
	Price float64 `json:"price"`

	// AppFee is the application fee of the order
	AppFee float64 `json:"app_fee"`

	// PaymentToken is the payment token of the order
	PaymentToken *string `json:"payment_token"`

	// Status is the status of the order
	Status string `json:"status"`

	// Reviewed is the review status of the order
	Reviewed *bool `json:"reviewed,omitempty"`
}

// FromModel is a function that converts an order model to an current
// user order DTO.
//
// m: The order model.
// reviewed: The review status of the order.
//
// Returns the current user order DTO.
func (o CurrentUserOrderDTO) FromModel(m *models.Order, reviewed *bool) *CurrentUserOrderDTO {
	// Check if the order status is pending
	if m.Status == enums.Pending.Label() {
		reviewed = nil
	}

	return &CurrentUserOrderDTO{
		ID:           m.ID,
		Date:         m.CreatedAt.Format("2006-01-02"),
		Vendor:       VendorDTO{}.FromModel(&m.Bookings[0].Vendor),
		CourtType:    m.Bookings[0].Court.CourtType.Type,
		Price:        m.Price,
		AppFee:       m.AppFee,
		PaymentToken: m.PaymentToken,
		Status:       m.Status,
		Reviewed:     reviewed,
	}
}
