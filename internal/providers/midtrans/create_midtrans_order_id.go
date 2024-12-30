package midtrans

import "strconv"

// CreateMidtransOrderId is a function that creates a new Midtrans order ID.
//
// orderID: the ID of the order.
//
// Returns the new Midtrans order ID.
func CreateMidtransOrderId(orderID uint) string {
	return "MID-Order" + "-" + strconv.Itoa(int(orderID))
}
