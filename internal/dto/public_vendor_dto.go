package dto

import (
	"main/data/models"
)

// PublicVendorDTO is a struct that defines the vendor data transfer object.
type PublicVendorDTO struct {
	// ID is the primary key of the vendor.
	ID uint `json:"id"`

	// Name is the name of the vendor.
	Name string `json:"name"`

	// Address is the address of the vendor.
	Address string `json:"address"`

	// OpenTime is the opening time of the vendor.
	OpenTime string `json:"open_time"`

	// CloseTime is the closing time of the vendor.
	CloseTime string `json:"close_time"`
}

// FromModel is a function that converts a vendor model to a public vendor DTO.
//
// m: The vendor model.
//
// Returns the public vendor DTO.
func (v PublicVendorDTO) FromModel(m *models.Vendor) *PublicVendorDTO {
	// Get the open and close time
	openTime, _ := m.OpenTime.Value()
	closeTime, _ := m.CloseTime.Value()

	return &PublicVendorDTO{
		ID:        v.ID,
		Name:      v.Name,
		Address:   v.Address,
		OpenTime:  openTime.(string),
		CloseTime: closeTime.(string),
	}
}
