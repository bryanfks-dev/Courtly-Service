package repository

import (
	"log"
	"main/core/enums"
	"main/core/types"
	"main/data/models"
	"main/internal/providers/mysql"
)

// CourtRepository is a struct that defines the court repository.
type CourtRepository struct{}

// NewCourtRepository is a factory function that returns a new instance of the court repository.
//
// Returns a new instance of the court repository.
func NewCourtRepository() *CourtRepository {
	return &CourtRepository{}
}

// Create is a function that creates a new court.
//
// court: The court object.
//
// Returns an error if any.
func (*CourtRepository) Create(court *models.Court) error {
	// Create the court
	err := mysql.Conn.Create(court).Error

	if err != nil {
		log.Println("Error creating court: " + err.Error())

		return err
	}

	return nil
}

// GetNewestUsingVendorIDCourtType is a function that returns the newest court by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the court and an error if any.
func (*CourtRepository) GetNewestUsingVendorIDCourtType(vendorID uint, courtType string) (*models.Court, error) {
	// Create a new court object
	var court models.Court

	// Get the courts by vendor ID and court type
	err :=
		mysql.Conn.Preload("Vendor").Joins("CourtType").Where("vendor_id = ?", vendorID).Where("CourtType.type = ?", courtType).Order("created_at desc").First(&court).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting newest court using vendor ID and court type: " + err.Error())

		return nil, err
	}

	return &court, nil
}

// Get is a function that returns all the courts.
//
// Returns the courts and an error if any.
func (*CourtRepository) Get() (*[]models.Court, error) {
	// Create courts array
	var courts []models.Court

	// Subquery to get the minimum id for each vendor id and court type id
	subQuery :=
		mysql.Conn.Model(&models.Court{}).Select("MIN(id)").Group("vendor_id, court_type_id")

	// Get the courts
	err := mysql.Conn.Preload("Vendor").Preload("CourtType").
		Where("id IN (?)", subQuery).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting all courts: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// GetUsingVendorName is a function that returns all the courts by vendor name.
//
// vendorName: The vendor name.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetUsingVendorName(vendorName string) (*[]models.Court, error) {
	// Create courts array
	var courts []models.Court

	// Subquery to get the minimum id for each vendor id and court type id
	subQuery :=
		mysql.Conn.Model(&models.Court{}).Select("MIN(courts.id)").Joins("JOIN vendors ON vendors.id = courts.vendor_id").Where("vendors.name LIKE ?", "%"+vendorName+"%").Group("vendor_id, court_type_id")

	// Get the courts
	err := mysql.Conn.Preload("Vendor").Preload("CourtType").
		Where("id IN (?)", subQuery).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting all courts using vendor name: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// GetUsingCourtType is a function that returns all the courts by court type.
//
// courtType: The court type.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetUsingCourtType(courtType string) (*[]models.Court, error) {
	// Create courts array
	var courts []models.Court

	// Subquery to get the minimum id for each vendor id and court type id
	subQuery :=
		mysql.Conn.Model(&models.Court{}).Select("MIN(courts.id)").Joins("JOIN court_types ON court_types.id = courts.court_type_id").Where("court_types.type = ?", courtType).Group("vendor_id, court_type_id")

	// Get the courts
	err := mysql.Conn.Preload("Vendor").Preload("CourtType").
		Where("id IN (?)", subQuery).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting courts using court type: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// GetUsingCourtTypeVendorName is a function that returns all the courts by court type and vendor name.
//
// courtType: The court type.
// vendorName: The vendor name.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetUsingCourtTypeVendorName(courtType string, vendorName string) (*[]models.Court, error) {
	// Create courts array
	var courts []models.Court

	// Subquery to get the minimum id for each vendor id and court type id
	subQuery :=
		mysql.Conn.Model(&models.Court{}).Select("MIN(courts.id)").Joins("JOIN vendors ON vendors.id = courts.vendor_id").Joins("JOIN court_types ON court_types.id = courts.court_type_id").Where("vendors.name LIKE ?", "%"+vendorName+"%").Where("court_types.type = ?", courtType).Group("vendor_id, court_type_id")

	// Get the courts
	err := mysql.Conn.Preload("Vendor").Preload("CourtType").
		Where("id IN (?)", subQuery).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting all courts using vendor name and court type: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// GetUsingID is a function that returns the courts by ID.
//
// courtID: The court ID.
//
// Returns the courts and an error if any.
func (*CourtRepository) GetUsingID(courtID uint) (*models.Court, error) {
	// Create court with rating object
	var court models.Court

	// Get the courts by ID
	err :=
		mysql.Conn.Preload("Vendor").Preload("CourtType").Where("id = ?", courtID).Find(&court).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting court using id: " + err.Error())

		return nil, err
	}

	return &court, nil
}

// GetUsingVendorIDCourtType is a function that returns the courts by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns the vendor courts and an error if any.
func (*CourtRepository) GetUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Court, error) {
	// Create a new court object
	var courts []models.Court

	// Get the courts by vendor ID and court type
	err :=
		mysql.Conn.Preload("Vendor").Joins("CourtType").Where("vendor_id = ?", vendorID).Where("CourtType.type = ?", courtType).Find(&courts).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting courts using vendor id and court type: " + err.Error())

		return nil, err
	}

	return &courts, nil
}

// CheckExistsUsingVendorIDCourtType is a function that checks if the courts exist by vendor ID and court type.
//
// vendorID: The vendor ID.
// courtType: The court type.
//
// Returns a boolean and an error if any.
func (*CourtRepository) CheckExistUsingVendorIDCourtType(vendorID uint, courtType string) (bool, error) {
	// count is the number of courts
	var count int64

	// Get the courts by vendor ID and court type
	err :=
		mysql.Conn.Model(&models.Court{}).Joins("CourtType").Where("vendor_id = ?", vendorID).Where("CourtType.type = ?", courtType).Limit(1).Count(&count).Error

	// Return an error if any
	if err != nil {
		log.Println("Error checking court exist using vendor id and court type: " + err.Error())

		return false, err
	}

	return count > 0, nil
}

// GetCountsUsingVendorID is a function that returns the counts of the courts by vendor ID.
//
// vendorID: The vendor ID.
//
// Returns the counts of the courts map and an error if any.
func (*CourtRepository) GetCountsUsingVendorID(vendorID uint) (*types.CourtCountsMap, error) {
	// Create a new struct for the results
	var results struct {
		FootballCount   int64
		BasketballCount int64
		VolleyballCount int64
		TennisCount     int64
		BadmintonCount  int64
	}

	// Get the counts of the courts by vendor ID
	err := mysql.Conn.Model(&models.Court{}).Select(`
        COUNT(CASE WHEN court_type_id = ? THEN 1 END) AS football_count,
        COUNT(CASE WHEN court_type_id = ? THEN 1 END) AS basketball_count,
        COUNT(CASE WHEN court_type_id = ? THEN 1 END) AS volleyball_count,
        COUNT(CASE WHEN court_type_id = ? THEN 1 END) AS tennis_count,
        COUNT(CASE WHEN court_type_id = ? THEN 1 END) AS badminton_count
    `, enums.GetCourtTypeID(enums.Football.Label()), enums.GetCourtTypeID(enums.Basketball.Label()), enums.GetCourtTypeID(enums.Volleyball.Label()), enums.GetCourtTypeID(enums.Tennis.Label()), enums.GetCourtTypeID(enums.Badminton.Label())).
		Where("vendor_id = ?", vendorID).Scan(&results).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting court counts using vendor id: " + err.Error())

		return nil, err
	}

	return &types.CourtCountsMap{
		enums.Football.Label():   results.FootballCount,
		enums.Basketball.Label(): results.BasketballCount,
		enums.Volleyball.Label(): results.VolleyballCount,
		enums.Tennis.Label():     results.TennisCount,
		enums.Badminton.Label():  results.BadmintonCount,
	}, nil
}

// UpdateUsingVendorIDCourtType is a method to update court using the given vendor id and court type.
//
// vendorID: The vendor id
// courtType: The court type
// pricePerHour: The court price
//
// Return error if any
func (*CourtRepository) UpdateUsingVendorIDCourtType(vendorID uint, courtType string, pricePerHour float64) error {
	// Update court
	err := mysql.Conn.Model(models.Court{}).Where("vendor_id = ?", vendorID).Where("court_type_id = ?", enums.GetCourtTypeID(courtType)).Updates(models.Court{
		Price: pricePerHour,
	}).Error

	// Return error if any
	if err != nil {
		log.Println("Error updating court using vendor id and court type: " + err.Error())

		return err
	}

	return nil
}

// DeleteUsingCourtIDsVendorID is a function that deletes the courts using court IDs and vendor ID.
//
// courtIDs: The court IDs.
// vendorID: The vendor ID.
//
// Returns an error if any.
func (*CourtRepository) DeleteUsingCourtIDsVendorID(courtIDs []uint, vendorID uint) error {
	// Delete the courts
	err := mysql.Conn.Where("id IN (?)", courtIDs).Where("vendor_id = ?", vendorID).Delete(models.Court{}).Error

	// Return an error if any
	if err != nil {
		log.Println("Error deleting courts using court ids and vendor id: " + err.Error())

		return err
	}

	return nil
}
