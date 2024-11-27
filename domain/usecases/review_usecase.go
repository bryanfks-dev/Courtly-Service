package usecases

import (
	"log"
	"main/data/models"
	"main/internal/repository"
)

// ReviewUseCase is a struct that defines the review use case.
type ReviewUseCase struct {
	ReviewRepository *repository.ReviewRepository
}

// NewReviewUseCase is a factory function that returns a new instance of the ReviewUseCase.
//
// r: The review repository.
//
// Returns a new instance of the ReviewUseCase.
func NewReviewUseCase(r *repository.ReviewRepository) *ReviewUseCase {
	return &ReviewUseCase{
		ReviewRepository: r,
	}
}

// GetCourtReviewsUsingIDType is a use case that handles the request to
// get the reviews of a court using the court id and type.
//
// vendorID: The id of the vendor.
// courtType: The type of the court.
//
// Returns a slice of maps containing the reviews of the court.
func (r *ReviewUseCase) GetCourtReviewsUsingIDType(vendorID uint, courtType string) (*[]models.Review, error) {
	// Get the reviews using the vendor ID and court type
	reviews, err := r.ReviewRepository.GetUsingVendorIDCourtType(vendorID, courtType)

	// Check if there is an error
	if err != nil {
		log.Println("Error getting reviews using vendor ID and court type:", err)

		return nil, err
	}

	// Return the reviews and an error if any
	return reviews, err
}
