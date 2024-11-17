package dto

// CurrentVendorResponseData is a struct that represents the response data for the current vendor.
type CurrentVendorResponseData struct {
	// Vendor is the current vendor.
	Vendor *CurrentVendor `json:"vendor"`
}
