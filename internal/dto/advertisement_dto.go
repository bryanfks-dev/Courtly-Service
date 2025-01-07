package dto

import (
	"fmt"
	"main/data/models"
	"main/delivery/http/router"
)

// AdvertisementDTO is a data transfer object that represents the ad entity.
type AdvertisementDTO struct {
	// ID is the unique identifier of the ad.
	ID uint `json:"id"`

	// ID is the unique identifier of the ad.
	ImageUrl string `json:"image_url"`

	// Vendor is the vendor of the ad.
	Vendor *VendorDTO `json:"vendor"`

	// CourtType is the type of the court.
	CourtType string `json:"court_type"`
}

// FromModel is a function that converts an ad model to an ad DTO.
//
// m: The advertisement model.
//
// Returns an ad DTO.
func (a AdvertisementDTO) FromModel(m *models.Advertisement) *AdvertisementDTO {
	// adImagePath is the path to the ad image.
	adImagePath := fmt.Sprintf("%s/%s", router.UserProfiles, m.Image)

	return &AdvertisementDTO{
		ID:        m.ID,
		ImageUrl:  adImagePath,
		CourtType: m.CourtType.Type,
		Vendor:    VendorDTO{}.FromModel(&m.Vendor),
	}
}
