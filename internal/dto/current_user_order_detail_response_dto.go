package dto

// CurrentUserOrderDetailResponseDTO is a struct that represents the current user order
// detail response data transfer object.
type CurrentUserOrderDetailResponseDTO struct {
	// OrderDetail is the order detail.
	OrderDetail *CurrentUserOrderDetailDTO `json:"order_detail"`
}
