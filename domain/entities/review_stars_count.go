package entities

// ReviewStarsCount is a struct that defines the review stars count.
type ReviewStarsCount struct {
	// OneStar is the count of one star reviews.
	OneStar int64

	// TwoStars is the count of two star reviews.
	TwoStars int64

	// ThreeStars is the count of three star reviews.
	ThreeStars int64

	// FourStars is the count of four star reviews.
	FourStars int64

	// FiveStars is the count of five star reviews.
	FiveStars int64
}
