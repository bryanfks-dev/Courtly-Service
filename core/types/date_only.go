package types

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DateOnly is a struct that defines a date only type.
type DateOnly struct {
	time.Time
}

// GromDatType is a method that returns the GORM data type of the date only type.
//
// Returns The GORM data type of the date only type.
func (DateOnly) GormDataType() string {
	return "date"
}

// GormDBDataType is a method that returns the database data type of the date only type.
//
// db: The GORM database.
// field: The schema field.
//
// Returns The database data type of the date only type.
func (DateOnly) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "date"
}

// Value is a method that returns the value of the date only type.
//
// Returns The value of the date only type.
func (dateOnly DateOnly) Value() driver.Value {
	// If the date is zero, return nil
	if dateOnly.IsZero() {
		return nil
	}

	return dateOnly.Time.Format("Jan 02, 2006")
}
