package dto

import "main/data/models"

// CurrentVendorCourtDTO is a struct that defines the current vendor court data transfer object.
type CurrentVendorCourtDTO struct {
	// ID is the primary key of the court.
	ID uint `json:"id"`

	// Name is the name of the court.
	Name string `json:"name"`

	// CourtType is the type of the court.
	Type string `json:"type"`

	// Name is the name of the court.
	Price float64 `json:"price"`
}

// FromModel is a function that converts a court model to a current vendor court DTO.
//
// m: The court model.
//
// Returns the current vendor court DTO.
func (c CurrentVendorCourtDTO) FromModel(m *models.Court) *CurrentVendorCourtDTO {
	return &CurrentVendorCourtDTO{
		ID:    m.ID,
		Name:  m.Name,
		Type:  m.CourtType.Type,
		Price: m.Price,
	}
}
