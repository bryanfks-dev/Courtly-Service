package dto

import "main/data/models"

// AdvertisementDTO is a data transfer object that represents
// the advertisement entity.
type AdvertisementsResponseDTO struct {
	Advertisements *[]AdvertisementDTO `json:"ads"`
}

func (a AdvertisementsResponseDTO) FromModels(m *[]models.Advertisement) *AdvertisementsResponseDTO {
	// Create a slice of advertisement DTOs
	dto := []AdvertisementDTO{}

	// Convert the advertisements to advertisement DTOs
	for _, ad := range *m {
		dto = append(dto, *AdvertisementDTO{}.FromModel(&ad))
	}

	return &AdvertisementsResponseDTO{
		Advertisements: &dto,
	}
}
