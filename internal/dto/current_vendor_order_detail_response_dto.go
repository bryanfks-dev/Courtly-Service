package dto

// CurrentVendorOrderDetailResponseDTO is a struct that represents the current vendor order
// detail response data transfer object.
type CurrentVendorOrderDetailResponseDTO struct {
	// OrderDetail is the order detail.
	OrderDetail *CurrentVendorOrderDetailDTO `json:"order_detail"`
}
