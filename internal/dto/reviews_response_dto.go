package dto

import "main/data/models"

// ReviewsResponseDTO is a struct that defines the reviews response DTO.
type ReviewsResponseDTO struct {
	// Reviews is a slice of review DTOs
	Reviews []ReviewDTO `json:"reviews"`
}

// FromModels is a function that converts a slice of review models to a reviews response DTO.
//
// m: The slice of review models.
//
// Returns the reviews response DTO.
func (r ReviewsResponseDTO) FromModels(m *[]models.Review) *ReviewsResponseDTO {
	// reviews is a slice of review DTOs
	var reviews []ReviewDTO

	// Iterate over the review models
	for _, review := range *m {
		// Append the review DTO to the reviews slice
		reviews = append(reviews, ReviewDTO{}.FromModel(&review))
	}

	return &ReviewsResponseDTO{
		Reviews: reviews,
	}
}
