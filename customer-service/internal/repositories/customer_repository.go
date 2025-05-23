package repositories

import "github.com/brandao07/cj-restaurant/customer-service/internal/dtos"

type CustomerRepository interface {
	CreateCustomer(dto dtos.CreateCustomerDTO) (*dtos.Customer, error)
	UpdateCustomer(dto dtos.UpdateCustomerDTO) error
	GetCustomerByRestaurantTableID(restaurantTableID int) (*dtos.Customer, error)
	GetCustomerByID(id int) (*dtos.Customer, error)
}
