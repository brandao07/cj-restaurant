package dtos

// CreateCustomerDTO represents the payload to create a new customer
type CreateCustomerDTO struct {
	MenuIDs           []int `json:"menu_ids"`
	RestaurantTableID int   `json:"restaurant_table_id"`
	OrderID           int   `json:"order_id"`
}

func NewCreateCustomerDTO(menuIDs []int, restaurantTableID, orderID int) *CreateCustomerDTO {
	return &CreateCustomerDTO{
		MenuIDs:           menuIDs,
		RestaurantTableID: restaurantTableID,
		OrderID:           orderID,
	}
}

// Customer represents a customer response
type Customer struct {
	ID                int    `json:"id"`
	Status            string `json:"status"`
	RestaurantTableID int    `json:"restaurant_table_id"`
	OrderID           int    `json:"order_id"`
}

func NewCustomer(id int, status string, restaurantTableID, orderID int) *Customer {
	return &Customer{
		ID:                id,
		Status:            status,
		RestaurantTableID: restaurantTableID,
		OrderID:           orderID,
	}
}

// UpdateCustomerDTO represents the payload to update a customer
type UpdateCustomerDTO struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func NewUpdateCustomerDTO(id int, status string) *UpdateCustomerDTO {
	return &UpdateCustomerDTO{
		ID:     id,
		Status: status,
	}
}
