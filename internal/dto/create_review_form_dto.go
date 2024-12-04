package dto

// CreateReviewFormDTO is a struct that defines the data transfer object 
// for creating a review.
type CreateReviewFormDTO struct {
	// Rating is the rating of the review
	Rating int8 `json:"rating"`

	// Review is the review content
	Review string `json:"review"`
}
