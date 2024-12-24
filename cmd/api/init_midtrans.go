package api

import "main/internal/providers/midtrans"

// initMidtrans is a helper function that initializes the
// Midtrans client.
func initMidtrans() {
	midtrans.CreateClient()
}