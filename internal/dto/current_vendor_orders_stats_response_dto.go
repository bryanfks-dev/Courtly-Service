package dto

// CurrentVendorOrdersStatsResponseDTO is a struct that represents the
// response of current vendor orders stats DTO.
type CurrentVendorOrdersStatsResponseDTO struct {
	// TotalOrders is the total number of orders.
	TotalOrders int `json:"total_orders"`

	// TotalOrdersToday is the total number of orders today.
	TotalOrdersToday int `json:"total_orders_today"`

	// RecentOrders is the recent orders.
	RecentOrders *[]CurrentVendorOrderDTO `json:"recent_orders"`
}
