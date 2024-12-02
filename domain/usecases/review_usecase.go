package usecases

import (
	"log"
	"main/data/models"
	"main/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

// ReviewUseCase is a struct that defines the review use case.
type ReviewUseCase struct {
	AuthUseCase      *AuthUseCase
	ReviewRepository *repository.ReviewRepository
}

// NewReviewUseCase is a factory function that returns a new instance of the ReviewUseCase.
//
// r: The review repository.
//
// Returns a new instance of the ReviewUseCase.
func NewReviewUseCase(a *AuthUseCase, r *repository.ReviewRepository) *ReviewUseCase {
	return &ReviewUseCase{
		AuthUseCase:      a,
		ReviewRepository: r,
	}
}

// GetCourtReviewsUsingIDCourtType is a use case that handles the request to
// get the reviews of a court using the vendor id and court type.
//
// vendorID: The id of the vendor.
// courtType: The type of the court.
//
// Returns a slice of maps containing the reviews of the court.
func (r *ReviewUseCase) GetCourtReviewsUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Review, error) {
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

// GetCurrentVendorReviews is a use case that handles the request to get the current vendor's reviews.
//
// token: The JWT token.
//
// Returns a slice of current vendor's reviews and erros if any.
func (r *ReviewUseCase) GetCurrentVendorReviews(token *jwt.Token) (*[]models.Review, error) {
	// Get the vendor ID from the token
	claims := r.AuthUseCase.DecodeToken(token)

	// Get the reviews using the vendor ID
	reviews, err := r.ReviewRepository.GetUsingVendorID(uint(claims.Id))

	// Check if there is an error
	if err != nil {
		log.Println("Error getting reviews using vendor ID:", err)

		return nil, err
	}

	return reviews, err
}
