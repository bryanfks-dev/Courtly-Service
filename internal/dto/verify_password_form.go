package dto

// VerifyPasswordForm is a data transfer object that represents the form
type VerifyPasswordForm struct {
	// Password is the password of the user
	Password string `json:"password"`
}
