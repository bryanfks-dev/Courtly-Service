package midtrans

import (
	"main/core/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

// MidtransClient is a global variable that holds the Midtrans client.
var MidtransClient *coreapi.Client

// CreateClient is a function that creates a new Midtrans client.
//
// Returns void
func CreateClient() {
	MidtransClient = &coreapi.Client{}

	MidtransClient.New(config.MidtransConfig.ServerKey, midtrans.Sandbox)
}
