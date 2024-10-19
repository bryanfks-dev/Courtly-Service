package handlers

import (
	"log"
	"main/core/constants"
	"main/core/types"
	"main/data/models"
	"main/domain/usecases"
	"main/internal/dto"
	"main/internal/providers/database"
	"main/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// sanitizeRegisterForm is a helper function that sanitizes the register input.
//
// data: The register form data.
//
// Returns void
func sanitizeRegisterForm(data *dto.RegisterForm) {
	data.Username = strings.TrimSpace(data.Username)
	data.PhoneNumber = strings.TrimSpace(data.PhoneNumber)
}

// validateRegisterForm is a helper function that validates the register input.
func validateRegisterForm(data dto.RegisterForm) (bool, types.ResponseMsg) {
	// Create an empty error map
	errs := make(types.ResponseMsg)

	// Check if the username is blank
	if utils.IsBlank(data.Username) {
		errs["username"] = append(errs["username"], "Username is required")
	}

	// Check if the username is too short
	if len(data.Username) < constants.MINIMUM_USERNAME_LENGTH {
		errs["username"] = append(errs["username"], "Username is too short")
	}

	// Check if the phone number is blank
	if utils.IsBlank(data.PhoneNumber) {
		errs["phone_number"] = append(errs["phone_number"], "Phone number is required")
	}

	// Check if the password is blank
	if utils.IsBlank(data.Password) {
		errs["password"] = append(errs["password"], "Password is required")
	}

	// Check if the password is too short
	if len(data.Password) < constants.MINIMUM_PASSWORD_LENGTH {
		errs["password"] = append(errs["password"], "Password is too short")
	}

	// Check if the password and confirm password are the same
	if data.Password != data.ConfirmPassword {
		errs["confirm_password"] = append(errs["confirm_password"], "Password and confirm password do not match")
	}

	// Check if there are any errors
	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// Register is a function that handles the register request.
// Endpoint: POST /register
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func Register(c echo.Context) error {
	// Create a new RegisterForm object
	form := new(dto.RegisterForm)

	// Bind the request body to the RegisterForm object
	if err := c.Bind(form); err != nil {
		log.Default().Println("Error binding the request body: ", err)

		return err
	}

	// Sanitize the form
	sanitizeRegisterForm(form)

	// Validate the form
	if ok, errs := validateRegisterForm(*form); !ok {
		return c.JSON(http.StatusBadRequest, dto.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errs,
			Data:       nil,
		})
	}

	// Hash the password
	hashedPwd, err := usecases.HashPassword(form.Password)

	// Check if there is an error hashing the password
	if err != nil {
		log.Fatal("Error hashing the password: ", err)

		return c.JSON(http.StatusInternalServerError, dto.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error hashing the password",
			Data:       nil,
		})
	}

	// Create a new user
	newUser := models.User{
		Username:    form.Username,
		Password:    hashedPwd,
		PhoneNumber: form.PhoneNumber,
	}

	// Register the user into the database
	database.Conn.Create(&newUser)

	return c.JSON(http.StatusCreated, dto.Response{
		StatusCode: http.StatusCreated,
		Message:    "User created successfully",
		Data: map[string]any{
			"user": newUser,
		},
	})
}
