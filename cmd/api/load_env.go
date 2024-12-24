package api

import (
	"main/core/config"
	"sync"
)

// LoadEnv is a function that loads the environment variables.
// It loads the database, JWT, midtrans, and server configuration.
//
// Returns void.
func LoadEnv() {
	// Wait group to wait for all the configurations to load
	var wg sync.WaitGroup

	// Add the number of configurations to load
	wg.Add(4)

	// Load the configurations in parallel
	go func() {
		config.DBConfig.LoadData()
		wg.Done()
	}()

	go func() {
		config.JWTConfig.LoadData()
		wg.Done()
	}()

	go func() {
		config.MidtransConfig.LoadData()
		wg.Done()
	} ()

	go func() {
		config.ServerConfig.LoadData()

		wg.Done()
	}()

	// Wait for all the configurations to load
	wg.Wait()
}
