package main

import (
	"fmt"
	"log"

	"DMP2S/api/rest"
)

func main() {
	fmt.Println("Starting REST API server...")

	// Initialize database and run migrations
	fmt.Println("Running database migrations...")
	// db.InitDatabase() // Running auto migrations - .env LOADED

	// Call the server setup function
	err := rest.StartRESTServer()
	if err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
