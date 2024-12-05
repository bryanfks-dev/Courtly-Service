package usecases

import (
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// ReviewUseCase is a struct that defines the review use case.
type ReviewUseCase struct {
	AuthUseCase       *AuthUseCase
	ReviewRepository  *repository.ReviewRepository
	BookingRepository *repository.BookingRepository
	CourtRepository   *repository.CourtRepository
}

// NewReviewUseCase is a factory function that returns a new instance of the ReviewUseCase.
//
// a: The auth use case.
// r: The review repository.
// b: The booking repository.
// c: The court repository.
//
// Returns a new instance of the ReviewUseCase.
func NewReviewUseCase(a *AuthUseCase, r *repository.ReviewRepository, b *repository.BookingRepository, c *repository.CourtRepository) *ReviewUseCase {
	return &ReviewUseCase{
		AuthUseCase:       a,
		ReviewRepository:  r,
		BookingRepository: b,
		CourtRepository:   c,
	}
}

// GetCurrentVendorReviewCount is a use case that handles the request to get the current vendor's review count.
//
// token: The JWT token.
//
// Returns the review count and an error if any.
func (r *ReviewUseCase) GetCurrentVendorReviewCount(token *jwt.Token) (int64, error) {
	// Get the vendor ID from the token
	claims := r.AuthUseCase.DecodeToken(token)

	// Get the review count using the vendor ID
	return r.ReviewRepository.GetCountUsingVendorID(claims.Id)
}

// GetCurrentVendorStarCounts is a use case that handles the request to get the current vendor's star counts.
//
// token: The JWT token.
//
// Returns the star counts and an error if any.
func (r *ReviewUseCase) GetCurrentVendorStarCounts(token *jwt.Token) (*entities.ReviewStarsCount, error) {
	// Get the vendor ID from the token
	claims := r.AuthUseCase.DecodeToken(token)

	// Get the star counts using the vendor ID
	return r.ReviewRepository.GetStarCountsUsingVendorID(claims.Id)
}

// CalculateTotalRating is a function that calculates the total rating of the reviews.
//
// starCount: The star count of the reviews.
// reviewCount: The total number of reviews.
//
// Returns the total rating.
func (r *ReviewUseCase) CalculateTotalRating(starCount *entities.ReviewStarsCount, reviewCount int64) float64 {
	// Check if there are no reviews
	if reviewCount == 0 {
		return 0
	}

	// Formula to calculate the total rating:
	// (1 * OneStar + 2 * TwoStars + 3 * ThreeStars + 4 * FourStars + 5 * FiveStars)
	// -----------------------------------------------------------------------------
	//                           Total Reviews

	return (float64(starCount.OneStar) + float64(2*starCount.TwoStars) + float64(3*starCount.ThreeStars) + float64(4*starCount.FourStars) + float64(5*starCount.FiveStars)) / float64(reviewCount)
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
	return r.ReviewRepository.GetUsingVendorID(claims.Id)
}

// GetReviewCountUsingVendorIDCourtType is a use case that handles the request to get the
// review count using the vendor ID and court type.
//
// vendorID: The id of the vendor.
// courtType: The type of the court.
//
// Returns the review count and an error if any.
func (r *ReviewUseCase) GetReviewCountUsingVendorIDCourtType(vendorID uint, courtType string) (int64, error) {
	// Get the review count using the vendor ID and court type
	return r.ReviewRepository.GetCountUsingVendorIDCourtType(vendorID, courtType)
}

// GetStarCountsUsingVendorIDCourtType is a use case that handles the request to get the
// star counts using the vendor ID and court type.
//
// vendorID: The id of the vendor.
// courtType: The type of the court.
//
// Returns the star counts and an error if any.
func (r *ReviewUseCase) GetStarCountsUsingVendorIDCourtType(vendorID uint, courtType string) (*entities.ReviewStarsCount, error) {
	// Get the star counts using the vendor ID and court type
	return r.ReviewRepository.GetStarCountsUsingVendorIDCourtType(vendorID, courtType)
}

// GetReviewsUsingVendorIDCourtType is a use case that handles the request to get the
// reviews using the vendor ID and court type.
// 
// vendorID: The id of the vendor.
// courtType: The type of the court.
//
// Returns the reviews of court and error if any.
func (r *ReviewUseCase) GetReviewsUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Review, error) {
	// Get the reviews using the vendor ID and court type
	return r.ReviewRepository.GetUsingVendorIDCourtType(vendorID, courtType)
}

// SanitizeCreateReviewForm is a use case that sanitizes the create review form.
//
// form: The create review form.
//
// Returns nothing.
func (r *ReviewUseCase) SanitizeCreateReviewForm(form *dto.CreateReviewFormDTO) {
	// Sanitize the review form
	form.Review = strings.TrimSpace(form.Review)
}

// ValidateCreateReviewForm is a use case that validates the create review form.
//
// form: The create review form.
//
// Returns a form error response message.
func (r *ReviewUseCase) ValidateCreateReviewForm(form *dto.CreateReviewFormDTO) types.FormErrorResponseMsg {
	// Create an empty error map
	errs := make(types.FormErrorResponseMsg)

	// Check if rating is valid
	if form.Rating <= 0 {
		errs["rating"] = append(errs["rating"], "Rating must be greater than 0")
	}

	// Check if rating is valid
	if form.Rating > 5 {
		errs["rating"] = append(errs["rating"], "Rating must be less than or equal to 5")
	}

	// Check if theres any error
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// ProcessCreateReview is a use case that processes the creation of a review.
//
// token: The JWT token.
// vendorID: The id of the vendor.
// courtID: The id of the court.
// form: The create review form.
//
// Returns an error if any.
func (r *ReviewUseCase) ProcessCreateReview(token *jwt.Token, vendorID int, courtID int, form *dto.CreateReviewFormDTO) (*models.Review, *entities.ProcessError) {
	// Get the user ID from the token
	claims := r.AuthUseCase.DecodeToken(token)

	// Check if user already book the court
	booked, err := r.BookingRepository.CheckUserHasBookCourt(claims.Id, uint(vendorID), uint(courtID))

	// Check if there is an error
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while checking if user has booked the court",
		}
	}

	// Return an error if user has not booked the court
	if !booked {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "User has not booked the court",
		}
	}

	// Get court using court id
	court, err := r.CourtRepository.GetUsingID(uint(courtID))

	// Check if there is an error
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while getting the court",
		}
	}

	// Check if user has reviewed the court
	reviewed, err := r.ReviewRepository.CheckUserHasReviewCourtType(claims.Id, uint(vendorID), court.CourtType.Type)

	// Check if there is an error
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while checking if user has reviewed the court",
		}
	}

	// Return an error if user has reviewed the court
	if reviewed {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "User has already reviewed the court",
		}
	}

	// Create a new review object
	review := &models.Review{
		UserID:      claims.Id,
		VendorID:    uint(vendorID),
		CourtTypeID: court.CourtTypeID,
		Rating:      form.Rating,
		Review:      form.Review,
	}

	// Create the review
	err = r.ReviewRepository.Create(review)

	// Check if there is an error
	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "An error occurred while creating the review",
		}
	}

	return review, nil
}
