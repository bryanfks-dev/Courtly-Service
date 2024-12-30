package dto

import (
	"main/data/models"
	"main/internal/providers/midtrans"
)

// OrderDetailDTO is a struct that defines the OrderDetailDTO
type OrderDetailDTO struct {
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

	// PaymentToken is the payment token of the order
	PaymentToken *string `json:"payment_token"`

	// Status is the status of the order
	Status string `json:"status"`

	// Bookings is the bookings of the order
	Bookings *[]BookingDTO `json:"bookings"`
}

// FromModel is a method that converts a model to a DTO
//
// m: The order model
//
// Returns the DTO
func (o OrderDetailDTO) FromModel(m *models.Order) *OrderDetailDTO {
	// bookingDto is a placeholder for the booking DTO
	bookingDtos := []BookingDTO{}

	// Loop through the bookings
	for _, booking := range m.Bookings {
		// Append the booking DTO
		bookingDtos = append(bookingDtos, *BookingDTO{}.FromModel(&booking))
	}

	return &OrderDetailDTO{
		ID:              m.ID,
		MidtransOrderID: midtrans.CreateMidtransOrderId(m.ID),
		OrderDate:       m.Bookings[0].Date.Format("2006-01-02"),
		CreatedDate:     m.CreatedAt.Format("2006-01-02"),
		Price:           m.Price,
		AppFee:          m.AppFee,
		PaymentToken:    m.PaymentToken,
		Status:          m.Status,
		Bookings:        &bookingDtos,
	}
}
