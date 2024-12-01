package enums

// PaymentStatus is an enum that defines the payment status.
type PaymentStatus int

const (
	Success PaymentStatus = iota
	Pending
)

// Label is a function that returns the label of the payment status.
//
// Returns the label of the payment status.
func (p PaymentStatus) Label() string {
	return map[PaymentStatus]string{
		Success: "Success",
		Pending: "Pending",
	}[p]
}
