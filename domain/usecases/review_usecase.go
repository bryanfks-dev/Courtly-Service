package usecases

import (
	"context"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/internal/dto"
	"main/internal/repository"
	"strings"
	"sync"

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
func (r *ReviewUseCase) GetCurrentVendorStarCounts(token *jwt.Token) (*types.StarCountsMap, error) {
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
func (r *ReviewUseCase) CalculateTotalRating(starCount *types.StarCountsMap, reviewCount int64) float64 {
	// Check if there are no reviews
	if reviewCount == 0 {
		return 0.0
	}

	// Formula to calculate the total rating:
	// (1 * OneStar + 2 * TwoStars + 3 * ThreeStars + 4 * FourStars + 5 * FiveStars)
	// -----------------------------------------------------------------------------
	//                           Total Reviews

	return (float64((*starCount)[1]) + float64(2*(*starCount)[2]) + float64(3*(*starCount)[3]) + float64(4*(*starCount)[4]) + float64(5*(*starCount)[5])) / float64(reviewCount)
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

// ValidateRatingParam is a use case that validates the rating parameter.
//
// rating: The rating parameter.
//
// Returns a boolean value.
func (r *ReviewUseCase) ValidateRatingParam(rating int) bool {
	// Check if rating is valid
	// Rating must be greater than 0 and less than or equal to 5

	// Check if rating is less or equal to 0
	if rating <= 0 {
		return false
	}

	// Check if rating is greater than 5
	if rating > 5 {
		return false
	}

	return true
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

// GetCourtTypeReviews is a use case that handles the request to get the court type reviews.
//
// courtType: The type of the court.
//
// Returns the reviews map and an error if any.
func (r *ReviewUseCase) GetCourtTypeReviews(vendorID uint, courtType string, rating *int) (*types.CourtReviewsMap, error) {
	// Create a new context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Create a new wait group for concurrency and error
	var (
		wg  sync.WaitGroup
		err error
	)

	// Create reviews map
	reviews := make(types.CourtReviewsMap)

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the review count
		reviewCount, e :=
			r.ReviewRepository.GetCountUsingVendorIDCourtType(vendorID, courtType)

		// Check if there is an error
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the review count to the reviews map
		reviews["reviews_total"] = reviewCount
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the star counts
		starCounts, e :=
			r.ReviewRepository.GetStarCountsUsingVendorIDCourtType(vendorID, courtType)

		// Check if there is an error
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the star counts to the reviews map
		reviews["star_counts"] = starCounts
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Create a variable to store the reviews and error
		var (
			records *[]models.Review
			e       error
		)

		// Get the reviews
		// Check if the rating query parameter is empty
		if rating != nil {
			records, e =
				r.ReviewRepository.GetUsingVendorIDCourtTypeRating(vendorID, courtType, *rating)
		} else {
			records, e = r.ReviewRepository.GetUsingVendorIDCourtType(vendorID, courtType)
		}

		// Check if there is an error
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the reviews to the reviews map
		reviews["reviews"] = records
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Return an error if any
	if err != nil {
		return nil, err
	}

	// Calculate the total rating
	reviews["total_rating"] = r.CalculateTotalRating(reviews["star_counts"].(*types.StarCountsMap), reviews["reviews_total"].(int64))

	return &reviews, err
}

// GetCurrentVendorReviews is a use case that handles the request to get the current vendor's reviews.
//
// token: The JWT token.
// rating: The rating of the review.
//
// Returns the reviews map and an error if any.
func (r *ReviewUseCase) GetCurrentVendorReviews(token *jwt.Token, rating *int) (*types.CourtReviewsMap, error) {
	// Get the vendor ID from the token
	claims := r.AuthUseCase.DecodeToken(token)

	// Create a new context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	// Defer the cancel function
	defer cancel()

	// Create a new wait group for concurrency and error
	var (
		wg  sync.WaitGroup
		err error
	)

	// Create a reviews map variable
	reviews := make(types.CourtReviewsMap)

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the review count
		reviewCount, e := r.ReviewRepository.GetCountUsingVendorID(claims.Id)

		// Check if there is an error
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the review count to the reviews map
		reviews["reviews_total"] = reviewCount
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the star counts
		starCounts, e := r.ReviewRepository.GetStarCountsUsingVendorID(claims.Id)

		// Check if there is an error
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the star counts to the reviews map
		reviews["star_counts"] = starCounts
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Create a variable to store the reviews and error
		var (
			records *[]models.Review
			e       error
		)

		// Get the reviews
		// Check if the rating query parameter is empty
		if rating != nil {
			records, e =
				r.ReviewRepository.GetUsingVendorIDRating(claims.Id, *rating)
		} else {
			records, e = r.ReviewRepository.GetUsingVendorID(claims.Id)
		}

		// Check if there is an error
		if e != nil {
			err = e

			cancel()

			return
		}

		// Add the reviews to the reviews map
		reviews["reviews"] = records
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if there is an error
	if err != nil {
		return nil, err
	}

	// Calculate the total rating
	reviews["total_rating"] = r.CalculateTotalRating(reviews["star_counts"].(*types.StarCountsMap), reviews["reviews_total"].(int64))

	return &reviews, nil
}

// GetReviewsUsingVendorIDCourtTypeRating is a use case that handles the request to get the
// reviews using the vendor ID, court type, and rating.
//
// vendorID: The id of the vendor.
// courtType: The type of the court.
// rating: The rating of the review.
//
// Returns the reviews and an error if any.
func (r *ReviewUseCase) GetReviewsUsingVendorIDCourtTypeRating(vendorID uint, courtType string, rating int) (*[]models.Review, error) {
	// Get the reviews using the vendor ID, court type, and rating
	return r.ReviewRepository.GetUsingVendorIDCourtTypeRating(vendorID, courtType, rating)
}

// GetCurrentVendorReviewsUsingRating is a use case that handles the request to get the current vendor's
// reviews using the rating.
//
// token: The JWT token.
// rating: The rating of the review.
//
// Returns the reviews and an error if any.
func (r *ReviewUseCase) GetCurrentVendorReviewsUsingRating(token *jwt.Token, rating int) (*[]models.Review, error) {
	// Get the vendor ID from the token
	claims := r.AuthUseCase.DecodeToken(token)

	// Get the reviews using the vendor ID and rating
	return r.ReviewRepository.GetUsingVendorIDRating(claims.Id, rating)
}
