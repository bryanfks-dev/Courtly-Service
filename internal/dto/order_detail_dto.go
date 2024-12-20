package dto

import "main/data/models"

// OrderDetailDTO is a struct that defines the OrderDetailDTO
type OrderDetailDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

	// Date is the date of the order
	Date string `json:"date"`

	// Price is the price of the order
	Price float64 `json:"price"`

	// AppFee is the application fee of the order
	AppFee float64 `json:"app_fee"`

	// Status is the status of the order
	Status string `json:"status"`

	// Bookings is the bookings of the order
	Bookings []BookingDTO `json:"bookings"`
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
		ID:       m.ID,
		Date:     m.Bookings[0].Date.Format("2006-01-02"),
		Price:    m.Price,
		AppFee:   m.AppFee,
		Status:   m.Status,
		Bookings: bookingDtos,
	}
}
