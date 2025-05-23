package handlers

import (
	"net/http"

	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/customer-service/internal/services"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	Service *services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{Service: service}
}

// CreateCustomer godoc
// @Summary      Create a new customer
// @Description  Creates a new customer for a restaurant table
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customer  body      dtos.CreateCustomerDTO  true  "Customer data"
// @Success      201       {object}  dtos.Customer
// @Failure      400       {string}  string  "Invalid request"
// @Failure      500       {string}  string  "Failed to create customer"
// @Router       /customers [post]
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	dto := dtos.CreateCustomerDTO{}
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	customer, err := h.Service.CreateCustomer(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create customer: "+err.Error())
	}

	return c.JSON(http.StatusCreated, customer)
}

// UpdateCustomer godoc
// @Summary      Update a customer
// @Description  Updates an existing customer's information
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customer  body      dtos.UpdateCustomerDTO  true  "Customer update data"
// @Success      200
// @Failure      400       {string}  string  "Invalid request"
// @Failure      500       {string}  string  "Failed to update customer"
// @Router       /customers [patch]
func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	dto := dtos.UpdateCustomerDTO{}
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request: "+err.Error())
	}

	err := h.Service.UpdateCustomer(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update customer: "+err.Error())
	}

	return c.NoContent(http.StatusOK)
}
