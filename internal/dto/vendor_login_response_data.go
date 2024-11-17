package dto

// VendorLoginResponseData is a type that represents the response data of the vendor login response.
type VendorLoginResponseData struct {
	// Vendor is the current vendor.
	Vendor *CurrentVendor `json:"vendor"`

	// Token is the JWT token.
	Token string `json:"token"`
}
