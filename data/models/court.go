package models

import "time"

// Court is the model for the court table.
type Court struct {
	// ID is the primary key of the court.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// VendorID is the foreign key of the vendor.
	VendorID uint   `gorm:"not null;index"`
	Vendor   Vendor `gorm:"foreignKey:VendorID"`

	// CourtTypeID is the foreign key of the court type.
	CourtTypeID uint      `gorm:"not null;index"`
	CourtType   CourtType `gorm:"foreignKey:CourtTypeID"`

	// Name is the name of the court.
	Name string `gorm:"not null;type:varchar(255)"`

	// Price is the price of the court.
	Price float64 `gorm:"not null"`

	// Image is the image of the court.
	Image string `gorm:"not null"`

	// CreatedAt is the time when the court was cwreated.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt is the time when the court was updated.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Reviews is the list of reviews that have the court.
	Reviews []Review `gorm:"foreignKey:CourtTypeID"`
}
