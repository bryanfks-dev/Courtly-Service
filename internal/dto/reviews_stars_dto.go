package dto

import "main/domain/entities"

// ReviewsStarsDTO is a struct that defines the reviews stars DTO.
type ReviewsStarsDTO struct {
	// OneStar is the number of one star reviews
	OneStar int `json:"1"`

	// TwoStars is the number of two star reviews
	TwoStars int `json:"2"`

	// ThreeStars is the number of three star reviews
	ThreeStars int `json:"3"`

	// FourStars is the number of four star reviews
	FourStars int `json:"4"`

	// FiveStars is the number of five star reviews
	FiveStars int `json:"5"`
}

// FromEntity is a function that converts a review stars count entity to a reviews stars DTO.
//
// e: The review stars count entity.
//
// Returns the reviews stars DTO.
func (r ReviewsStarsDTO) FromEntity(e *entities.ReviewStarsCount) *ReviewsStarsDTO {
	return &ReviewsStarsDTO{
		OneStar:    int(e.OneStar),
		TwoStars:   int(e.TwoStars),
		ThreeStars: int(e.ThreeStars),
		FourStars:  int(e.FourStars),
		FiveStars:  int(e.FiveStars),
	}
}
