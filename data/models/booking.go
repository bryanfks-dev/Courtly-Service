package models

import (
	"main/core/shared"
)

// Booking is the model for the booking table.
type Booking struct {
	// ID is the primary key of the book.
	ID uint `gorm:"primary_key;autoIncrement"`

	// OrderID is the foreign key of the order.
	OrderID uint  `gorm:"not null"`
	Order   Order `gorm:"foreignKey:OrderID"`

	// UserID is the foreign key of the user.
	UserID uint `gorm:"not null;index"`
	User   User `gorm:"foreignKey:UserID"`

	// VendorID is the foreign key of the vendor.
	VendorID uint   `gorm:"not null;index"`
	Vendor   Vendor `gorm:"foreignKey:VendorID"`

	// CourtID is the foreign key of the court.
	CourtID uint  `gorm:"not null;index;constraint:OnDelete:SET NULL"`
	Court   Court `gorm:"foreignKey:CourtID"`

	// Date is the date of the book was created.
	Date shared.DateOnly `gorm:"autoCreateTime;type:DATE"`

	//	BookStartTime is the start time of the book.
	BookStartTime shared.TimeOnly `gorm:"not null"`

	// BookEndTime is the end time of the book.
	BookEndTime shared.TimeOnly `gorm:"not null"`
}
