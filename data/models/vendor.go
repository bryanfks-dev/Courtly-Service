package models

import (
	"main/core/shared"
	"time"
)

// Vendor is the model for the vendor table.
type Vendor struct {
	// ID is the primary key of the vendor.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// Name is the name of the vendor.
	Name string `gorm:"not null;unique;type:varchar(255)"`

	// Address is the address of the vendor.
	Address string `gorm:"not null;type:varchar(255)"`

	// Email is the email of the vendor.
	Email string `gorm:"not null;unique;type:varchar(255);index"`

	// PhoneNumber is the phone number of the vendor.
	Password string `gorm:"not null;type:varchar(255)"`

	// OpenTime is the opening time of the vendor.
	OpenTime shared.TimeOnly `gorm:"not null"`

	// CloseTime is the closing time of the vendor.
	CloseTime shared.TimeOnly `gorm:"not null"`

	// CreatedAt is the time when the user was created.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt is the time when the user was updated.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Bookings is the list of bookings that have the vendor.
	Bookings []Booking `gorm:"foreignKey:VendorID"`

	// Courts is the list of courts that have the vendor.
	Courts []Court `gorm:"foreignKey:VendorID"`

	// Reviews is the list of reviews that have the vendor.
	Reviews []Review `gorm:"foreignKey:VendorID"`
}
