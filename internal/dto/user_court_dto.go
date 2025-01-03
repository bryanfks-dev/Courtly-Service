package dto

import (
	"fmt"
	"main/core/types"
	"main/data/models"
	"main/delivery/http/router"
)

// UserCourtDTO is a struct that defines the court data transfer object
// for user client type.
type UserCourtDTO struct {
	// ID is the primary key of the court.
	ID uint `json:"id"`

	// Name is the name of the court.
	Name string `json:"name"`

	// Vendor is the vendor of the court.
	Vendor *VendorDTO `json:"vendor"`

	// CourtType is the type of the court.
	Type string `json:"type"`

	// Name is the name of the court.
	Price float64 `json:"price"`

	// Rating is the rating of the court.
	Rating *float64 `json:"rating,omitempty"`

	// ImageUrl is the image URL of the court.
	ImageUrl string `json:"image_url"`
}

// FromModel is a function that converts a court model to a court DTO.
//
// m: The court model.
//
// Returns the user court DTO.
func (c UserCourtDTO) FromModel(m *models.Court) *UserCourtDTO {
	// courtImagePath is the path to the court image.
	courtImagePath := fmt.Sprintf("%s/%s", router.CourtImages, m.Image)

	return &UserCourtDTO{
		ID:       m.ID,
		Name:     m.Name,
		Vendor:   VendorDTO{}.FromModel(&m.Vendor),
		Type:     m.CourtType.Type,
		Price:    m.Price,
		Rating:   nil,
		ImageUrl: courtImagePath,
	}
}

// FromCourtMap is a function that converts a court map to a court DTO.
//
// m: The court map.
//
// Returns the user court DTO.
func (c UserCourtDTO) FromCourtMap(m *types.CourtMap) *UserCourtDTO {
	// Get the court
	court := m.GetCourt()

	// courtImagePath is the path to the court image.
	courtImagePath := fmt.Sprintf("%s/%s", router.CourtImages, court.Image)

	// Get the rating
	rating := m.GetTotalRating()

	return &UserCourtDTO{
		ID:       court.ID,
		Name:     court.Name,
		Vendor:   VendorDTO{}.FromModel(&court.Vendor),
		Type:     court.CourtType.Type,
		Price:    court.Price,
		ImageUrl: courtImagePath,
		Rating:   &rating,
	}
}
