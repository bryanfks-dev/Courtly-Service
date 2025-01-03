package dto

// UpdateCourtFormDTO is a struct that defines the update court form data transfer object.
type UpdateCourtFormDTO struct {
	// PricePerHour is the price per hour of the court.
	PricePerHour float64 `json:"price_per_hour"`
}
