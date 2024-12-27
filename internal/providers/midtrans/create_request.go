package midtrans

import (
	"log"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// CreateRequest is a function that creates a new request.
//
// orderID: the ID of the order.
// courtID: the ID of the court.
// amount: the amount of the transaction.
//
// Returns an error if there is any.
func CreateRequest(orderID uint, courtID uint, amount int64) (*string, error) {
	// Create a new request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-Order" + strconv.Itoa(int(orderID)) + "-" + strconv.Itoa(int(courtID)),
			GrossAmt: amount,
		},
		EnabledPayments: []snap.SnapPaymentType{snap.PaymentTypeGopay, snap.PaymentTypeShopeepay, snap.PaymentTypeBCAVA, snap.PaymentTypeBRIVA},
	}

	// Create a new transaction
	res, err := MidtransClient.CreateTransaction(req)

	// Check if there is an error
	if err != nil {
		log.Println("Error creating transaction: ", err)

		return nil, err
	}

	return &res.RedirectURL, nil
}
