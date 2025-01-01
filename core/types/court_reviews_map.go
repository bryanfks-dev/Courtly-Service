package types

// CourtReviewsMap is a map that maps a string to any.
// CourtReviewsMap should contains as follows:
// {
//     "reviews_total": ...,
//     "star_counts": ...,
//     "reviews": ...,
//     "total_rating": ...
// }
type CourtReviewsMap map[string]any
