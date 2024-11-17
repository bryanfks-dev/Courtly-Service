package dto

// VendorLoginForm is a struct that represents the login form
// that is sent by the vendor.
type VendorLoginForm struct {
	// Email is the email of the vendor.
	Email string `json:"email"`

	// Password is the password of the vendor.
	Password string `json:"password"`
}