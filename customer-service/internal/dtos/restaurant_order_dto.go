package dtos

// CreateRestaurantOrderDTO represents a restaurant order creation request
type CreateRestaurantOrderDTO struct {
	MenuIDs []int   `json:"menu_ids"`
	Total   float32 `json:"total"`
}

func NewCreateRestaurantOrderDTO(menuIDs []int, total float32, status string) *CreateRestaurantOrderDTO {
	return &CreateRestaurantOrderDTO{
		MenuIDs: menuIDs,
		Total:   total,
	}
}

// RestaurantOrder represents a restaurant order response
type RestaurantOrder struct {
	ID      int     `json:"id"`
	MenuIDs []int   `json:"menu_ids"`
	Total   float32 `json:"total"`
	Status  string  `json:"status"`
}

func NewRestaurantOrder(id int, menuIDs []int, total float32, status string) *RestaurantOrder {
	return &RestaurantOrder{
		ID:      id,
		MenuIDs: menuIDs,
		Total:   total,
		Status:  status,
	}
}
