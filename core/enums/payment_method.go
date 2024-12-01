package enums

// PaymentMethod is an enum that defines the payment methods.
type PaymentMethod int

const (
	OVO PaymentMethod = iota
	Dana
	Gopay
	ShopeePay
	BCA
	BNI
	BRI
)

// Label is a function that returns the label of the payment method.
//
// Returns the label of the payment method.
func (p PaymentMethod) Label() string {
	return map[PaymentMethod]string{
		OVO:       "OVO",
		Dana:      "Dana",
		Gopay:     "Gopay",
		ShopeePay: "ShopeePay",
		BCA:       "BCA",
		BNI:       "BNI",
		BRI:       "BRI",
	}[p]
}
