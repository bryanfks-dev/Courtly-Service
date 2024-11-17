package models

import (
	"time"
)

// User is a struct that represents the user models.
type User struct {
	// ID is the primary key of the user.
	ID uint `gorm:"primaryKey;autoIncrement"`

	// Username is the username of the user.
	Username string `gorm:"not null;unique;type:varchar(255);index"`

	// PhoneNumber is the phone number of the user.
	PhoneNumber string `gorm:"not null;unique;type:varchar(20)"`

	// ProfilePicture is the profile picture of the user.
	ProfilePicture string `gorm:"type:varchar(255)"`

	// Password is the password of the user.
	Password string `gorm:"not null;type:varchar(255)"`

	// CreatedAt is the time when the user was created.
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// UpdatedAt is the time when the user was updated.
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
