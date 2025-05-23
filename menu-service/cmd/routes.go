package main

import (
    "github.com/brandao07/cj-restaurant/menu-service/internal/handlers"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
    MenuHandler *handlers.MenuHandler
}

func NewRouter(h *Handlers) *echo.Echo {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    api := e.Group("/api")
    menus := api.Group("/menus")
    menus.POST("", h.MenuHandler.CreateMenu)
    menus.PATCH("/:id/deactivate", h.MenuHandler.DeactivateMenu)

    return e
}