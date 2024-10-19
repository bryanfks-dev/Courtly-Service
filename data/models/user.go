package models

import (
	"time"
)

// User is a struct that represents the user models.
type User struct {
	// ID is the primary key of the user.
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`

	// Username is the username of the user.
	Username string `json:"username" gorm:"not null;unique;type:varchar(255)"`

	// PhoneNumber is the phone number of the user.
	PhoneNumber string `json:"phone_number" gorm:"not null;unique;type:varchar(20)"`

	// Password is the password of the user.
	Password string `json:"-" gorm:"not null;type:varchar(255)"`

	// CreatedAt is the time when the user was created.
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`

	// UpdatedAt is the time when the user was updated.
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
