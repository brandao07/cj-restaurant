package handlers

import (
	"net/http"

	"github.com/brandao07/cj-restaurant/customer-service/internal/services"
	"github.com/labstack/echo/v4"
)

type RestaurantTableHandler struct {
	Service *services.RestaurantTableService
}

func NewRestaurantTableHandler(service *services.RestaurantTableService) *RestaurantTableHandler {
	return &RestaurantTableHandler{Service: service}
}

// GetPendingPaymentTables godoc
// @Summary      Get tables with pending payment
// @Description  Returns a list of restaurant tables where customers have not yet paid
// @Tags         restaurant-tables
// @Produce      json
// @Success      200  {array}   dtos.RestaurantTable
// @Failure      500  {string}  string  "Failed to get pending payment tables"
// @Router       /restaurant-tables/pending-payment [get]
func (h *RestaurantTableHandler) GetPendingPaymentTables(c echo.Context) error {
	tables, err := h.Service.GetPendingPaymentTables()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get pending payment tables: "+err.Error())
	}

	return c.JSON(http.StatusOK, tables)
}

// GetPendingDeliveryTables godoc
// @Summary      Get tables with pending delivery
// @Description  Returns a list of restaurant tables where orders are pending delivery
// @Tags         restaurant-tables
// @Produce      json
// @Success      200  {array}   dtos.RestaurantTable
// @Failure      500  {string}  string  "Failed to get pending delivery tables"
// @Router       /restaurant-tables/pending-delivery [get]
func (h *RestaurantTableHandler) GetPendingDeliveryTables(c echo.Context) error {
	tables, err := h.Service.GetPendingDeliveryTables()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get pending delivery tables: "+err.Error())
	}

	return c.JSON(http.StatusOK, tables)
}
