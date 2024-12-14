package types

import "main/data/models"

// CourtMap is a type that represent a court map.
// This map type should be formatted as following
// {
//     "court": ...,
//     "total_rating": ...
// }
type CourtMap map[string]any

// GetCourt is a function that returns the court from the court map.
//
// Returns the court.
func (c CourtMap) GetCourt() models.Court {
	return c["court"].(models.Court)
}

// GetTotalRating is a function that returns the total rating from the court map.
//
// Returns the total rating.
func (c CourtMap) GetTotalRating() float64 {
	return c["total_rating"].(float64)
}
