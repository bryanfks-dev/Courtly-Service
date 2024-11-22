package models

// Order is the model for the order table.
type Order struct {
	// ID is the primary key of the order.
	ID uint `gorm:"primary_key;autoIncrement"`

	// BookID is the foreign key of the book.
	BookID uint
	Book   Book `gorm:"foreignKey:BookID"`

	// PaymentMethodID is the foreign key of the payment method.
	PaymentMethodID uint
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`

	// Price is the price of the order.
	Price float64 `gorm:"not null"`

	// AppFee is the app fee of the order.
	AppFee float64 `gorm:"not null"`
}
