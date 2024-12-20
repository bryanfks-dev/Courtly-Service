package usecases

import (
	"main/data/models"
	"main/domain/entities"
	"main/internal/repository"
	"main/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// OrderUseCase is a struct that defines the OrderUseCase
type OrderUseCase struct {
	AuthUseCase     *AuthUseCase
	OrderRepository *repository.OrderRepository
}

// NewOrderUseCase is a function that returns a new OrderUseCase
//
// a: The AuthUseCase
// o: The OrderRepository
//
// Returns a pointer to the OrderUseCase struct
func NewOrderUseCase(a *AuthUseCase, o *repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		AuthUseCase:     a,
		OrderRepository: o,
	}
}

// GetCurrentUserOrders is a method that gets the current user order from the database.
//
// token: The JWT token.
// courtType: The court type.
//
// Returns the orders and an error if any.
func (o *OrderUseCase) GetCurrentUserOrders(token *jwt.Token, courtType *string) (*[]models.Order, error) {
	// Get the user ID from the JWT
	claims := o.AuthUseCase.DecodeToken(token)

	// Check if the court type is not empty
	if courtType != nil && utils.IsBlank(*courtType) {
		// Get the orders using the user ID
		return o.OrderRepository.GetUsingUserID(claims.Id)
	}

	// Get the orders using the user ID
	return o.OrderRepository.GetUsingUserIDCourtType(claims.Id, *courtType)
}

// GetCurrentUserOrderDetail is a method that gets the current user order detail from the database.
//
// token: The JWT token.
// orderID: The order ID.
//
// Returns the order and an error if any.
func (o *OrderUseCase) GetCurrentUserOrderDetail(token *jwt.Token, orderID uint) (*models.Order, *entities.ProcessError) {
	// Get the user ID from the JWT
	claims := o.AuthUseCase.DecodeToken(token)

	// Get the order using the user ID and order ID
	order, err := o.OrderRepository.GetUsingIDUserID(orderID, claims.Id)

	// Return an error if any
	if err == gorm.ErrRecordNotFound {
		return nil, &entities.ProcessError{
			ClientError: true,
			Message:     "Order not found",
		}
	}

	if err != nil {
		return nil, &entities.ProcessError{
			ClientError: false,
			Message:     "Failed to get user order",
		}
	}

	return order, nil
}
