package dto

import (
	"fmt"
	"main/core/config"
	"main/core/types"
	"main/data/models"
	"main/delivery/http/router"
)

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

	// Rating is the rating of the court.
	Rating float64 `json:"rating,omitempty"`

	// ImageUrl is the image URL of the court.
	ImageUrl string `json:"image_url"`
}

// FromModel is a function that converts a court model to a court DTO.
//
// m: The court model.
//
// Returns the court DTO.
func (c CourtDTO) FromModel(m *models.Court) *CourtDTO {
	// courtImagePath is the path to the court image.
	courtImagePath := fmt.Sprintf("%s:%d%s/%s", config.ServerConfig.Host, config.ServerConfig.Port, router.CourtImages, m.Image)

	return &CourtDTO{
		ID:       m.ID,
		Name:     m.Name,
		Vendor:   c.Vendor.FromModel(&m.Vendor),
		Type:     m.CourtType.Type,
		Price:    m.Price,
		ImageUrl: courtImagePath,
	}
}

// FromCourtMap is a function that converts a court map to a court DTO.
//
// m: The court map.
//
// Returns the court DTO.
func (c CourtDTO) FromCourtMap(m *types.CourtMap) *CourtDTO {
	// Get the court
	court := m.GetCourt()

	// courtImagePath is the path to the court image.
	courtImagePath := fmt.Sprintf("%s:%d%s/%s", config.ServerConfig.Host, config.ServerConfig.Port, router.CourtImages, court.Image)

	return &CourtDTO{
		ID:       court.ID,
		Name:     court.Name,
		Vendor:   PublicVendorDTO{}.FromModel(&court.Vendor),
		Type:     court.CourtType.Type,
		Price:    court.Price,
		ImageUrl: courtImagePath,
		Rating:   m.GetTotalRating(),
	}
}
