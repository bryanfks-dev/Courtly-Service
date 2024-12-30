package midtrans

import "strconv"

// MidtransIDToOrderID is a function that converts a Midtrans ID to an order ID.
//
// midtransID: The Midtrans ID.
//
// Returns the order ID and error if any.
func MidtransIDToOrderID(midtransID string) (uint, error) {
	// Convert the Midtrans ID to an order ID
	orderID, err := strconv.Atoi(midtransID[len("MID-Order-"):])

	// Return an error if any
	if err != nil {
		return 0, err
	}

	return uint(orderID), nil
}
