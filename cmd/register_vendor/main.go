package main

import (
	"bufio"
	"fmt"
	"main/core/config"
	"main/core/shared"
	"main/data/models"
	"main/domain/usecases"
	"main/internal/providers/mysql"
	"main/internal/repository"
	"main/pkg/utils"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// RegisterForm is a struct that defines the register form.
type RegisterForm struct {
	// Name is the name of the vendor.
	Name string

	// Address is the address of the vendor.
	Address string

	// Email is the email of the vendor.
	Email string

	// Password is the password of the vendor.
	Password string

	// OpenTime is the open time of the vendor.
	OpenTime shared.TimeOnly

	// CloseTime is the close time of the vendor.
	CloseTime shared.TimeOnly
}

// sanitizeForm is a helper function that sanitizes the register input.
//
// form: The register form form.
//
// Returns void
func sanitizeForm(form *RegisterForm) {
	form.Name = strings.TrimSpace(form.Name)
	form.Address = strings.TrimSpace(form.Address)
	form.Email = strings.TrimSpace(form.Email)
}

// validateForm is a function that validates the register form.
//
// form: The register form.
//
// Returns void
func validateForm(form *RegisterForm) {
	// Check if the name is blank
	if utils.IsBlank(form.Name) {
		panic("Name is required")
	}

	// Check if the address is blank
	if utils.IsBlank(form.Address) {
		panic("Address is required")
	}

	// Check if the email is blank
	if utils.IsBlank(form.Email) {
		panic("Email is required")
	}

	// Check if the email is valid
	if !utils.IsValidEmail(form.Email) {
		panic("Email is invalid")
	}

	// Check if the open time is zero
	if form.OpenTime.IsZero() {
		panic("Open time is required")
	}

	// Check if the close time is zero
	if form.CloseTime.IsZero() {
		panic("Close time is required")
	}

	// Check if the close time is before the open time
	if form.CloseTime.Time.Before(form.OpenTime.Time) {
		panic("Close time must be after open time")
	}
}

// registerVendor is a function that registers a vendor.
//
// form: The register form.
//
// Returns void
func registerVendor(form *RegisterForm) {
	// Create a new auth use case
	a := usecases.NewAuthUseCase()

	// Generate vendor password
	generatedPassword := form.Email

	// Hash the password
	hashedPassword, err := a.HashPassword(generatedPassword)

	// Return an error if any
	if err != nil {
		return
	}

	// Set the password
	form.Password = hashedPassword

	// Create a new vendor object
	vendor := models.Vendor{
		Name:      form.Name,
		Address:   form.Address,
		Email:     form.Email,
		Password:  form.Password,
		OpenTime:  form.OpenTime,
		CloseTime: form.CloseTime,
	}

	// Create a new vendor repository
	v := repository.NewVendorRepository()

	// Create the vendor
	err = v.Create(&vendor)

	// Return an error if any
	if err != nil {
		panic("\nFailed to register vendor: " + err.Error())
	}

	fmt.Println("\nVendor registered successfully!")

	// Print the account details
	fmt.Println("Email: ", vendor.Email)
	fmt.Println("Password: ", generatedPassword)
}

// main is the entry point of the program.
func main() {
	// Load the environment variables
	err := godotenv.Load()

	// Check if there is an error loading the environment variables
	if err != nil {
		panic("Error loading environment variables: " + err.Error())
	}

	// Load the database configuration
	config.DBConfig.LoadData()

	// Connect to the database
	err = mysql.Connect()

	// Check if there is an error connecting to the database
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	// Close the database connection
	defer func() {
		err := mysql.CloseConnection()

		// Check if there is an error closing the database connection
		if err != nil {
			panic("Error closing the database connection: " + err.Error())
		}
	}()

	fmt.Println("Register vendor program")
	fmt.Println("=====================================")

	// Get the vendor register form
	form := RegisterForm{}

	// Create a new reader instance
	reader := bufio.NewReader(os.Stdin)

	// Get the vendor name
	fmt.Print("Enter vendor name: ")
	line, err := reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get vendor name: " + err.Error())
	}

	// Set the vendor name
	form.Name = line

	// Get the vendor address
	fmt.Print("Enter vendor address: ")
	line, err = reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get vendor address: " + err.Error())
	}

	// Set the vendor address
	form.Address = line

	// Get the vendor email
	fmt.Print("Enter vendor email: ")
	line, err = reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get vendor name: " + err.Error())
	}

	// Set the vendor email
	form.Email = line

	// Get the vendor open time
	fmt.Print("Enter vendor open time (HH:MM): ")
	line, err = reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get vendor open time: " + err.Error())
	}

	// Trim the line
	line = strings.TrimSpace(line)

	// Parse the open time
	openTime, err := time.Parse("15:04", line)

	// Return an error if any
	if err != nil {
		panic("Failed to parse open time: " + err.Error())
	}

	// Set the open time
	form.OpenTime = shared.TimeOnly{Time: openTime}

	// Get the vendor close time
	fmt.Print("Enter vendor close time (HH:MM): ")
	line, err = reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get vendor close time: " + err.Error())
	}

	// Trim the line
	line = strings.TrimSpace(line)

	// Parse the close time
	closeTime, err := time.Parse("15:04", line)

	// Return an error if any
	if err != nil {
		panic("Failed to parse close time: " + err.Error())
	}

	// Set the close time
	form.CloseTime = shared.TimeOnly{Time: closeTime}

	// Sanitize the form
	sanitizeForm(&form)

	// Validate the form
	validateForm(&form)

	// Register the vendor
	registerVendor(&form)
}
