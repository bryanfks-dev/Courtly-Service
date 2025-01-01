package dto

import (
	"main/core/types"
	"main/data/models"
)

// ReviewsResponseDTO is a struct that defines the reviews response DTO.
type ReviewsResponseDTO struct {
	// TotalRating is the total rating of the reviews
	TotalRating float64 `json:"total_rating"`

	// ReviewsTotal is the total number of reviews
	ReviewsTotal int64 `json:"reviews_total"`

	// Stars is the reviews stars DTO
	Stars *ReviewsStarsDTO `json:"stars"`

	// Reviews is a slice of review DTOs
	Reviews *[]ReviewDTO `json:"reviews"`
}

// FromMap is a function that converts a reviews map to a reviews response DTO.
//
// m: The reviews map.
//
// Returns a pointer to the reviews response DTO.
func (r ReviewsResponseDTO) FromMap(m *types.CourtReviewsMap) *ReviewsResponseDTO {
	// Create a slice of review DTOs
	dtos := []ReviewDTO{}

	// Convert the reviews to review DTOs
	for _, review := range *(*m)["reviews"].(*[]models.Review) {
		dtos = append(dtos, *ReviewDTO{}.FromModel(&review))
	}

	return &ReviewsResponseDTO{
		TotalRating:  (*m)["total_rating"].(float64),
		ReviewsTotal: (*m)["reviews_total"].(int64),
		Stars:        ReviewsStarsDTO{}.FromMap((*m)["star_counts"].(*types.StarCountsMap)),
		Reviews:      &dtos,
	}
}
