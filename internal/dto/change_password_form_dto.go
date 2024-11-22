package dto

// ChangePasswordFormDTO is a data transfer object that represents the form data for changing the password.
type ChangePasswordFormDTO struct {
	// OldPassword is the current password of the user.
	OldPassword string `json:"old_password"`

	// NewPassword is the new password of the user.
	NewPassword string `json:"new_password"`

	// ConfirmPassword is the confirmation of the new password.
	ConfirmPassword string `json:"confirm_password"`
}
