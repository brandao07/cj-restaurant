package models

import "time"

// Database Model

type Customer struct {
	ID                int        `db:"id"`
	CreatedAt         time.Time  `db:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at"`
	Status            string     `db:"status"`
	RestaurantTableID int        `db:"restaurant_table_id"`
	OrderID           int        `db:"order_id"`
}
