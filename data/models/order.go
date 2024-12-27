package models

import "time"

// Order is the model for the order table.
type Order struct {
	// ID is the primary key of the order.
	ID uint `gorm:"primary_key;autoIncrement"`

	// PaymentMethodID is the foreign key of the payment method.
	PaymentMethodID uint         `gorm:"default:null"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	// Price is the price of the order.
	Price float64 `gorm:"not null"`

	// AppFee is the app fee of the order.
	AppFee float64 `gorm:"not null"`

	// PaymentToken is the payment token of the order.
	PaymentToken *string `gorm:"default:null"`

	// Status is the status of the order.
	Status string `gorm:"not null"`

	// Bookings is the list of bookings.
	Bookings []Booking `gorm:"foreignKey:OrderID"`

	// CreatedAt is the time the order was created.
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
