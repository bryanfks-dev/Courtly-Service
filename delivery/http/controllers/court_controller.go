package controllers

import (
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CourtController is a struct that defines the CourtController
type CourtController struct {
	CourtUseCase *usecases.CourtUseCase
}

// NewCourtController is a factory function that returns a new instance of the CourtController.
//
// c: The court use case.
//
// Returns a new instance of the CourtController.
func NewCourtController(c *usecases.CourtUseCase) *CourtController {
	return &CourtController{
		CourtUseCase: c,
	}
}

// GetCourts is a controller that handles the get courts endpoint.
// Endpoint: GET /courts
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCourts(c echo.Context) error {
	// Get the courts
	courts, err := co.CourtUseCase.GetCourts()

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to get courts",
			Data:    nil,
		})
	}

	// Convert the court models to court DTOs
	courtsDTO := co.CourtUseCase.ConvertCourtModelsToDTOs(courts)

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success retrieve courts",
		Data:    courtsDTO,
	})
}

// GetCourtsUsingType is a controller that handles the get courts using type endpoint.
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCourtsUsingType(c echo.Context) error {
	// Get the court type from the URL
	courtType := c.Param("type")

	// Return an error if the court type is invalid
	if !co.CourtUseCase.ValidateCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get the courts
	courts, err := co.CourtUseCase.GetCourtsUsingType(courtType)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to get courts",
			Data:    nil,
		})
	}

	// Convert the court models to court DTOs
	courtsDTO := co.CourtUseCase.ConvertCourtModelsToDTOs(courts)

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success retrieve courts",
		Data:    courtsDTO,
	})
}

// GetCurrentVendorCourtTypes is a controller that handles the get current vendor court types endpoint.
// Endpoint: GET /vendors/me/courts/types
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCurrentVendorCourtTypes(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor court types
	courtTypes, err := co.CourtUseCase.GetCurrentVendorCourtTypes(cc.Token)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to get current vendor court types",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success retrieve current vendor court types",
		Data:    dto.CourtTypeResponseDTO{}.FromModels(*courtTypes),
	})
}

// GetCurrentVendorCourtType is a controller that handles the get current vendor court type endpoint.
// Endpoint: GET /vendors/me/courts/:type
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCurrentVendorCourtType(c echo.Context) error {
	// Get the court type from the URL
	courtType := c.Param("type")

	// Return an error if the court type is invalid
	if !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor courts
	courts, err := co.CourtUseCase.GetCurrentVendorCourtsUsingType(cc.Token, courtType)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to get current vendor courts",
			Data:    nil,
		})
	}

	// Convert the court models to court DTOs
	courtsDTO := co.CourtUseCase.ConvertCourtModelsToDTOs(courts)

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Success retrieve current vendor courts",
		Data: dto.VendorCourtTypeResponse{
			CourtType: courtType,
			Courts:    courtsDTO,
		},
	})
}
