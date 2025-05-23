package handlers

import (
	"github.com/brandao07/cj-restaurant/order-service/internal/services"
)

type RestaurantOrderHandler struct {
	Service *services.RestaurantOrderService
}

func NewRestaurantOrderHandler(service *services.RestaurantOrderService) *RestaurantOrderHandler {
	return &RestaurantOrderHandler{Service: service}
}
