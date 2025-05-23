package data

import (
	"github.com/brandao07/cj-restaurant/customer-service/internal/common"
	"github.com/brandao07/cj-restaurant/customer-service/internal/data/models"
	"github.com/brandao07/cj-restaurant/customer-service/internal/dtos"
)

func (db *DB) GetPendingPaymentTables() ([]dtos.RestaurantTable, error) {
	query := `
		SELECT DISTINCT restaurant_table_id
		FROM customers
		WHERE status IN ('WAITING', 'EATING', 'NEEDS_PAYMENT') AND deleted_at IS NULL
		ORDER BY order_id ASC
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []dtos.RestaurantTable
	for rows.Next() {
		var table models.RestaurantTable
		if err := rows.Scan(&table.ID); err != nil {
			return nil, err
		}
		tables = append(tables, *dtos.NewRestaurantTable(table.ID, common.TableStatusUnavailable))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}

func (db *DB) GetPendingDeliveryTables() ([]dtos.RestaurantTable, error) {
	query := `
		SELECT DISTINCT restaurant_table_id
		FROM customers
		WHERE status IS 'WAITING' AND deleted_at IS NULL
		ORDER BY order_id ASC
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []dtos.RestaurantTable
	for rows.Next() {
		var table models.RestaurantTable
		if err := rows.Scan(&table.ID); err != nil {
			return nil, err
		}
		tables = append(tables, *dtos.NewRestaurantTable(table.ID, common.TableStatusUnavailable))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tables, nil
}
