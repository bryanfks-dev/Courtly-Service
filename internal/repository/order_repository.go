package repository

import (
	"log"
	"main/data/models"
	"main/internal/providers/mysql"

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

// GetUsingUserID is a method that returns the orders by the given user ID.
//
// userID: The ID of the user.
//
// Returns the orders and an error if any.
func (*OrderRepository) GetUsingUserID(userID uint) (*[]models.Order, error) {
	// orders is a placeholder for the orders
	var orders []models.Order

	// Get the orders from the database
	err :=
		mysql.Conn.Preload("PaymentMethod").Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Where("bookings.user_id = ?", userID).Group("orders.id").Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting orders using user id: " + err.Error())

		return nil, err
	}

	return &orders, nil
}

// GetUsingUserIDCourtType is a method that returns the orders by the given user ID and court type.
//
// userID: The ID of the user.
// courtType: The type of the court.
//
// Returns the orders and an error if any.
func (*OrderRepository) GetUsingUserIDCourtType(userID uint, courtType string) (*[]models.Order, error) {
	// orders is a placeholder for the orders
	var orders []models.Order

	// Get the orders from the database
	err :=
		mysql.Conn.Preload("PaymentMethod").Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Joins("JOIN courts ON courts.id = bookings.court_id").
			Joins("JOIN court_types ON court_types.id = courts.court_type_id").
			Where("bookings.user_id = ?", userID).Group("orders.id").
			Where("court_types.type = ?", courtType).
			Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting orders using user id and court type: " + err.Error())

		return nil, err
	}

	return &orders, nil
}

// GetUsingID is a method that returns the order by the given ID.
//
// orderID: The ID of the order.
// userID: The ID of the user.
//
// Returns the order and an error if any.
func (*OrderRepository) GetUsingIDUserID(orderID uint, userID uint) (*models.Order, error) {
	// order is a placeholder for the order
	var order models.Order

	// Get the order from the database
	err :=
		mysql.Conn.Preload("PaymentMethod").Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Where("orders.id = ?", orderID).Where("bookings.user_id = ?", userID).
			First(&order).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting order using order id and user id: " + err.Error())

		return nil, err
	}

	return &order, nil
}
