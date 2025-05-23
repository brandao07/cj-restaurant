package main

import (
	"fmt"
	"os"

	"github.com/brandao07/cj-restaurant/menu-service/internal/data"
	"github.com/brandao07/cj-restaurant/menu-service/internal/handlers"
	"github.com/brandao07/cj-restaurant/menu-service/internal/services"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	DSN  string
	Port string
}

func loadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DB_DSN is not set in environment")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return nil, fmt.Errorf("SERVER_PORT is not set in environment")
	}

	return &Config{
		DSN:  dsn,
		Port: port,
	}, nil
}

func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
		return
	}

	// Initialize the database connection
	db, err := data.NewDB(config.DSN + "?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return
	}

	// Close the database connection
	defer func() {
		if err := db.Close(); err != nil {
			log.Error("Error closing the database connection:", err)
		}
	}()

	// Initialize Repositories, Services, and Handlers
	itemRepository := db
	menuRepository := db
	menuService := services.NewMenuService(menuRepository, itemRepository)
	menuHandler := handlers.NewMenuHandler(menuService)

	appHandlers := Handlers{
		MenuHandler: menuHandler,
	}

	// Set up the server
    e := NewRouter(&appHandlers)
    addr := ":" + config.Port
    log.Info("Starting server on ", addr)
    e.Logger.Fatal(e.Start(addr))
}
