package dto

import (
	"main/data/models"
)

// CurrentVendor is a struct that represents the current vendor dto.
type CurrentVendor struct {
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
func (c CurrentVendor) FromModel(m *models.Vendor) *CurrentVendor {
	return &CurrentVendor{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Address:   m.Address,
		OpenTime:  m.OpenTime.Format("15:04"),
		CloseTime: m.CloseTime.Format("15:04"),
	}
}
