package midtrans

import (
	"main/core/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var MidtransClient *snap.Client

// CreateClient is a function that creates a new Midtrans client.
//
// Returns 
func CreateClient() {
	MidtransClient = &snap.Client{}

	MidtransClient.New(config.MidtransConfig.ApiKey, midtrans.Sandbox)
}