package services

import (
	"github.com/brandao07/cj-restaurant/order-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/order-service/internal/repositories"
)

type RestaurantOrderService struct {
	RestaurantOrderRepository repositories.RestaurantOrderRepository
}

func NewRestaurantOrderService(restaurantOrderRepository repositories.RestaurantOrderRepository) *RestaurantOrderService {
	return &RestaurantOrderService{
		RestaurantOrderRepository: restaurantOrderRepository,
	}
}

func (s *RestaurantOrderService) CreateRestaurantOrder(dto *dtos.CreateRestaurantOrderDTO) (*dtos.RestaurantOrder, error) {
	order, err := s.RestaurantOrderRepository.CreateRestaurantOrder(dto)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *RestaurantOrderService) UpdateRestaurantOrder(dto dtos.UpdateRestaurantOrderDTO) error {
	err := s.RestaurantOrderRepository.UpdateRestaurantOrder(dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *RestaurantOrderService) GetRestaurantOrderByID(id int) (*dtos.RestaurantOrder, error) {
	order, err := s.RestaurantOrderRepository.GetRestaurantOrderByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
