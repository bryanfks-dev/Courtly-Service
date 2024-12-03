package controllers

import (
	"log"
	"main/core/enums"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get courts",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success retrieve courts",
		Data:    dto.CourtsResponseDTO{}.FromCourtModels(courts),
	})
}

// GetCourtUsingID is a controller that handles the get court using ID endpoint.
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCourtUsingID(c echo.Context) error {
	// Get the court ID from the URL
	id := c.Param("id")

	// Convert the ID to an integer
	courtID, err := strconv.Atoi(id)

	// Return an error if the ID is invalid
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court ID",
			Data:    nil,
		})
	}

	// Get the courts
	court, err := co.CourtUseCase.GetCourtUsingID(uint(courtID))

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get court",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success retrieve court",
		Data: dto.CourtResponseDTO{
			Court: dto.CourtDTO{}.FromModel(court), // Return the first court
		},
	})
}

// GetCourtsUsingCourtType is a controller that handles the get courts using type endpoint.
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCourtsUsingCourtType(c echo.Context) error {
	// Get the court type from the URL
	courtType := c.Param("type")

	// Return an error if the court type is invalid
	if !co.CourtUseCase.ValidateCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get the courts
	courts, err := co.CourtUseCase.GetCourtsUsingCourtType(courtType)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get courts",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success retrieve courts",
		Data:    dto.CourtsResponseDTO{}.FromCourtModels(courts),
	})
}

// GetCurrentVendorCourtsUsingCourtType is a controller that handles the get
// current vendor courts using court type endpoint.
// Endpoint: GET /vendors/me/courts/:type
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCurrentVendorCourtsUsingCourtType(c echo.Context) error {
	// Get the court type from the URL
	courtType := c.Param("type")

	// Return an error if the court type is invalid
	if !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor courts
	courts, err := co.CourtUseCase.GetCurrentVendorCourtsUsingCourtType(cc.Token, courtType)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get current vendor courts",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success retrieve current vendor courts",
		Data:    dto.VendorCourtResponseDTO{}.FromCourtModels(courts),
	})
}

// CreateNewCourt is a controller that handles the create new court endpoint.
// Endpoint: POST /vendors/me/courts/types/:type/new
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) CreateNewCourt(c echo.Context) error {
	// Get the court type from the URL
	courtType := c.Param("type")

	// Return an error if the court type is invalid
	if !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Create a new CreateNeCourtFormDTO object
	form := new(dto.CreateNewCourtFormDTO)

	// Bind the request body to the UserLoginForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding request body: ", err)

		return err
	}

	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor courts
	court, err := co.CourtUseCase.CreateNewCourt(cc.Token, courtType, form)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to create new court",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success create new court",
		Data: dto.CourtResponseDTO{
			Court: dto.CourtDTO{}.FromModel(court),
		},
	})
}
