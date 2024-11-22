package types

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// TimeOnly is a struct that represents a time only type.
type TimeOnly struct {
	time.Time
}

// GormDataType is a method that returns the GORM data type of the time only type.
//
// Returns The GORM data type of the time only type.
func (TimeOnly) GormDataType() string {
	return "time"
}

// GormDBDataType is a method that returns the database data type of the time only type.
//
// db: The GORM database.
// field: The schema field.
//
// Returns The database data type of the time only type.
func (TimeOnly) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}

// Value is a method that returns the value of the time only type.
//
// Returns The value of the time only type.
func (timeOnly TimeOnly) Value() driver.Value {
	// If the time is zero, return nil
	if timeOnly.IsZero() {
		return nil
	}
	
	return timeOnly.Time.Format("15:04")
}
