package repository

import (
	"log"
	"main/core/enums"
	"main/data/models"
	"main/internal/providers/mysql"
	"time"

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
		mysql.Conn.Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Where("bookings.user_id = ?", userID).Group("orders.id").
			Order("orders.created_at DESC").
			Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting orders using user id: " + err.Error())

		return nil, err
	}

	return &orders, nil
}

// GetUsingVendorID is a method that returns the orders by the given vendor ID.
//
// vendorID: The ID of the vendor.
//
// Returns the orders and an error if any.
func (*OrderRepository) GetUsingVendorID(vendorID uint) (*[]models.Order, error) {
	// orders is a placeholder for the orders
	var orders []models.Order

	// Get the orders from the database
	err :=
		mysql.Conn.Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Preload("Bookings.User").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Where("bookings.vendor_id = ?", vendorID).Group("orders.id").
			Order("orders.created_at DESC").
			Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting orders using vendor id: " + err.Error())

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
		mysql.Conn.Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Joins("JOIN courts ON courts.id = bookings.court_id").
			Joins("JOIN court_types ON court_types.id = courts.court_type_id").
			Where("bookings.user_id = ?", userID).Group("orders.id").
			Where("court_types.type = ?", courtType).
			Order("orders.created_at desc").
			Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting orders using user id and court type: " + err.Error())

		return nil, err
	}

	return &orders, nil
}

// GetUsingVendorIDCourtType is a method that returns the orders by the given vendor ID and court type.
//
// vendorID: The ID of the vendor.
// courtType: The type of the court.
//
// Returns the orders and an error if any.
func (*OrderRepository) GetUsingVendorIDCourtType(vendorID uint, courtType string) (*[]models.Order, error) {
	// orders is a placeholder for the orders
	var orders []models.Order

	// Get the orders from the database
	err :=
		mysql.Conn.Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Preload("Bookings.User").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Joins("JOIN courts ON courts.id = bookings.court_id").
			Joins("JOIN court_types ON court_types.id = courts.court_type_id").
			Where("bookings.vendor_id = ?", vendorID).Group("orders.id").
			Where("court_types.type = ?", courtType).
			Order("orders.created_at desc").
			Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting orders using user id and court type: " + err.Error())

		return nil, err
	}

	return &orders, nil
}

// GetUsingID is a method that returns the order by the given order ID.
//
// orderID: The ID of the order.
//
// Returns the order and an error if any.
func (*OrderRepository) GetUsingID(orderID uint) (*models.Order, error) {
	// order is a placeholder for the order
	var order models.Order

	// Get the order from the database
	err :=
		mysql.Conn.Preload("Bookings", func(db *gorm.DB) *gorm.DB {
			return db.Order("Bookings.book_start_time ASC")
		}).Preload("Bookings.Court.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Where("orders.id = ?", orderID).
			First(&order).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting order using order id: " + err.Error())

		return nil, err
	}

	return &order, nil
}

// UpdatePaymentTokenUsingID is a method that updates the payment token using the given order ID.
//
// tx: The database transaction.
// paymentToken: The payment token.
// orderID: The ID of the order.
//
// Returns an error if any.
func (*OrderRepository) UpdatePaymentTokenUsingID(tx *gorm.DB, paymentToken string, orderID uint) error {
	// Update the payment token using the order ID
	err := tx.Model(&models.Order{}).Where("id = ?", orderID).Update("payment_token", paymentToken).Error

	// Return an error if any
	if err != nil {
		log.Println("Error updating payment token using order id: " + err.Error())

		return err
	}

	return nil
}

// UpdatePaymentStatusUsingID is a method that updates the payment status using the given order ID.
//
// orderID: The ID of the order.
// status: The status of the payment.
//
// Returns an error if any.
func (*OrderRepository) UpdatePaymentStatusUsingID(orderID uint, status string) error {
	// Update the payment status using the order ID
	err := mysql.Conn.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error

	// Return an error if any
	if err != nil {
		log.Println("Error updating payment status using order id: " + err.Error())

		return err
	}

	return nil
}

// GetTotalUsingVendorID is a method that used to get vendor total order using the
// given vendor ID.
//
// vendorID: the id of the vendor
//
// Returns order count and error if any
func (*OrderRepository) GetTotalUsingVendorID(vendorID uint) (*int64, error) {
	// count is a placeholder for the count
	var count int64

	// Get the orders count from the database
	err :=
		mysql.Conn.Model(&models.Order{}).Joins("JOIN bookings ON bookings.order_id = orders.id").Where("bookings.vendor_id = ?", vendorID).Where("orders.status = ?", enums.Success.Label()).Group("bookings.order_id").Count(&count).Error

	if err != nil {
		log.Println("Error getting order total using vendor id: " + err.Error())

		return nil, err
	}

	return &count, err
}

// GetTotalTodayUsingVendorID is a method that used to get vendor total order
// today using the given vendor id.
//
// vendorID: the id of the vendor
//
// Returns the total order today and error if any
func (*OrderRepository) GetTotalTodayUsingVendorID(vendorID uint) (*int64, error) {
	// count is a placeholder for the count
	var count int64

	// Get the current date
	today := time.Now().Format("2006-01-02")

	// Get the orders count from the database
	err :=
		mysql.Conn.Model(&models.Order{}).Joins("JOIN bookings ON bookings.order_id = orders.id").Where("orders.status = ?", enums.Success.Label()).Where("vendor_id = ?", vendorID).Where("DATE(created_at) = ?", today).Count(&count).Error

	if err != nil {
		log.Println("Error getting order total today using vendor id: " + err.Error())

		return nil, err
	}

	return &count, err
}

// GetNLatestUsingVendorID is a method to get the n latest order using the given
// vendor id.
//
// vendorID: the id of the vendor
// n: the number of order to fetch
//
// Return n of latest order and error if any
func (*OrderRepository) GetNLatestUsingVendorID(vendorID uint, n int) (*[]models.Order, error) {
	// orders is a placeholder for the orders
	var orders []models.Order

	// Get the orders from the database
	err :=
		mysql.Conn.Preload("Bookings").Preload("Bookings.Vendor").
			Preload("Bookings.Court").Preload("Bookings.Court.CourtType").
			Preload("Bookings.User").
			Joins("JOIN bookings ON bookings.order_id = orders.id").
			Where("bookings.vendor_id = ?", vendorID).Group("orders.id").
			Order("orders.created_at DESC").Limit(n).
			Find(&orders).Error

	// Return an error if any
	if err != nil {
		log.Println("Error getting n latest orders using vendor id: " + err.Error())

		return nil, err
	}

	return &orders, nil
}
