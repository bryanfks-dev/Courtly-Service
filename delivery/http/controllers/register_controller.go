package controllers

import (
	"log"
	"main/domain/usecases"
	"main/internal/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterController is a struct that defines the register controller.
type RegisterController struct {
	RegisterUseCase *usecases.RegisterUseCase
}

// NewRegisterController is a factory function that returns a new instance of the RegisterController.
//
// r: The register use case.
//
// Returns a new instance of the RegisterController.
func NewRegisterController(r *usecases.RegisterUseCase) *RegisterController {
	return &RegisterController{RegisterUseCase: r}
}

// UserRegister is a function that handles the user register request.
// Endpoint: POST /auth/user/register
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func (r RegisterController) UserRegister(c echo.Context) error {
	// Create a new UserRegisterForm dto object
	form := new(dto.UserRegisterFormDTO)

	// Bind the request body to the UserRegisterForm object
	if err := c.Bind(form); err != nil {
		log.Println("Error binding the request body: ", err)

		return err
	}

	// Sanitize the form
	r.RegisterUseCase.SanitizeUserRegisterForm(form)

	// Validate the form
	if errs := r.RegisterUseCase.ValidateUserRegisterForm(form); errs != nil {
		return c.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: errs,
			Data:    nil,
		})
	}

	// Check if there is an error processing the form
	if err := r.RegisterUseCase.ProcessUserRegister(form); err != nil {
		// Check if the error is a client error
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

	// Register the user
	user, err := r.RegisterUseCase.CreateNewUser(form)

	// Return an error if any
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "An error occurred while registering the user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "User registered successfully",
		Data: dto.CurrentUserRegisterResponseDTO{
			User: dto.CurrentUserDTO{}.FromModel(user),
		},
	})
}
