package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
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
// Returns The value of the time only type and an error if any.
func (timeOnly TimeOnly) Value() (driver.Value, error) {
	// If the time is zero, return nil
	if timeOnly.IsZero() {
		return nil, nil
	}

	return timeOnly.GetTime().Format("15:04"), nil
}

// GetTime is a method that returns the time of the time only type.
//
// Returns The time of the time only type.
func (timeOnly TimeOnly) GetTime() time.Time {
	return timeOnly.Time
}

// Scan is a method that scans the time only type.
//
// value: The value to scan.
//
// Returns An error if any.
func (timeOnly *TimeOnly) Scan(value any) error {
	scanned, ok := value.([]byte)

	// If the value is not a byte slice, return an error
	if !ok {
		return errors.New(fmt.Sprintln("Failed to scan Time value: ", value))
	}

	// Parse the scanned string
	scannedTime, err := time.Parse("15:04:05", string(scanned))

	// If there is an error parsing the time, return the error
	if err != nil {
		return err
	}

	// Set the time
	*timeOnly = TimeOnly{Time: scannedTime}

	return nil
}
