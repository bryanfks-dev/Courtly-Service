package dto

// CreateNewCourtFormDTO is a struct that defines the create new court form data transfer object.
type CreateNewCourtFormDTO struct {
	// PricePerHour is the price per hour of the court.
	PricePerHour float64 `json:"price_per_hour"`

	// CourtImage is the image of the court.
	CourtImage string `json:"court_image"`
}
