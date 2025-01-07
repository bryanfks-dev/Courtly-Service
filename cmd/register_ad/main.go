package main

import (
	"bufio"
	"fmt"
	"main/core/config"
	"main/core/enums"
	"main/data/models"
	"main/internal/providers/mysql"
	"main/internal/repository"
	"main/pkg/utils"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// RegisterForm is a struct that defines the register form.
type RegisterForm struct {
	// ImageName is the name of the image.
	ImageName string

	// VendorID is the id of the vendor
	VendorID uint

	// CourtType is the type of the court.
	CourtType string
}

// Repository initialization
var (
	VendorRepository        *repository.VendorRepository
	AdvertisementRepository *repository.AdvertisementRepository
)

// sanitizeForm is a helper function that sanitizes the register input.
//
// form: The register form form.
//
// Returns void
func sanitizeForm(form *RegisterForm) {
	form.ImageName = strings.TrimSpace(form.ImageName)
	form.CourtType = strings.TrimSpace(form.CourtType)
	form.CourtType = 
		strings.ToUpper(form.CourtType[:1]) + strings.ToLower(form.CourtType[1:])
}

// validateForm is a function that validates the register form.
//
// form: The register form.
//
// Returns void
func validateForm(form *RegisterForm) {
	// Check if the image name is blank
	if utils.IsBlank(form.ImageName) {
		panic("Image name is required")
	}

	// Check if the vendor id is invalid
	if form.VendorID <= 0 {
		panic("Invalid vendor id")
	}

	// Try to get the vendor using the vendor id
	_, err := VendorRepository.GetUsingID(form.VendorID)

	// Check if there is an error
	if err == gorm.ErrRecordNotFound {
		panic("Vendor not found")
	}

	// Check if there is an error
	if err != nil {
		panic(err.Error())
	}

	// Check if the court type is blank
	if utils.IsBlank(form.CourtType) {
		panic("Court type is required")
	}

	// Check if the court type is invalid
	if !enums.InCourtType(form.CourtType) {
		panic("Invalid court type")
	}
}

// registerAd is a function that registers an advertisement.
//
// form: The register form.
//
// Returns void
func registerAd(form *RegisterForm) {
	// Create the advertisement
	ad := models.Advertisement{
		Image:       form.ImageName,
		VendorID:    form.VendorID,
		CourtTypeID: enums.GetCourtTypeID(form.CourtType),
	}

	// Create the advertisement
	err := AdvertisementRepository.Create(&ad)

	// Check if there is an error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\nAdvertisement registered successfully!")
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

	fmt.Println("Register advertisement program")
	fmt.Println("=====================================")

	// Get the vendor register form
	form := RegisterForm{}

	// Create a new reader instance
	reader := bufio.NewReader(os.Stdin)

	// Get the image name
	fmt.Print("Enter image name[Include file extension][See assets/ads]: ")

	// Read the image name
	line, err := reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get image name: " + err.Error())
	}

	// Set the image name
	form.ImageName = line

	// Get the vendor id
	fmt.Print("Enter vendor id: ")

	// Read the vendor id
	line, err = reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get vendor id: " + err.Error())
	}

	// Convert the vendor id to uint
	vendorID, err :=
		strconv.ParseUint(strings.TrimSpace(line), 10, 64)

	// Return an error if any
	if err != nil {
		panic("Failed to convert vendor id: " + err.Error())
	}

	// Set the vendor id
	form.VendorID = uint(vendorID)

	// Get the court type
	fmt.Print("Enter court type[Football|Basketball|Tennis|Volleyball|Badminton]: ")

	// Read the court type
	line, err = reader.ReadString('\n')

	// Return an error if any
	if err != nil {
		panic("Failed to get court type: " + err.Error())
	}

	// Set the court type
	form.CourtType = line

	// Sanitize the form
	sanitizeForm(&form)

	// Validate the form
	validateForm(&form)

	// Register the advertisement
	registerAd(&form)
}
