package models

import (
	"main/core/types"
)

// Review is the model for the review table.
type Review struct {
	// ID is the primary key of the review.
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

	// Rating is the rating of the review.
	Rating int8 `gorm:"not null"`

	// Review is the review of the review.
	Review string `gorm:"not null;type:text"`

	// Date is the date of the review was created.
	Date types.DateOnly `gorm:"type:DATEautoCreateTime"`
}
