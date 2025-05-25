package services

import (
	"fmt"

	"github.com/brandao07/cj-restaurant/customer-service/internal/clients"
	"github.com/brandao07/cj-restaurant/customer-service/internal/common"
	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/customer-service/internal/repositories"
)

type CustomerService struct {
	CustomerRepository repositories.CustomerRepository
	MenuServiceClient  clients.MenuServiceClient
	OrderServiceClient clients.OrderServiceClient
}

func NewCustomerService(customerRepository repositories.CustomerRepository, menuServiceURL string, orderServiceURL string) *CustomerService {
	menuServiceClient := clients.NewMenuServiceClient(menuServiceURL)
	orderServiceClient := clients.NewOrderServiceClient(orderServiceURL)
	return &CustomerService{
		CustomerRepository: customerRepository,
		MenuServiceClient:  menuServiceClient,
		OrderServiceClient: orderServiceClient,
	}
}

func (s *CustomerService) CreateCustomer(dto dtos.CreateCustomerDTO) (*dtos.Customer, error) {
	// Connect to Menu-Service Microservice to fetch the menus
	menus, err := s.MenuServiceClient.GetMenus(dto.MenuIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch menus: %w", err)
	}

	var totalPrice float32
	menuIds := make([]int, len(menus))

	// Calculate the total price and collect menu IDs
	for _, menu := range menus {
		totalPrice += menu.Price
		menuIds = append(menuIds, menu.ID)
	}

	//Connect to Order-Service Microservice to create a new order
	createOrderDTO := dtos.CreateRestaurantOrderDTO{
		MenuIDs: menuIds,
		Total:   totalPrice,
	}

	order, err := s.OrderServiceClient.CreateOrder(createOrderDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	dto.OrderID = order.ID

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
