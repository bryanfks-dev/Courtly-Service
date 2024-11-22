package dto

// UserRegisterForm is a struct that represents the register form
// that is sent by the user.
type UserRegisterForm struct {
	// Username is the username of the user.
	Username string `json:"username"`

	// Email is the email of the user.
	PhoneNumber string `json:"phone_number"`

	// Password is the password of the user.
	Password string `json:"password"`

	// ConfirmPassword is the confirmation password of the user.
	ConfirmPassword string `json:"confirm_password"`
}

// VendorRegisterForm is a struct that represents the register form
type VendorRegisterForm struct {
	// Name is the name of the vendor.
	Name string

	// Address is the address of the vendor.
	Address string

	// Email is the email of the vendor.
	Email string

	// Password is the password of the vendor.
	Password string
}
