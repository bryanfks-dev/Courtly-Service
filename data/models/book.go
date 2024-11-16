package models

import "time"

// Book is the model for the book table.
type Book struct {
	// ID is the primary key of the book.
	ID uint `gorm:"primary_key;autoIncrement"`

	// UserID is the foreign key of the user.
	UserID uint
	User   User

	// VendorID is the foreign key of the vendor.
	VendorID uint
	Vendor   Vendor

	// CourtTypeID is the foreign key of the court type.
	CourtTypeID uint
	CourtType   CourtType

	// Date is the date of the book was created.
	Date time.Time `gorm:"autoCreateTime"`

	//	BookStartTime is the start time of the book.
	BookStartTime time.Time `gorm:"not null"`

	// BookEndTime is the end time of the book.
	BookEndTime   time.Time `gorm:"not null"`
}