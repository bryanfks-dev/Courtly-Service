package models

import "time"

// Advertisement is the model for the advertisement table.
type Advertisement struct {
	// ID is the primary key of the advertisement.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// Vendor is the foreign key of the vendor.
	VendorID uint   `gorm:"not null"`
	Vendor   Vendor `gorm:"foreignKey:VendorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// CourtType is the foreign key of the court type.
	CourtTypeID uint      `gorm:"not null"`
	CourtType   CourtType `gorm:"foreignKey:CourtTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// Image is the image of the advertisement.
	Image string `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}
