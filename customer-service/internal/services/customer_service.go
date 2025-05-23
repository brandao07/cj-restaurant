package services

import (
	"fmt"

	"github.com/brandao07/cj-restaurant/customer-service/internal/common"
	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/customer-service/internal/repositories"
)

type CustomerService struct {
	CustomerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) *CustomerService {
	return &CustomerService{
		CustomerRepository: customerRepository,
	}
}

func (s *CustomerService) CreateCustomer(dto dtos.CreateCustomerDTO) (*dtos.Customer, error) {
	//TODO: Implement the logic to create a new customer
	// Connect to Order-Service Microservice to create a new order
	// order, err := s.OrderService.CreateOrder(dto)...

	customer, err := s.CustomerRepository.CreateCustomer(dto)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *CustomerService) UpdateCustomer(dto dtos.UpdateCustomerDTO) error {
	oldCustomer, err := s.CustomerRepository.GetCustomerByID(dto.ID)
	if err != nil {
		return err
	}

	// Status transition logic
	statusOrder := map[string]int{
		common.CustomerStatusWaiting:      1,
		common.CustomerStatusEating:       2,
		common.CustomerStatusNeedsPayment: 3,
		common.CustomerStatusPaid:         4,
	}

	oldStatusOrder := statusOrder[oldCustomer.Status]
	newStatusOrder := statusOrder[dto.Status]

	if newStatusOrder != oldStatusOrder+1 {
		return fmt.Errorf("invalid status transition from %s to %s", oldCustomer.Status, dto.Status)
	}

	err = s.CustomerRepository.UpdateCustomer(dto)
	if err != nil {
		return err
	}
	return nil
}
