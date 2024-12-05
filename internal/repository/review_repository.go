package repository

import (
	"log"
	"main/data/models"
	"main/domain/entities"
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
	err := mysql.Conn.Model(&models.Review{}).Preload("CourtType", "type = ?", courtType).Where("vendor_id = ?", vendorID).Count(&count).Error

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
func (*ReviewRepository) GetStarCountsUsingVendorID(vendorID uint) (*entities.ReviewStarsCount, error) {
	// Create a new count variable
	var counts entities.ReviewStarsCount

	// Get the courts by vendor ID and star
	err := mysql.Conn.Model(&models.Review{}).Select(`
        COUNT(CASE WHEN stars = 1 THEN 1 END),
        COUNT(CASE WHEN stars = 2 THEN 1 END),
        COUNT(CASE WHEN stars = 3 THEN 1 END),
        COUNT(CASE WHEN stars = 4 THEN 1 END),
        COUNT(CASE WHEN stars = 5 THEN 1 END)
    `).
		Where("vendor_id = ?", vendorID).Scan(&counts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting star counts using vendor id: " + err.Error())

		return nil, err
	}

	return &counts, nil
}

// GetStarCountsUsingVendorIDCourtType is a function that returns the star count by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the star count and an error if any.
func (*ReviewRepository) GetStarCountsUsingVendorIDCourtType(vendorID uint, courtType string) (*entities.ReviewStarsCount, error) {
	// Create a new count variable
	var counts entities.ReviewStarsCount

	// Get the courts by vendor ID and court type
	err := mysql.Conn.Model(&models.Review{}).Preload("CourtType", "type = ?", courtType).Select(`
        COUNT(CASE WHEN stars = 1 THEN 1 END),
        COUNT(CASE WHEN stars = 2 THEN 1 END),
        COUNT(CASE WHEN stars = 3 THEN 1 END),
        COUNT(CASE WHEN stars = 4 THEN 1 END),
        COUNT(CASE WHEN stars = 5 THEN 1 END)
    `).
		Where("vendor_id = ?", vendorID).Scan(&counts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting star counts using vendor id and court type: " + err.Error())

		return nil, err
	}

	return &counts, nil
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
	err := mysql.Conn.Preload("CourtType", "type = ?", courtType).Where("vendor_id = ?", vendorID).Find(&reviews).Error

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
	err := mysql.Conn.Where("vendor_id = ?", vendorID).Find(&reviews).Error

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
	err := mysql.Conn.Model(&models.Review{}).Where("user_id = ? AND vendor_id = ? AND court_type = ?", userID, vendorID, courtType).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking user has review for court type: " + err.Error())

		return false, err
	}

	return count > 0, nil
}
