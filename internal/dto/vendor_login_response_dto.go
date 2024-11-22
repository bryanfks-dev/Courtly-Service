package dto

// VendorLoginResponseDTO is a type that represents the response data of the vendor login response.
type VendorLoginResponseDTO struct {
	// Vendor is the current vendor.
	Vendor *CurrentVendorDTO `json:"vendor"`

	// Token is the JWT token.
	Token string `json:"token"`
}
