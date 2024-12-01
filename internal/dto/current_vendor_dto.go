package dto

import (
	"main/data/models"
)

// CurrentVendorDTO is a struct that represents the current vendor dto.
type CurrentVendorDTO struct {
	// ID is the primary key of the vendor.
	ID uint `json:"id"`

	// Name is the name of the vendor.
	Name string `json:"name"`

	// Username is the username of the vendor.
	Email string `json:"email"`

	// Username is the username of the vendor.
	Address string `json:"address"`

	// OpenTime is the open time of the vendor.
	OpenTime string `json:"open_time"`

	// CloseTime is the close time of the vendor.
	CloseTime string `json:"close_time"`
}

// FromModel creates a CurrentVendor DTO from a Vendor model.
//
// m: The vendor model.
//
// Returns a CurrentVendor DTO.
func (c CurrentVendorDTO) FromModel(m *models.Vendor) *CurrentVendorDTO {
	return &CurrentVendorDTO{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Address:   m.Address,
		OpenTime:  m.OpenTime.String(),
		CloseTime: m.CloseTime.String(),
	}
}
