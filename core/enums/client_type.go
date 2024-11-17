package enums

import "main/data/models"

// UserRole is an enum that defines the user roles.
type ClientType int

const (
	User ClientType = iota
	Vendor
)

// String is a function that returns the string representation of the user role.
//
// Returns a string containing the user role.
func (c ClientType) String() string {
	return [...]string{"User", "Vendor"}[c]
}

// Model is a function that returns the model of the client type.
//
// 
func (c ClientType) Model() any {
	return [...]any{models.User{}, models.Vendor{}}[c]
}
