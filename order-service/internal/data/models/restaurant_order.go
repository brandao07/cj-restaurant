package models

import "time"

type RestaurantOrder struct {
	ID        int        `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	MenuIDs   []int      `db:"menu_ids"`
	Total     float32    `db:"total"`
	Status    string     `db:"status"`
}
