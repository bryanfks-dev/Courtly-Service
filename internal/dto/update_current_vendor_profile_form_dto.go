package dto

// UpdateCurrentVendorProfileDTO is a struct that defines the data
// needed to update the current vendor profile.
type UpdateCurrentVendorProfileDTO struct {
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}
