package data

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func NewDB(dsn string) (*DB, error) {
	var db *sqlx.DB
	var err error

	maxAttempts := 10
	delay := 3 * time.Second

	for i := 1; i <= maxAttempts; i++ {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			break
		}

		fmt.Printf("Attempt %d: Could not connect to DB: %v\n", i, err)
		time.Sleep(delay)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB after %d attempts: %w", maxAttempts, err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	return &DB{db}, nil
}
