package config

import (
	"log"
	"main/pkg/utils"
)

// Midtrans is a struct that holds the configuration for the Midtrans API.
type Midtrans struct {
	// ServerKey is the server key for the Midtrans API.
	ServerKey string
}

// MidtransConfig is a global variable that holds the Midtrans configuration.
var MidtransConfig = Midtrans{}

// LoadData is a method that loads the midtrans configuration from the environment variables.
func (m Midtrans) LoadData() {
	// Get the server key from the environment variables
	key := utils.GetEnv("MIDTRANS_SERVER_KEY", "")

	// Check if the server key is empty
	if utils.IsBlank(key) {
		log.Fatal("MIDTRANS_SERVER_KEY is required")

		return
	}

	m.ServerKey = key

	MidtransConfig = m
}
