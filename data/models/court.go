package models

import "time"

// Court is the model for the court table.
type Court struct {
	// ID is the primary key of the court.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// VendorID is the foreign key of the vendor.
	VendorID uint
	Vendor   Vendor `gorm:"foreignKey:VendorID"`

	// CourtTypeID is the foreign key of the court type.
	CourtTypeID uint
	CourtType   CourtType `gorm:"foreignKey:CourtTypeID"`

	// Name is the name of the court.
	Name string `gorm:"not null;type:varchar(255)"`

	// Price is the price of the court.
	Price float64 `gorm:"not null"`

	// CreatedAt is the time when the court was created.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt is the time when the court was updated.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
