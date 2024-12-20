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
	// Get the court type from the query parameter
	courtType := c.QueryParam("type")

	// Return an error if the court type is invalid
	if !utils.IsBlank(courtType) && !enums.InCourtType(courtType) {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid court type",
			Data:    nil,
		})
	}

	// Get the vendor name from the query parameter
	vendorName := c.QueryParam("search")

	// Get the courts
	courtMaps, err := co.CourtUseCase.GetCourts(&courtType, &vendorName)

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
		Data:    dto.CourtsResponseDTO{}.FromCourtMaps(courtMaps),
	})
}

// GetVendorCourtsUsingCourtType is a controller that handles the get vendor courts using court type endpoint.
// Endpoint: GET /vendors/:id/courts/:type
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetVendorCourtsUsingCourtType(c echo.Context) error {
	// Get the vendor id from the URL
	vendorID, err := strconv.Atoi(c.Param("id"))

	// Return an error if the vendor id is invalid
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid vendor id",
			Data:    nil,
		})
	}

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

	// Get the courts
	courts, err := co.CourtUseCase.GetVendorCourtsUsingCourtType(uint(vendorID), courtType)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get vendor courts",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success retrieve vendor courts",
		Data:    dto.CourtsResponseDTO{}.FromCourtMaps(courts),
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
		Data:    dto.CurrentVendorCourtsResponseDTO{}.FromModels(courts),
	})
}

// CreateNewCourt is a controller that handles the create new court endpoint.
// Endpoint: POST /vendors/me/courts/:type/new
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

	// Validate the login form
	errs := co.CourtUseCase.ValidateCreateNewCourtForm(form)

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

	// Get the current vendor courts
	court, err := co.CourtUseCase.CreateNewCourt(cc.Token, courtType, form)

	// Return an error if any
	if err != nil {
		// Return an error if the client error is true
		if err.ClientError {
			return c.JSON(http.StatusForbidden, dto.ResponseDTO{
				Success: false,
				Message: err.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: err.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success create new court",
		Data: dto.CurrentVendorCourtResponseDTO{
			Court: dto.CurrentVendorCourtDTO{}.FromModel(court),
		},
	})
}

// AddCourt is a controller that handles the add court endpoint.
// Endpoint: POST /vendors/me/courts/:type
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) AddCourt(c echo.Context) error {
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

	// Add the court
	court, err := co.CourtUseCase.AddCourt(cc.Token, courtType)

	// Return an error if any
	if err != nil {
		// Return an error if the client error is true
		if err.ClientError {
			return c.JSON(http.StatusForbidden, dto.ResponseDTO{
				Success: false,
				Message: err.Message,
				Data:    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: err.Message,
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success add court",
		Data: dto.CurrentVendorCourtResponseDTO{
			Court: dto.CurrentVendorCourtDTO{}.FromModel(court),
		},
	})
}

// GetCurrentVendorCourtStats is a controller that handles the get current vendor
// court stats endpoint.
// Endpoint: GET /vendors/me/courts/stats
//
// c: The echo context.
//
// Returns an error if any.
func (co *CourtController) GetCurrentVendorCourtStats(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Get the current vendor court stats
	courtCounts, err := co.CourtUseCase.GetCurrentVendorCourtStats(cc.Token)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Failed to get current vendor court stats",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Success retrieve current vendor court stats",
		Data: &dto.CurrentVendorCourtStatsResponseDTO{
			FootballCourtCount:   (*courtCounts)[enums.Football.Label()],
			BasketballCourtCount: (*courtCounts)[enums.Basketball.Label()],
			TennisCourtCount:     (*courtCounts)[enums.Tennis.Label()],
			VolleyballCourtCount: (*courtCounts)[enums.Volleyball.Label()],
			BadmintonCourtCount:  (*courtCounts)[enums.Badminton.Label()],
		},
	})
}
