package controllers

import (
	"context"
	"log"
	"main/data/models"
	"main/domain/entities"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

// ReviewController is a struct that defines the ReviewController
type ReviewController struct {
	ReviewUseCase *usecases.ReviewUseCase
	CourtUseCase  *usecases.CourtUseCase
}

// NewReviewController is a factory function that returns a new instance of the ReviewController.
//
// r: The review use case.
//
// Returns a new instance of the ReviewController.
func NewReviewController(r *usecases.ReviewUseCase, c *usecases.CourtUseCase) *ReviewController {
	return &ReviewController{
		ReviewUseCase: r,
		CourtUseCase:  c,
	}
}

// GetCourtTypeReviews is a controller that handles the request to
// get the reviews of a court using the court id and court type.
// Endpoint: GET /vendors/:id/courts/types/:type/reviews
//
// c: The echo context.
//
// Returns a response containing the reviews of the court.
func (r *ReviewController) GetCourtTypeReviews(c echo.Context) error {
	// Get the id from the URL
	id := c.Param("id")

	// Convert the id to an integer
	vendorID, err := strconv.Atoi(id)

	// Check if the id is invalid
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid vendor ID",
			Data:    nil,
		})
	}

	// Get the court type from the URL
	courtType := c.Param("type")

	// Validate the court type
	if !r.CourtUseCase.ValidateCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Create a new context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Create a new wait group for concurrency
	var (
		wg sync.WaitGroup
	)

	// Create review count, star counts and reviews
	var (
		reviewCount int64
		starCounts  *entities.ReviewStarsCount
		reviews     *[]models.Review
	)

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the review count
		reviewCount, err =
			r.ReviewUseCase.GetReviewCountUsingVendorIDCourtType(uint(vendorID), courtType)

		// Check if there is an error
		if err != nil {
			err = c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get review count",
				Data:    nil,
			})

			cancel()
		}
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the star counts
		starCounts, err =
			r.ReviewUseCase.GetStarCountsUsingVendorIDCourtType(uint(vendorID), courtType)

		// Check if there is an error
		if err != nil {
			err = c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get star counts",
				Data:    nil,
			})

			cancel()
		}
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the reviews
		reviews, err =
			r.ReviewUseCase.GetReviewsUsingVendorIDCourtType(uint(vendorID), courtType)

		// Check if there is an error
		if err != nil {
			err = c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get reviews",
				Data:    nil,
			})

			cancel()
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if there is an error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Reviews retrieved successfully",
		Data: dto.ReviewsResponseDTO{}.FromModels(
			r.ReviewUseCase.CalculateTotalRating(starCounts, reviewCount),
			int(reviewCount),
			starCounts,
			reviews,
		),
	})
}

// GetCurrentVendorReviews is a controller that handles the request to
// get the reviews of the current vendor.
// Endpoint: GET /vendors/me/reviews
//
// c: The echo context.
//
// Returns a response containing the reviews of the current vendor.
func (r *ReviewController) GetCurrentVendorReviews(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Create a new context with a cancel function
	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Create a new wait group for concurrency and error
	var (
		wg  sync.WaitGroup
		err error
	)

	// Create review count, star counts and reviews
	var (
		reviewCount int64
		starCounts  *entities.ReviewStarsCount
		reviews     *[]models.Review
	)

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the review count
		reviewCount, err = r.ReviewUseCase.GetCurrentVendorReviewCount(cc.Token)

		// Check if there is an error
		if err != nil {
			err = c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get vendor review count",
				Data:    nil,
			})

			cancel()
		}
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the star counts
		starCounts, err = r.ReviewUseCase.GetCurrentVendorStarCounts(cc.Token)

		// Check if there is an error
		if err != nil {
			err = c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get vendor review star counts",
				Data:    nil,
			})

			cancel()
		}
	}()

	// Add a new wait group
	wg.Add(1)

	go func() {
		// Defer the wait group
		defer wg.Done()

		// Get the reviews
		reviews, err =
			r.ReviewUseCase.GetCurrentVendorReviews(cc.Token)

		// Check if there is an error
		if err != nil {
			err = c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Success: false,
				Message: "Failed to get vendor reviews",
				Data:    nil,
			})

			cancel()
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if there is an error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Vendor reviews retrieved successfully",
		Data: dto.ReviewsResponseDTO{}.FromModels(
			r.ReviewUseCase.CalculateTotalRating(starCounts, reviewCount),
			int(reviewCount),
			starCounts,
			reviews,
		),
	})
}

// CreateReview is a controller that handles the request to create a review.
// Endpoint: POST /vendors/:id/courts/types/:type/reviews
//
// c: The echo context.
//
// Returns a response containing the created review.
func (r *ReviewController) CreateReview(c echo.Context) error {
	// Get the vendor id from the URL
	id := c.Param("vendorID")

	// Convert the id to an integer
	vendorID, err := strconv.Atoi(id)

	// Check if the id is invalid
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid vendor ID",
			Data:    nil,
		})
	}

	// Get the court type from the URL
	id = c.Param("courtID")

	// Convert the id to an integer
	courtID, err := strconv.Atoi(id)

	// Check if the id is invalid
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court ID",
			Data:    nil,
		})
	}

	// Create a new CreateReviewForm dto object
	form := new(dto.CreateReviewFormDTO)

	// Bind the request body to the CreateReviewForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Sanitize the review form
	r.ReviewUseCase.SanitizeCreateReviewForm(form)

	// Validate the login form
	errs := r.ReviewUseCase.ValidateCreateReviewForm(form)

	// Check if there are any errors
	if errs != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Process the creation of the review
	review, processErr := r.ReviewUseCase.ProcessCreateReview(cc.Token, vendorID, courtID, form)

	// Check if there is an error
	if processErr != nil {
		// Check if the error is a client error
		if processErr.ClientError {
			return c.JSON(http.StatusForbidden, dto.ResponseDTO{
				Success: false,
				Message: processErr.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: processErr.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, dto.ResponseDTO{
		Success: true,
		Message: "Review created successfully",
		Data: dto.ReviewResponseDTO{
			Review: dto.ReviewDTO{}.FromModel(review),
		},
	})
}
