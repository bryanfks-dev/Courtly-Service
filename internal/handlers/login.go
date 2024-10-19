package handlers

import (
	"log"
	"main/core/types"
	"main/data/models"
	"main/domain/entities"
	"main/domain/usecases"
	"main/internal/dto"
	"main/internal/providers/database"
	"main/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// user is a global variable that holds the user model.
var user models.User

// validateLoginForm is a helper function that validates the login input.
//
// data: The login form data.
//
// Returns a boolean indicates the form is valid or not and a map of errors.
func validateLoginForm(data entities.LoginForm) (bool, types.ResponseMsg) {
	// Create an empty error map
	errs := make(types.ResponseMsg)

	// Check if the username is blank
	if utils.IsBlank(data.Username) {
		errs["username"] = append(errs["username"], "Username is required")
	}

	// Check if the password is blank
	if utils.IsBlank(data.Password) {
		errs["password"] = append(errs["password"], "Password is required")
	}

	// Fetch the user from the database
	err := database.Conn.Model(models.User{}).Where("username = ?", data.Username).Find(&user).Error

	// Check if there is an error fetching the user
	if err != nil {
		log.Fatal("Error fetching user: ", err)

		return false, nil
	}

	// Check if the user exists
	if (user == models.User{}) {
		errs["username"] = append(errs["username"], "Username does not exist")
	}

	// Check if the password is correct
	if (user != models.User{} && !usecases.VerifyPassword(data.Password, user.Password)) {
		errs["password"] = append(errs["password"], "Password is incorrect")
	}

	// Check if there are any errors
	if (len(errs)) > 0 {
		return false, errs
	}

	return true, nil
}

// Login is a function that handles the login request.
// Endpoint: POST /api/v1/login
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func Login(c echo.Context) error {
	// Create a new LoginForm object
	form := new(entities.LoginForm)

	// Bind the request body to the LoginForm object
	if err := c.Bind(form); err != nil {
		log.Default().Println("Error binding request body: ", err)

		return err
	}

	// Validate the form
	if ok, errs := validateLoginForm(*form); !ok {
		return c.JSON(http.StatusBadRequest, dto.Response{
			StatusCode: http.StatusBadRequest,
			Message:    errs,
			Data:       nil,
		})
	}

	// Generate a token
	token, err := usecases.GenerateToken(user.ID)

	// Check if there is an error generating the token
	if err != nil {
		log.Fatal("Error generating token: ", err)

		return c.JSON(http.StatusInternalServerError, dto.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error generating token",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		StatusCode: http.StatusOK,
		Message:    "Login Success",
		Data: map[string]any{
			"user":  user,
			"token": token,
		},
	})
}
