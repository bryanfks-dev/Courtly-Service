package dto

import "main/data/models"

// OrderDTO is a struct that represents the order data transfer object.
type OrderDTO struct {
	// ID is the ID of the order
	ID uint `json:"id"`

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
	return &OrderDTO{
		ID:            m.ID,
		PaymentMethod: m.PaymentMethod.Method,
		Price:         m.Price,
		AppFee:        m.AppFee,
		Status:        m.Status,
	}
}
