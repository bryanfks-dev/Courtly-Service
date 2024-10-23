package controllers

import (
	"log"
	"main/core/types"
	"main/data/models"
	"main/domain/usecases"
	"main/internal/dto"
	"main/internal/providers/mysql"
	"main/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// validateLoginForm is a helper function that validates the login input.
//
// data: The login form data.
//
// Returns the user model and a map of errors.
func validateLoginForm(data dto.LoginForm) (models.User, types.ResponseMsg) {
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

	// Define a user model
	var user models.User

	// Fetch the user from the database
	err := mysql.Conn.Model(models.User{}).Where("username = ?", data.Username).Find(&user).Error

	// Check if there is an error fetching the user
	if err != nil {
		log.Fatal("Error fetching user: ", err)

		return user, nil
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
		return user, errs
	}

	return user, nil
}

// Login is a function that handles the login request.
// Endpoint: POST /api/v1/login
//
// c: The echo context.
//
// Returns an error response if there is an error, otherwise a success response.
func Login(c echo.Context) error {
	// Create a new LoginForm object
	form := new(dto.LoginForm)

	// Bind the request body to the LoginForm object
	if err := c.Bind(form); err != nil {
		log.Default().Println("Error binding request body: ", err)

		return err
	}

	// Validate the login form
	user, errs := validateLoginForm(*form)

	// Check if there are any errors
	if errs != nil {
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
