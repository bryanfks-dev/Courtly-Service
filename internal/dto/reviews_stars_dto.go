package dto

import "main/core/types"

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

// FromMap is a function that converts a review stars count map to a reviews stars DTO.
//
// m: The review stars count map.
//
// Returns the reviews stars DTO.
func (r ReviewsStarsDTO) FromMap(m *types.StartCountsMap) *ReviewsStarsDTO {
	return &ReviewsStarsDTO{
		OneStar:    int((*m)[1]),
		TwoStars:   int((*m)[2]),
		ThreeStars: int((*m)[3]),
		FourStars:  int((*m)[4]),
		FiveStars:  int((*m)[5]),
	}
}
