package data

import (
	"github.com/brandao07/cj-restaurant/order-service/internal/data/models"
	"github.com/brandao07/cj-restaurant/order-service/internal/dtos"
)

func (db *DB) CreateRestaurantOrder(dto *dtos.CreateRestaurantOrderDTO) (*dtos.RestaurantOrder, error) {
	query := `
        INSERT INTO restaurant_orders (menu_ids, total)
        VALUES ($1, $2)
        RETURNING id, status
    `
	var id int
	var status string
	err := db.QueryRow(query, dto.MenuIDs, dto.Total).Scan(&id, &status)
	if err != nil {
		return nil, err
	}
	return dtos.NewRestaurantOrder(id, dto.MenuIDs, dto.Total, status), nil
}

func (db *DB) UpdateRestaurantOrder(dto dtos.UpdateRestaurantOrderDTO) error {
	query := `
		UPDATE restaurant_orders
		SET status = $1
		WHERE id = $2
	`
	_, err := db.Exec(query, dto.Status, dto.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetRestaurantOrderByID(id int) (*dtos.RestaurantOrder, error) {
	query := `
		SELECT id, menu_ids, total, status
		FROM restaurant_orders
		WHERE id = $1 AND deleted_at IS NULL
	`
	var order models.RestaurantOrder
	err := db.QueryRow(query, id).Scan(&order.ID, &order.MenuIDs, &order.Total, &order.Status)
	if err != nil {
		return nil, err
	}
	return dtos.NewRestaurantOrder(order.ID, order.MenuIDs, order.Total, order.Status), nil
}
