package dto

// ReviewResponseDTO is a struct that defines the data transfer object for a review response.
type ReviewResponseDTO struct {
	// Review is the review data transfer object
	Review *ReviewDTO `json:"review"`
}
