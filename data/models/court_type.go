package models

// CourtType is a struct that represents the court type models.
type CourtType struct {
	// ID is the primary key of the court type.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// Type is the type of the court.
	Type string `gorm:"not null;unique;type:varchar(255);index"`
}
