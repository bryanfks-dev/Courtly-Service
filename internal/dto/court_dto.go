package dto

import "main/data/models"

// CourtDTO is a struct that defines the court data transfer object.
type CourtDTO struct {
	// ID is the primary key of the court.
	ID uint `json:"id"`

	// Name is the name of the court.
	Name string `json:"name"`

	// Vendor is the vendor of the court.
	Vendor *PublicVendorDTO `json:"vendor"`

	// CourtType is the type of the court.
	Type string `json:"court_type"`

	// Name is the name of the court.
	Price float64 `json:"price"`
}

// FromModel is a function that converts a court model to a court DTO.
//
// m: The court model.
//
// Returns the court DTO.
func (c CourtDTO) FromModel(m *models.Court) *CourtDTO {
	return &CourtDTO{
		ID:     m.ID,
		Name:   m.Name,
		Vendor: c.Vendor.FromModel(&m.Vendor),
		Type:   m.CourtType.Type,
		Price:  m.Price,
	}
}
