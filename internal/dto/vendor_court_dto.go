package dto

import "main/data/models"

// VendorCourtDTO is a struct that defines the VendorCourt DTO.
type VendorCourtDTO struct {
	// ID is the court ID.
	ID uint `json:"id"`

	// Name is the court name.
	Name string `json:"name"`

	// Price is the court price.
	Price float64 `json:"price"`
}

// FormModel creates a VendorCourts DTO from a Court model.
//
// m: The Court model.
//
// Returns a VendorCourts DTO.
func (v VendorCourtDTO) FromModel(m *models.Court) *VendorCourtDTO {
	return &VendorCourtDTO{
		ID:    m.ID,
		Name:  m.Name,
		Price: m.Price,
	}
}
