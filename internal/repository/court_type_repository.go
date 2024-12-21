package repository

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"
)

// CourtTypeRepository is a struct that defines the CourtTypeRepository
type CourtTypeRepository struct{}

// NewCourtTypeRepository is a function that returns a new CourtTypeRepository
//
// Returns a pointer to the CourtTypeRepository struct
func NewCourtTypeRepository() *CourtTypeRepository {
	return &CourtTypeRepository{}
}

// GetUsingType is a method that returns the court type by the given type.
//
// courtType: The type of the court.
//
// Returns the court type and an error if any.
func (*CourtTypeRepository) GetUsingType(courtType string) (*models.CourtType, error) {
	// model is a placeholder for the court type
	var model models.CourtType

	// Get the court type from the database
	err := mysql.Conn.Where("type = ?", courtType).First(&model).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting court type using type: " + err.Error())

		return nil, err
	}

	return &model, nil
}
