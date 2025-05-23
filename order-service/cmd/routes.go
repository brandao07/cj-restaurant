package main

import (
	"net/http"

	"github.com/brandao07/cj-restaurant/order-service/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handlers struct {
	RestaurantOrderHandler *handlers.RestaurantOrderHandler
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

	// api := e.Group("/api")
	// orders := api.Group("/orders")
	// orders.POST("", h.OrderHandler.CreateOrder)
	// orders.PATCH("/:id", h.OrderHandler.UpdateOrder)

	return e
}
