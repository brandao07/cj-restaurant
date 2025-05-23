package services

import "github.com/brandao07/cj-restaurant/order-service/internal/repositories"

type RestaurantOrderService struct {
	RestaurantOrderRepository repositories.RestaurantOrderRepository
}

func NewRestaurantOrderService(restaurantOrderRepository repositories.RestaurantOrderRepository) *RestaurantOrderService {
	return &RestaurantOrderService{
		RestaurantOrderRepository: restaurantOrderRepository,
	}
}
