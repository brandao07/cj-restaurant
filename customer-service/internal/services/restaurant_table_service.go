package services

import (
	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/customer-service/internal/repositories"
)

type RestaurantTableService struct {
	RestaurantTableRepository repositories.RestaurantTableRepository
}

func NewRestaurantTableService(restaurantTableRepository repositories.RestaurantTableRepository) *RestaurantTableService {
	return &RestaurantTableService{
		RestaurantTableRepository: restaurantTableRepository,
	}
}

func (s *RestaurantTableService) GetPendingPaymentTables() ([]dtos.RestaurantTable, error) {
	tables, err := s.RestaurantTableRepository.GetPendingPaymentTables()
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (s *RestaurantTableService) GetPendingDeliveryTables() ([]dtos.RestaurantTable, error) {
	tables, err := s.RestaurantTableRepository.GetPendingDeliveryTables()
	if err != nil {
		return nil, err
	}
	return tables, nil
}
