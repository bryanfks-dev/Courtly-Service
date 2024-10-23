package models

import "time"

// BlacklistedToken is a struct that represents the blacklisted token models.
type BlacklistedToken struct {
	// ID is the primary key of the user.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// Token is the token that is blacklisted.
	Token string `gorm:"not null;unique;type:varchar(500)"`

	// ExpiresAt is the time when the token expires.
	ExpiresAt time.Time `gorm:"not null"`
}
