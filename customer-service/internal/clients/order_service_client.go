package clients

import (
	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/go-resty/resty/v2"
)

type OrderServiceClient interface {
	CreateOrder(dto dtos.CreateRestaurantOrderDTO) (dtos.RestaurantOrder, error)
}

type orderServiceClient struct {
	client *resty.Client
}

func NewOrderServiceClient(baseURL string) OrderServiceClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json")

	return &orderServiceClient{
		client: client,
	}
}

func (c *orderServiceClient) CreateOrder(dto dtos.CreateRestaurantOrderDTO) (dtos.RestaurantOrder, error) {
	// implementation of CreateOrder method
	panic("CreateOrder method not implemented")
}
