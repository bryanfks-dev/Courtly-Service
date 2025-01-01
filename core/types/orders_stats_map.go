package types

// OrdersStatsMap is a map of order stats
// OrdersStatsMap consists of followings:
// {
//     "total_orders": ...,
//     "total_orders_today": ...,
//     "recent_orders": [...]
// }
type OrdersStatsMap map[string]any
