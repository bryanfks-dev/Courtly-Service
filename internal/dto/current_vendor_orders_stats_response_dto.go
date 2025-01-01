package dto

import (
	"main/core/types"
	"main/data/models"
)

// CurrentVendorOrdersStatsResponseDTO is a struct that represents the
// response of current vendor orders stats DTO.
type CurrentVendorOrdersStatsResponseDTO struct {
	// TotalOrders is the total number of orders.
	TotalOrders int64 `json:"total_orders"`

	// TotalOrdersToday is the total number of orders today.
	TotalOrdersToday int64 `json:"total_orders_today"`

	// RecentOrders is the recent orders.
	RecentOrders *[]CurrentVendorOrderDTO `json:"recent_orders"`
}

// FromMap is a function that converts the orders stats map to the current vendor orders stats response DTO.
//
// m: The orders stats map.
//
// Returns the current vendor orders stats response DTO.
func (c CurrentVendorOrdersStatsResponseDTO) FromMap(m *types.OrdersStatsMap) *CurrentVendorOrdersStatsResponseDTO {
	return &CurrentVendorOrdersStatsResponseDTO{
		TotalOrders:      (*m)["total_orders"].(int64),
		TotalOrdersToday: (*m)["total_orders_today"].(int64),
		RecentOrders:     CurrentVendorOrderDTO{}.FromModels((*m)["recent_orders"].(*[]models.Order)),
	}
}
