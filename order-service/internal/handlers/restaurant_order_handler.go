package handlers

import (
	"net/http"
	"strconv"

	"github.com/brandao07/cj-restaurant/order-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/order-service/internal/services"
	"github.com/labstack/echo/v4"
)

type RestaurantOrderHandler struct {
	Service *services.RestaurantOrderService
}

func NewRestaurantOrderHandler(service *services.RestaurantOrderService) *RestaurantOrderHandler {
	return &RestaurantOrderHandler{Service: service}
}

// CreateRestaurantOrder godoc
// @Summary      Create a new restaurant order
// @Description  Creates a new restaurant order with the provided menu IDs and total.
// @Tags         restaurant_orders
// @Accept       json
// @Produce      json
// @Param        order  body      dtos.CreateRestaurantOrderDTO  true  "Restaurant Order Data"
// @Success      201    {object}  dtos.RestaurantOrder
// @Router       /restaurant-orders [post]
func (h *RestaurantOrderHandler) CreateRestaurantOrder(c echo.Context) error {
	var dto dtos.CreateRestaurantOrderDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	order, err := h.Service.CreateRestaurantOrder(&dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create restaurant order: "+err.Error())
	}

	return c.JSON(http.StatusCreated, order)
}

// UpdateRestaurantOrder godoc
// @Summary      Update a restaurant order status
// @Description  Updates the status of an existing restaurant order.
// @Tags         restaurant_orders
// @Accept       json
// @Produce      json
// @Param        order  body      dtos.UpdateRestaurantOrderDTO  true  "Restaurant Order Update Data"
// @Success      200    "Order updated successfully"
// @Router       /restaurant-orders [put]
func (h *RestaurantOrderHandler) UpdateRestaurantOrder(c echo.Context) error {
	var dto dtos.UpdateRestaurantOrderDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body: "+err.Error())
	}

	if err := h.Service.UpdateRestaurantOrder(dto); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update restaurant order: "+err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// GetRestaurantOrderByID godoc
// @Summary      Get restaurant order by ID
// @Description  Retrieves a restaurant order by its unique ID.
// @Tags         restaurant_orders
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Restaurant Order ID"
// @Success      200  {object}  dtos.RestaurantOrder
// @Router       /restaurant-orders/{id} [get]
func (h *RestaurantOrderHandler) GetRestaurantOrderByID(c echo.Context) error {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid restaurant order ID: "+err.Error())
	}

	order, err := h.Service.GetRestaurantOrderByID(orderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get restaurant order: "+err.Error())
	}

	if order == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Restaurant order not found")
	}

	return c.JSON(http.StatusOK, order)
}
