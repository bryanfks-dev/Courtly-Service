package utils

import "main/core/types"

// CalculateTotalRating is a function that calculates the total rating of the reviews.
//
// starCount: The star count of the reviews.
// reviewCount: The total number of reviews.
//
// Returns the total rating.
func CalculateTotalRating(starCounts *types.StarCountsMap, reviewCount int64) float64 {
	// Check if there are no reviews
	if reviewCount == 0 {
		return 0.0
	}

	// Formula to calculate the total rating:
	// (1 * OneStarCount + 2 * TwoStarsCount + 3 * ThreeStarsCount + 4 * FourStarsCount + 5 * FiveStarsCount)
	// ------------------------------------------------------------------------------------------------------
	//                           							Total Reviews

	return (float64((*starCounts)[1]) + float64(2*(*starCounts)[2]) + float64(3*(*starCounts)[3]) + float64(4*(*starCounts)[4]) + float64(5*(*starCounts)[5])) / float64(reviewCount)
}
