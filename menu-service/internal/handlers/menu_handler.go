package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/brandao07/cj-restaurant/menu-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/menu-service/internal/services"
	"github.com/labstack/echo/v4"
)

type MenuHandler struct {
	Service *services.MenuService
}

func NewMenuHandler(service *services.MenuService) *MenuHandler {
	return &MenuHandler{Service: service}
}

func (h *MenuHandler) CreateMenu(c echo.Context) error {
	var dto dtos.CreateMenuDTO

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}

	fmt.Printf("%+v\n", dto)

	menu, err := h.Service.CreateMenu(&dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create menu: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, menu)
}

func (h *MenuHandler) DeactivateMenu(c echo.Context) error {
	id := c.Param("id")
	menuID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid menu ID: " + err.Error()})
	}

	if err := h.Service.DeactivateMenu(menuID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to deactivate menu: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
