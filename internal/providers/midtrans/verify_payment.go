package midtrans

import (
	"errors"
	"main/core/config"
	"main/core/enums"
	"main/internal/repository"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

// VerifyPayment is a function that verifies the payment from Midtrans.
//
// data: the data from the notification.
//
// Returns an error if there is any.
func VerifyPayment(data map[string]any) error {
	// Create new order repository
	orderRepository := repository.NewOrderRepository()

	// Create new Midtrans client
	client := &coreapi.Client{}

	client.New(config.MidtransConfig.ApiKey, midtrans.Sandbox)

	// Get midtrans order id from notification
	midtransOrderID, exist := data["order_id"].(string)

	// Check if order id not found
	if !exist {
		return errors.New("order ID not found")
	}

	// Check transaction to Midtrans with param orderId
	transactionStatusResp, err := client.CheckTransaction(midtransOrderID)

	if err != nil {
		return err
	}

	// Check if transaction status response is not nil
	if transactionStatusResp != nil {
		// Do set transaction status based on response from check transaction status
		if transactionStatusResp.TransactionStatus == "capture" {
			if transactionStatusResp.FraudStatus == "challenge" {
				// TODO set transaction status on your database to 'challenge'
				// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			} else if transactionStatusResp.FraudStatus == "accept" {
				// Convert midtrans order id to order id
				orderID, err := MidtransIDToOrderID(midtransOrderID)

				// Return an error if any
				if err != nil {
					return err
				}

				orderRepository.UpdatePaymentStatusUsingID(orderID, enums.Success.Label())
			}
		} else if transactionStatusResp.TransactionStatus == "settlement" {
			// Convert midtrans order id to order id
			orderID, err := MidtransIDToOrderID(midtransOrderID)

			// Return an error if any
			if err != nil {
				return err
			}

			orderRepository.UpdatePaymentStatusUsingID(orderID, enums.Success.Label())
		} else if transactionStatusResp.TransactionStatus == "deny" {
			// TODO you can ignore 'deny', because most of the time it allows payment retries
			// and later can become success
		} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
			// Convert midtrans order id to order id
			orderID, err := MidtransIDToOrderID(midtransOrderID)

			// Return an error if any
			if err != nil {
				return err
			}

			orderRepository.UpdatePaymentStatusUsingID(orderID, enums.Canceled.Label())
		}
	}

	return nil
}
