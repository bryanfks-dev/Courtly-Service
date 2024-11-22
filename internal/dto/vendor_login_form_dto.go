package dto

// VendorLoginFormDTO is a struct that represents the login form
// that is sent by the vendor.
type VendorLoginFormDTO struct {
	// Email is the email of the vendor.
	Email string `json:"email"`

	// Password is the password of the vendor.
	Password string `json:"password"`
}