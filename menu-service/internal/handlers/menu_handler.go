package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

// CreateMenu godoc
// @Summary      Create a new menu
// @Description  Creates a new menu with the given items and price
// @Tags         menus
// @Accept       json
// @Produce      json
// @Param        menu  body  dtos.CreateMenuDTO  true  "Menu to create"
// @Success      201   {object}  dtos.Menu
// @Router       /api/menus [post]
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

// DeactivateMenu godoc
// @Summary      Deactivate a menu
// @Description  Sets deleted_at for a menu
// @Tags         menus
// @Param        id   path      int  true  "Menu ID"
// @Success      200  {object}  map[string]string
// @Router       /api/menus/{id}/deactivate [patch]
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

// GetMenus godoc
//
// @Summary      Get menus by IDs or all active menus
// @Description  If the "ids" query parameter is provided (e.g. ?ids=1,2,3), fetches the menus with those IDs.
//
//	Otherwise, returns all active menus.
//
// @Tags         menus
// @Accept       json
// @Produce      json
// @Param        ids   query     []int  false  "List of Menu IDs to fetch (comma-separated)"
// @Success      200   {array}   dtos.Menu
// @Failure      400   {object}  map[string]string  "Invalid request"
// @Failure      500   {object}  map[string]string  "Internal server error"
// @Router       /menus [get]
func (h *MenuHandler) GetMenus(c echo.Context) error {
	if c.QueryParam("ids") != "" {
		return h.getMenusByIDs(c)
	}
	return h.getAllActiveMenus(c)
}

func (h *MenuHandler) getMenusByIDs(c echo.Context) error {
	idsParam := c.QueryParam("ids")
	idStrs := strings.Split(idsParam, ",")

	var ids []int
	for _, s := range idStrs {
		id, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format: " + s})
		}
		ids = append(ids, id)
	}

	menus, err := h.Service.GetMenusByIDs(ids)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch menus: " + err.Error()})
	}
	return c.JSON(http.StatusOK, menus)
}

func (h *MenuHandler) getAllActiveMenus(c echo.Context) error {
	menus, err := h.Service.GetAllActiveMenus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch active menus: " + err.Error()})
	}
	return c.JSON(http.StatusOK, menus)
}
