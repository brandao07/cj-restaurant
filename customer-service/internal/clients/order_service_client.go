package clients

import (
	"fmt"

	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/go-resty/resty/v2"
)

type OrderServiceClient interface {
	CreateOrder(dto dtos.CreateRestaurantOrderDTO) (*dtos.RestaurantOrder, error)
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

func (c *orderServiceClient) CreateOrder(dto dtos.CreateRestaurantOrderDTO) (*dtos.RestaurantOrder, error) {
	resp, err := c.client.R().
		SetBody(dto).
		SetResult(&dtos.RestaurantOrder{}).
		Post("/orders")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("menu service returned status %d", resp.StatusCode())
	}

	order := *resp.Result().(*dtos.RestaurantOrder)
	return &order, nil
}
