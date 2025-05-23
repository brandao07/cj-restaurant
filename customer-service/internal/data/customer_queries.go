package data

import (
	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
)

func (db *DB) CreateCustomer(dto dtos.CreateCustomerDTO) (*dtos.Customer, error) {
	query := `
		INSERT INTO customers (restaurant_table_id, order_id)
		VALUES ($1, $2)
		RETURNING id, status
	`
	var id int
	var status string
	err := db.QueryRow(query, dto.RestaurantTableID, dto.OrderID).Scan(&id, &status)
	if err != nil {
		return nil, err
	}
	return dtos.NewCustomer(id, status, dto.RestaurantTableID, dto.OrderID), nil
}

func (db *DB) UpdateCustomer(dto dtos.UpdateCustomerDTO) error {
	query := `
		UPDATE customers
		SET status = $1
		WHERE id = $2
	`
	_, err := db.Exec(query, dto.Status, dto.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetCustomerByRestaurantTableID(restaurantTableID int) (*dtos.Customer, error) {
	query := `
		SELECT id, status, restaurant_table_id, order_id
		FROM customers
		WHERE restaurant_table_id = $1 AND deleted_at IS NULL AND (status = 'EATING' OR status = 'WAITING')
	`
	var customer dtos.Customer
	err := db.QueryRow(query, restaurantTableID).
		Scan(&customer.ID, &customer.Status, &customer.RestaurantTableID, &customer.OrderID)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (db *DB) GetCustomerByID(id int) (*dtos.Customer, error) {
	query := `
		SELECT id, status, restaurant_table_id, order_id
		FROM customers
		WHERE id = $1 AND deleted_at IS NULL
	`
	var customer dtos.Customer
	err := db.QueryRow(query, id).
		Scan(&customer.ID, &customer.Status, &customer.RestaurantTableID, &customer.OrderID)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
