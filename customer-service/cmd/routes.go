package main

import (
	"net/http"

	"github.com/brandao07/cj-restaurant/customer-service/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	CustomerHandler        *handlers.CustomerHandler
	RestaurantTableHandler *handlers.RestaurantTableHandler
}

func NewRouter(h *Handlers) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	api := e.Group("/api")
	customers := api.Group("/customers")
	customers.POST("", h.CustomerHandler.CreateCustomer)
	customers.PATCH("/:id", h.CustomerHandler.UpdateCustomer)

	restaurant_tables := api.Group("/restaurant-tables")
	restaurant_tables.GET("/pending-payment", h.RestaurantTableHandler.GetPendingPaymentTables)
	restaurant_tables.GET("/pending-delivery", h.RestaurantTableHandler.GetPendingDeliveryTables)

	return e
}
