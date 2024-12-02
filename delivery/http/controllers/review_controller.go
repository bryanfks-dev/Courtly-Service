package controllers

import (
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"
	"strconv"

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

// GetCourtReviewsUsingIDCourtType is a controller that handles the request to
// get the reviews of a court using the court id and court type.
// Endpoint: GET /vendors/:id/courts/types/:type/reviews
//
// c: The echo context.
//
// Returns a response containing the reviews of the court.
func (r *ReviewController) GetCourtReviewsUsingIDCourtType(c echo.Context) error {
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

	// Get the reviews
	reviews, err :=
		r.ReviewUseCase.GetCourtReviewsUsingVendorIDCourtType(uint(vendorID), courtType)

	// Check if there is an error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get reviews",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Reviews retrieved successfully",
		Data:    dto.ReviewsResponseDTO{}.FromModels(reviews),
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

	// Get the current vendor reviews
	reviews, err := r.ReviewUseCase.GetCurrentVendorReviews(cc.Token)

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
		Data:    dto.ReviewsResponseDTO{}.FromModels(reviews),
	})
}
