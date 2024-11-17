package dto

import "main/data/models"

// VendorCourt is a struct that defines the VendorCourt DTO.
type VendorCourt struct {
	// ID is the court ID.
	ID uint `json:"id"`

	// Name is the court name.
	Name string `json:"name"`

	// Price is the court price.
	Price float64 `json:"price"`
}

// VendorCourtTypeResponse is a struct that defines the VendorCourtTypeResponse DTO.
type VendorCourtTypeResponse struct {
	// CourtType is the type of the court.
	CourtType string `json:"court_type"`

	// Courts is a list of Court DTO.
	Courts *[]VendorCourt `json:"courts"`
}

// FormModel creates a VendorCourts DTO from a Court model.
//
// m: The Court model.
//
// Returns a VendorCourts DTO.
func (v VendorCourt) FromModel(m models.Court) *VendorCourt {
	return &VendorCourt{
		ID:    m.ID,
		Name:  m.Name,
		Price: m.Price,
	}
}
