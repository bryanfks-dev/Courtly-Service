package controllers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// VerifyPasswordController is a struct to hold the VerifyPassword function
type VerifyPasswordController struct {
	authUseCase *usecases.AuthUseCase
	userUseCase *usecases.UserUseCase
}

// NewVerifyPasswordController is a factory function that returns a new instance of VerifyPasswordController
//
// a: Instance of usecases.AuthUseCase
// u: Instance of usecases.UserUseCase
//
// Returns a new instance of VerifyPasswordController
func NewVerifyPasswordController(a *usecases.AuthUseCase, u *usecases.UserUseCase) VerifyPasswordController {
	return VerifyPasswordController{authUseCase: a, userUseCase: u}
}

// VerifyPassword is a controller to handle the request to verify the password of the user
//
// c: Context of the HTTP request
//
// Returns an error if any
func (v *VerifyPasswordController) VerifyPassword(c echo.Context) error {
	// Get custom context
	cc := c.(*dto.CustomContext)

	// Decode the token
	claims := v.authUseCase.DecodeToken(cc.Token)

	// Bind the data
	data := new(dto.VerifyPasswordData)

	// Return an error if the form data is invalid
	if err := c.Bind(data); err != nil {
		log.Println("Error binding form data: ", err)

		return c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid form data",
			Data:    nil,
		})
	}

	// Get the user from the database
	user, err := v.userUseCase.GetUserByID(claims.Id)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "An error occurred while getting the user",
			Data:    nil,
		})
	}

	// Verify the password
	valid := v.authUseCase.VerifyPassword(data.Password, user.Password)

	// Return an error if the password is invalid
	if !valid {
		return c.JSON(http.StatusForbidden, dto.Response{
			Success: false,
			Message: "Invalid password",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Password verified successfully",
		Data:    nil,
	})
}
