package dto

import "main/data/models"

// ReviewDTO is a struct that defines the review DTO.
type ReviewDTO struct {
	// ID is the primary key of the review
	ID uint `json:"id"`

	// User is the user of the review
	User *UserDTO `json:"user"`

	// CourtType is the type of the court
	CourtType string `json:"court_type"`

	// Rating is the rating of the court
	Rating int8 `json:"rating"`

	// Review is the review of the court
	Review string `json:"review"`

	// Date is the date of the review
	Date string `json:"date"`
}

// FromModel is a function that converts a review model to a review DTO.
//
// m: The review model.
//
// Returns the review DTO.
func (r ReviewDTO) FromModel(m *models.Review) *ReviewDTO {
	// date is the date of the review
	date, _ := m.Date.Value()

	return &ReviewDTO{
		ID:        m.ID,
		User:      UserDTO{}.FromModel(&m.User),
		CourtType: m.CourtType.Type,
		Rating:    m.Rating,
		Review:    m.Review,
		Date:      date.(string),
	}
}
