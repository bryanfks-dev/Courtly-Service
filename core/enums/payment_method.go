package enums

// PaymentMethod is an enum that defines the payment methods.
type PaymentMethod int

const (
	OVO PaymentMethod = iota
	Dana
	Gopay
	ShopeePay
	BCA
	BRI
)

// paymentMethods is a list of payment methods.
var paymentMethods = []string{
	"OVO",
	"Dana",
	"Gopay",
	"Shopee Pay",
	"BCA",
	"BRI",
}

// paymentMethodsApiValue is a list of payment methods API value.
var paymentMethodsApiValue = map[string]int{
	"OVO":        1,
	"DANA":       2,
	"GOPAY":      3,
	"SHOPEE_PAY": 4,
	"BCA":        5,
	"BRI":        6,
}

// Label is a function that returns the label of the payment method.
//
// Returns the label of the payment method.
func (p PaymentMethod) Label() string {
	return paymentMethods[p]
}

// PaymentMethods is a function that returns the list of payment methods.
//
// Returns the list of payment methods.
func PaymentMethods() []string {
	return paymentMethods
}

// GetPaymentMethodID is a function that returns the ID of the payment method.
//
// val: The value of the payment method.
//
// Returns the ID of the payment method.
func GetPaymentMethodIDFromRequest(val string) int {
	return paymentMethodsApiValue[val]
}
