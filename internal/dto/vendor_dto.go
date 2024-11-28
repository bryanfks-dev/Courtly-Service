package dto

import "main/data/models"

// VendorDTO is a struct that represents the vendor data transfer object.
type VendorDTO struct {
	// ID is the ID of the vendor
	ID uint `json:"id"`

	// Name is the name of the vendor
	Name string `json:"name"`

	// Address is the address of the vendor
	Address string `json:"address"`

	// OpenTime is the open time of the vendor
	OpenTime string `json:"open_time"`

	// CloseTime is the close time of the vendor
	CloseTime string `json:"close_time"`
}

// FromModel creates a CurrentVendor DTO from a Vendor model.
//
// m: The vendor model.
//
// Returns a CurrentVendor DTO.
func (c VendorDTO) FromModel(m *models.Vendor) *VendorDTO {
	// Get the open and close time
	openTime, _ := m.OpenTime.Value()
	closeTime, _ := m.CloseTime.Value()

	return &VendorDTO{
		ID:        m.ID,
		Name:      m.Name,
		Address:   m.Address,
		OpenTime:  openTime.(string),
		CloseTime: closeTime.(string),
	}
}
