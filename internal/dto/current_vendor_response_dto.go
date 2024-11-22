package dto

// CurrentVendorResponseDTO is a struct that represents the response data for the current vendor.
type CurrentVendorResponseDTO struct {
	// Vendor is the current vendor.
	Vendor *CurrentVendorDTO `json:"vendor"`
}
