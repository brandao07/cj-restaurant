package repositories

import "github.com/brandao07/cj-restaurant/order-service/internal/dtos"

type RestaurantOrderRepository interface {
	CreateRestaurantOrder(dto *dtos.CreateRestaurantOrderDTO) (*dtos.RestaurantOrder, error)
	UpdateRestaurantOrder(dto dtos.UpdateRestaurantOrderDTO) error
	GetRestaurantOrderByID(id int) (*dtos.RestaurantOrder, error)
}
