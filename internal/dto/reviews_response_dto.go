package dto

import (
	"main/data/models"
	"main/domain/entities"
	"main/pkg/utils"
	"strconv"
)

// ReviewsResponseDTO is a struct that defines the reviews response DTO.
type ReviewsResponseDTO struct {
	// TotalRating is the total rating of the reviews
	TotalRating float64 `json:"total_rating"`

	// ReviewsTotal is the total number of reviews
	ReviewsTotal int `json:"reviews_total"`

	// Stars is the reviews stars DTO
	Stars *ReviewsStarsDTO `json:"stars"`

	// Reviews is a slice of review DTOs
	Reviews *[]ReviewDTO `json:"reviews"`
}

// FromModels is a function that converts a slice of review models to a reviews response DTO.
//
// rate: The total rating of the reviews.
// reviewCount: The total number of reviews.
// stars: The reviews stars DTO.
// m: The slice of review models.
//
// Returns the reviews response DTO.
func (r ReviewsResponseDTO) FromModels(rate float64, reviewCount int, stars *entities.ReviewStarsCount, m *[]models.Review) *ReviewsResponseDTO {
	// reviews is a slice of review DTOs
	reviews := []ReviewDTO{}

	// Iterate over the review models
	for _, review := range *m {
		// Append the review DTO to the reviews slice
		reviews = append(reviews, *ReviewDTO{}.FromModel(&review))
	}

	return &ReviewsResponseDTO{
		TotalRating:  rate,
		ReviewsTotal: reviewCount,
		Stars:        ReviewsStarsDTO{}.FromEntity(stars),
		Reviews:      &reviews,
	}
}
