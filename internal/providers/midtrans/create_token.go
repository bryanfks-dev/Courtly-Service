package midtrans

import (
	"log"
	"main/core/constants"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// CreateToken is a function that creates a new token for the transaction.
//
// orderID: The order ID.
// amount: The amount.
//
// Returns a pointer to the token and an error if any.
func CreateToken(orderID uint, amount int64) (*string, error) {
	// Create a new request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  CreateMidtransOrderId(orderID),
			GrossAmt: amount + int64(constants.APP_FEE_PRICE),
		},
		EnabledPayments: []snap.SnapPaymentType{snap.PaymentTypeGopay, snap.PaymentTypeShopeepay, snap.PaymentTypeBCAVA, snap.PaymentTypeBRIVA},
		Expiry: &snap.ExpiryDetails{
			StartTime: time.Now().Format("2006-01-02 15:04:05 -0700"),
			Unit:      "minutes",
			Duration:  60,
		},
	}

	// Create a new transaction
	res, err := MidtransClient.CreateTransaction(req)

	// Check if there is an error
	if err != nil {
		log.Println("Error creating transaction: ", err)

		return nil, err
	}

	return &res.Token, nil
}
