package dto

import (
	"main/data/models"
	"main/internal/providers/midtrans"
)

// CurrentVendorOrderDetailDTO is a struct that defines the CurrentVendorOrderDetailDTO
type CurrentVendorOrderDetailDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

	// MidtransOrderID is the midtrans order ID of the order
	MidtransOrderID string `json:"midtrans_order_id"`

	// OrderDate is the date of the order
	OrderDate string `json:"order_date"`

	// CreatedDate is the created date of the order
	CreatedDate string `json:"created_date"`

	// Price is the price of the order
	Price float64 `json:"price"`

	// AppFee is the application fee of the order
	AppFee float64 `json:"app_fee"`

	// Bookings is the bookings of the order
	Bookings *[]CurrentVendorBookingDTO `json:"bookings"`
}

// FromModel is a method that converts a model to a DTO
//
// m: The order model
//
// Returns the DTO
func (c CurrentVendorOrderDetailDTO) FromModel(m *models.Order) *CurrentVendorOrderDetailDTO {
	// bookingDto is a placeholder for the booking DTO
	bookingDtos := []CurrentVendorBookingDTO{}

	// Loop through the bookings
	for _, booking := range m.Bookings {
		// Append the booking DTO
		bookingDtos = append(bookingDtos, *CurrentVendorBookingDTO{}.FromModel(&booking))
	}

	return &CurrentVendorOrderDetailDTO{
		ID:              m.ID,
		MidtransOrderID: midtrans.CreateMidtransOrderId(m.ID),
		OrderDate:       m.Bookings[0].Date.Format("2006-01-02"),
		CreatedDate:     m.CreatedAt.Format("2006-01-02"),
		Price:           m.Price,
		AppFee:          m.AppFee,
		Bookings:        &bookingDtos,
	}
}
