package controllers

import (
	"log"
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"main/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ReviewController is a struct that defines the ReviewController
type ReviewController struct {
	ReviewUseCase *usecases.ReviewUseCase
}

// NewReviewController is a factory function that returns a new instance of the ReviewController.
//
// r: The review use case.
//
// Returns a new instance of the ReviewController.
func NewReviewController(r *usecases.ReviewUseCase) *ReviewController {
	return &ReviewController{
		ReviewUseCase: r,
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
	if !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get the rating query parameter
	ratingParam := c.QueryParam("rating")

	// Create a new rating variable
	var rating int

	// Check if the rating query parameter is empty
	if !utils.IsBlank(ratingParam) {
		// Convert the rating query parameter to an integer
		rating, err = strconv.Atoi(ratingParam)

		// Check if the star query parameter is invalid
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
				Success: false,
				Message: "Invalid rating query parameter",
				Data:    nil,
			})
		}

		// Validate the rating parameter
		if !r.ReviewUseCase.ValidateRatingParam(rating) {
			return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
				Success: false,
				Message: "Invalid rating parameter",
				Data:    nil,
			})
		}
	}

	// Get the reviews from the database
	reviewsMap, err := r.ReviewUseCase.GetCourtTypeReviews(uint(vendorID), courtType, &rating)

	// Check if there is an error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get court reviews",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Reviews retrieved successfully",
		Data:    dto.ReviewsResponseDTO{}.FromMap(reviewsMap),
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

	// Get the rating query parameter
	ratingParam := c.QueryParam("rating")

	// Create a new rating and error variable
	var (
		rating int
		err    error
	)

	// Check if the rating query parameter is empty
	if !utils.IsBlank(ratingParam) {
		// Convert the rating query parameter to an integer
		rating, err = strconv.Atoi(ratingParam)

		// Check if the star query parameter is invalid
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
				Success: false,
				Message: "Invalid rating query parameter",
				Data:    nil,
			})
		}

		// Validate the rating parameter
		if !r.ReviewUseCase.ValidateRatingParam(rating) {
			return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
				Success: false,
				Message: "Invalid rating parameter",
				Data:    nil,
			})
		}
	}

	// Get the reviews from the database
	reviewsMap, err := r.ReviewUseCase.GetCurrentVendorReviews(cc.Token, &rating)

	// Check if there is an error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get vendor reviews",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Vendor reviews retrieved successfully",
		Data:    dto.ReviewsResponseDTO{}.FromMap(reviewsMap),
	})
}

// CreateReview is a controller that handles the request to create a review.
// Endpoint: POST /vendors/:id/courts/:type/reviews
//
// c: The echo context.
//
// Returns a response containing the created review.
func (r *ReviewController) CreateReview(c echo.Context) error {
	// Get the vendor id from the URL
	id := c.Param("id")

	// Check if the vendor id is empty
	if utils.IsBlank(id) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Vendor ID is required",
			Data:    nil,
		})
	}

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

	// Check if the court type is empty
	if utils.IsBlank(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Court type is required",
			Data:    nil,
		})
	}

	// Validate the court type
	if !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
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
	review, processErr := r.ReviewUseCase.ProcessCreateReview(cc.Token, vendorID, courtType, form)

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
