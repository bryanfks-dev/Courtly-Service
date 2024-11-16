package models

// PaymentMethods is a struct that defines the payment methods model.
type PaymentMethod struct {
	// ID is the primary key of the payment method.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// Method is the name of the payment method.
	Method string `gorm:"not null;unique;type:varchar(255);index"`
}
