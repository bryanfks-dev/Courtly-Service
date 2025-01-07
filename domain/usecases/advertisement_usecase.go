package usecases

import (
	"main/data/models"
	"main/internal/repository"
)

// AdvertisementUseCase is the use case for the advertisement.
type AdvertisementUseCase struct {
	AdvertisementsRepository *repository.AdvertisementRepository
}

// NewAdvertisementUseCase is a constructor for the AdvertisementUseCase.
//
// a: The advertisement repository.
//
// Returns the AdvertisementUseCase instance.F
func NewAdvertisementUseCase(a *repository.AdvertisementRepository) *AdvertisementUseCase {
	return &AdvertisementUseCase{
		AdvertisementsRepository: a,
	}
}

// GetAdvertisements is a use case function to get all advertisements.
//
// Returns the advertisements and an error if any.
func (a *AdvertisementUseCase) GetAdvertisements() (*[]models.Advertisement, error) {
	return a.AdvertisementsRepository.GetAll()
}
