package main

import (
	"DMP2S/internal/infrastructure/db"
	"fmt"
	"log"

	"DMP2S/api/rest"
)

func main() {
	fmt.Println("Starting API server...")

	// Initialize database and run migrations
	fmt.Println("Running database migrations...")
	db.InitDatabase() // Running auto migrations

	// Call the server setup function
	err := rest.StartServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
