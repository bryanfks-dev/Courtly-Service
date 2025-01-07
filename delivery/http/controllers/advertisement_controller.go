package controllers

import (
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AdvertisementController is a struct that defines the AdvertisementController.
type AdvertisementController struct {
	AdvertisementUseCase *usecases.AdvertisementUseCase
}

// NewAdvertisementController is a factory function that returns a
// new instance of the AdvertisementController.
//
// a: the AdvertisementUseCase instance.
//
// Returns the AdvertisementController instance.
func NewAdvertisementController(a *usecases.AdvertisementUseCase) *AdvertisementController {
	return &AdvertisementController{
		AdvertisementUseCase: a,
	}
}

func (a *AdvertisementController) GetAdvertisements(c echo.Context) error {
	// Get all advertisements
	ads, err := a.AdvertisementUseCase.GetAdvertisements()

	// Check if there is an error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.ResponseDTO{
			Success: false,
			Message: "Failed to get advertisements",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &dto.ResponseDTO{
		Success: true,
		Message: "Successfully get advertisements",
		Data:    dto.AdvertisementsResponseDTO{}.FromModels(ads),
	})
}
