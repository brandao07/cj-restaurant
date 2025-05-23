package models

import "time"

// Database Model

type Menu struct {
	ID        int        `db:"id"`         
	CreatedAt time.Time  `db:"created_at"` 
	UpdatedAt time.Time  `db:"updated_at"` 
	DeletedAt *time.Time `db:"deleted_at"` 
	StarterID int        `db:"starter_id"`
    MainID    int        `db:"main_id"`
    DessertID int        `db:"dessert_id"`
    DrinkID   int        `db:"drink_id"`
	Price     float32    `db:"price"`
}
