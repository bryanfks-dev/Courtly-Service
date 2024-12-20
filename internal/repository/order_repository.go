package repository

import (
	"log"
	"main/data/models"

	"gorm.io/gorm"
)

// OrderRepository is a struct that defines the OrderRepository
type OrderRepository struct{}

// NewOrderRepository is a function that returns a new OrderRepository
//
// Returns a pointer to the OrderRepository struct
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

// Create is a method that creates an order in the database.
//
// tx: The database transaction.
// order: The order to create.
//
// Returns an error if any.
func (*OrderRepository) Create(tx *gorm.DB, order *models.Order) error {
	err := tx.Create(order).Error

	// Return an error if any
	if err != nil {
		log.Println("Error creating order: " + err.Error())

		return err
	}

	return nil
}
