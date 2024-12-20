package dto

import "main/data/models"

// OrderDTO is a struct that represents the order data transfer object.
type OrderDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

	// Bookings is the list of bookings
	Date string `json:"date"`

	// CourtType is the type of the court
	CourtType string `json:"court_type"`

	// VendorName is the vendor name of the order
	VendorName string `json:"vendor_name"`

	// PaymentMethod is the payment method of the order
	PaymentMethod string `json:"payment_method"`

	// Price is the price of the order
	Price float64 `json:"price"`

	// AppFee is the application fee of the order
	AppFee float64 `json:"app_fee"`

	// Status is the status of the order
	Status string `json:"status"`
}

// FromModel is a function that converts an order model to an order DTO.
//
// m: The order model.
//
// Returns the order DTO.
func (o OrderDTO) FromModel(m *models.Order) *OrderDTO {
	// Get the date
	date, _ := m.Bookings[0].Date.Value()

	return &OrderDTO{
		ID:            m.ID,
		Date:          date.(string),
		CourtType:     m.Bookings[0].Court.CourtType.Type,
		VendorName:    m.Bookings[0].Vendor.Name,
		PaymentMethod: m.PaymentMethod.Method,
		Price:         m.Price,
		AppFee:        m.AppFee,
		Status:        m.Status,
	}
}
