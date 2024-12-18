package repository

import (
	"log"
	"main/core/types"
	"main/data/models"
	"main/internal/providers/mysql"
)

// ReviewRepository is a struct that defines the ReviewRepository
type ReviewRepository struct{}

// NewReviewRepository is a factory function that returns a new instance of the ReviewRepository.
//
// Returns a new instance of the ReviewRepository.
func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{}
}

// Create is a function that creates a new review.
//
// review: The review to create.
//
// Returns an error if any.
func (*ReviewRepository) Create(review *models.Review) error {
	// Create the review
	err := mysql.Conn.Create(review).Error

	// Return an error if any
	if err != nil {
		log.Println("Error creating review: " + err.Error())

		return err
	}

	return nil
}

// GetCountUsingVendorID is a function that returns the count of reviews for vendor ID.
//
// vendorID: The vendor ID.
//
// Returns the count and an error if any.
func (*ReviewRepository) GetCountUsingVendorID(vendorID uint) (int64, error) {
	// Create new count variable
	var count int64

	// Get the count of courts by vendor ID
	err := mysql.Conn.Model(&models.Court{}).Where("vendor_id = ?", vendorID).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting count using vendor id: " + err.Error())

		return 0, err
	}

	return count, nil
}

// GetCountUsingVendorIDCourtType is a function that returns the count of reviews for vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the count and an error if any.
func (*ReviewRepository) GetCountUsingVendorIDCourtType(vendorID uint, courtType string) (int64, error) {
	// Create new count variable
	var count int64

	// Get the count of courts by vendor ID and court type
	err :=
		mysql.Conn.Model(&models.Review{}).Joins("CourtType").Where("vendor_id = ?", vendorID).Where("CourtType.type = ?", courtType).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting count using vendor id and court type: " + err.Error())

		return 0, err
	}

	return count, nil
}

// GetStarCountsUsingVendorID is a function that returns the star count by vendor ID.
//
// vendorID: The vendor ID.
//
// Returns the star count and an error if any.
func (*ReviewRepository) GetStarCountsUsingVendorID(vendorID uint) (*types.StarCountsMap, error) {
	// Create a new struct for the results
	var results struct {
		OneStar   int64
		TwoStar   int64
		ThreeStar int64
		FourStar  int64
		FiveStar  int64
	}

	// Get the courts by vendor ID and star
	err :=
		mysql.Conn.Model(&models.Review{}).Select(`
        COUNT(CASE WHEN rating = 1 THEN 1 END),
        COUNT(CASE WHEN rating = 2 THEN 1 END),
        COUNT(CASE WHEN rating = 3 THEN 1 END),
        COUNT(CASE WHEN rating = 4 THEN 1 END),
        COUNT(CASE WHEN rating = 5 THEN 1 END)
    `).
			Where("vendor_id = ?", vendorID).Scan(&results).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting star counts using vendor id: " + err.Error())

		return nil, err
	}

	return &types.StarCountsMap{
		1: results.OneStar,
		2: results.TwoStar,
		3: results.ThreeStar,
		4: results.FourStar,
		5: results.FiveStar,
	}, nil
}

// GetStarCountsUsingVendorIDCourtType is a function that returns the star count by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns map of the star count and an error if any.
func (*ReviewRepository) GetStarCountsUsingVendorIDCourtType(vendorID uint, courtType string) (*types.StarCountsMap, error) {
	// Create a new struct for the results
	var results struct {
		OneStar   int64
		TwoStar   int64
		ThreeStar int64
		FourStar  int64
		FiveStar  int64
	}

	// Get the courts by vendor ID and court type
	err :=
		mysql.Conn.Model(&models.Review{}).Joins("JOIN court_types on court_types.id = reviews.court_type_id").Select(`
        COUNT(CASE WHEN rating = 1 THEN 1 END),
        COUNT(CASE WHEN rating = 2 THEN 1 END),
        COUNT(CASE WHEN rating = 3 THEN 1 END),
        COUNT(CASE WHEN rating = 4 THEN 1 END),
        COUNT(CASE WHEN rating = 5 THEN 1 END)
    `).
			Where("vendor_id = ?", vendorID).Where("court_types.type = ?", courtType).Scan(&results).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting star counts using vendor id and court type: " + err.Error())

		return nil, err
	}

	return &types.StarCountsMap{
		1: results.OneStar,
		2: results.TwoStar,
		3: results.ThreeStar,
		4: results.FourStar,
		5: results.FiveStar,
	}, nil
}

// GetUsingVendorIDCourtType is a function that returns the reviews using the vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the reviews and an error if any.
func (*ReviewRepository) GetUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Review, error) {
	// reviews is a slice of maps containing the reviews of the court
	var reviews []models.Review

	// Get the reviews using the vendor ID and court type
	err :=
		mysql.Conn.Preload("User").Preload("Vendor").Joins("CourtType").Where("vendor_id = ?", vendorID).Where("CourtType.type = ?", courtType).Find(&reviews).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting reviews using vendor id and court type: " + err.Error())

		return nil, err
	}

	return &reviews, err
}

// GetUsingVendorID is a function that returns the reviews using the vendor ID.
//
// vendorID: The vendor ID.
//
// Returns the reviews and an error if any.
func (*ReviewRepository) GetUsingVendorID(vendorID uint) (*[]models.Review, error) {
	// reviews is a slice of maps containing the reviews of the court
	var reviews []models.Review

	// Get the reviews using the vendor ID
	err :=
		mysql.Conn.Preload("User").Preload("Vendor").Preload("CourtType").Where("vendor_id = ?", vendorID).Find(&reviews).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting reviews using vendor id: " + err.Error())

		return nil, err
	}

	return &reviews, err
}

// CheckUserHasReviewCourtType is a function that checks if the user has a review for the court type.
//
// userID: The user ID.
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns true if the user has a review for the court type.
func (*ReviewRepository) CheckUserHasReviewCourtType(userID uint, vendorID uint, courtType string) (bool, error) {
	// Create new count variable
	var count int64

	// Get the count of courts by vendor ID and court type
	err :=
		mysql.Conn.Model(&models.Review{}).Joins("CourtType").Where("user_id = ?", userID).Where("vendor_id = ?", vendorID).Where("CourtType.type = ?", courtType).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking user has review for court type: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

// GetUsingVendorIDCourtTypeRating is a function that returns the reviews using the vendor ID, court type, and rating.
//
// vendorID: The vendor ID.
// courtType: The court type.
// rating: The rating.
//
// Returns the reviews and an error if any.
func (*ReviewRepository) GetUsingVendorIDCourtTypeRating(vendorID uint, courtType string, rating int) (*[]models.Review, error) {
	// reviews is a slice of maps containing the reviews of the court
	var reviews []models.Review

	// Get the reviews using the vendor ID and court type
	err :=
		mysql.Conn.Preload("User").Preload("Vendor").Joins("CourtType").Where("vendor_id = ?", vendorID).Where("rating = ?", rating).Where("CourtType.type = ?", courtType).Find(&reviews).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting reviews using vendor id, court type, and rating: " + err.Error())

		return nil, err
	}

	return &reviews, err
}

// GetUsingVendorIDRating is a function that returns the reviews using the vendor ID and rating.
//
// vendorID: The vendor ID.
// rating: The rating.
//
// Returns the reviews and an error if any.
func (*ReviewRepository) GetUsingVendorIDRating(vendorID uint, rating int) (*[]models.Review, error) {
	// reviews is a slice of maps containing the reviews of the court
	var reviews []models.Review

	// Get the reviews using the vendor ID and court type
	err :=
		mysql.Conn.Preload("User").Preload("Vendor").Preload("CourtType").Where("vendor_id = ?", vendorID).Where("rating = ?", rating).Find(&reviews).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting reviews using vendor id and rating: " + err.Error())

		return nil, err
	}

	return &reviews, err
}

// GetAvgRatingUsingCourtTypeVendorID is a function that returns the average rating using
// the court type and vendor ID.
//
// courtType: The court type.
// vendorID: The vendor ID.
//
// Returns the average rating and an error if any.
func (*ReviewRepository) GetAvgRatingUsingCourtTypeVendorID(courtType string, vendorID uint) (float64, error) {
	// Create new count variable
	var avgRating float64

	// Get the count of courts by vendor ID and court type
	err :=
		mysql.Conn.Model(&models.Review{}).Joins("JOIN court_types ON court_types.id = reviews.court_type_id").Where("vendor_id = ?", vendorID).Select("COALESCE(AVG(rating), 0)").Scan(&avgRating).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting avg rating using court type and vendor id: " + err.Error())

		return 0, err
	}

	return avgRating, nil
}
