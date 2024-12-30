package dto

// CreateOrderDTO is a type that defines the create booking DTO.
type CreateOrderDTO struct {
	// VendorID is the vendor ID.
	VendorID uint `json:"vendor_id"`

	// Date is the book date.
	Date string `json:"date"`

	// Bookings is the bookings.
	Bookings *[]CreateOrderDTOInner `json:"bookings"`
}

// CreateOrderDTOInner is a type that defines the create
// booking DTO inner.
type CreateOrderDTOInner struct {
	// CourtID is the court ID.
	CourtID uint `json:"court_id"`

	// BookTime is the book time.
	BookTime []string `json:"book_times"`
}
