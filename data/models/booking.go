package models

import (
	"main/core/types"
)

// Booking is the model for the booking table.
type Booking struct {
	// ID is the primary key of the book.
	ID uint `gorm:"primary_key;autoIncrement"`

	// OrderID is the foreign key of the order.
	OrderId uint
	Order   Order `gorm:"foreignKey:OrderId"`

	// UserID is the foreign key of the user.
	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	// VendorID is the foreign key of the vendor.
	VendorID uint
	Vendor   Vendor `gorm:"foreignKey:VendorID"`

	// CourtTypeID is the foreign key of the court type.
	CourtTypeID uint
	CourtType   CourtType `gorm:"foreignKey:CourtTypeID"`

	// Date is the date of the book was created.
	Date types.DateOnly `gorm:"autoCreateTime;type:DATE"`

	//	BookStartTime is the start time of the book.
	BookStartTime types.TimeOnly `gorm:"not null"`

	// BookEndTime is the end time of the book.
	BookEndTime types.TimeOnly `gorm:"not null"`
}
