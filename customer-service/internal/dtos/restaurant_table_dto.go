package dtos

// RestaurantTable represents a restaurant table response
type RestaurantTable struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func NewRestaurantTable(id int, status string) *RestaurantTable {
	return &RestaurantTable{
		ID:     id,
		Status: status,
	}
}
