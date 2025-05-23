package repositories

import "github.com/brandao07/cj-restaurant/customer-service/internal/dtos"

type RestaurantTableRepository interface {
	GetPendingPaymentTables() ([]dtos.RestaurantTable, error)
	GetPendingDeliveryTables() ([]dtos.RestaurantTable, error)
}
