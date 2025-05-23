package models

import "time"

// Database Model

type Item struct {
	ID        int        `db:"id"`         
	CreatedAt time.Time  `db:"created_at"` 
	UpdatedAt time.Time  `db:"updated_at"` 
	DeletedAt *time.Time `db:"deleted_at"` 
	Name      string     `db:"name"`
	Type      string     `db:"type"`
}
