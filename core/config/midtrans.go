package config

import (
	"log"
	"main/pkg/utils"
)

// Midtrans is a struct that holds the configuration for the Midtrans API.
type Midtrans struct {
	// ApiKey is a string that holds the Midtrans  API key.
	ApiKey string
}

// MidtransConfig is a global variable that holds the Midtrans  configuration.
var MidtransConfig Midtrans = Midtrans{}

// LoadData is a method that loads the midtrans  configuration from the environment variables.
func (x Midtrans) LoadData() {
	// Get the server key from the environment variables
	key := utils.GetEnv("MIDTRANS_API_KEY", "")

	// Check if the server key is empty
	if utils.IsBlank(key) {
		log.Fatal("MIDTRANS_API_KEY is required")

		return
	}

	x.ApiKey = key

	MidtransConfig = x
}
