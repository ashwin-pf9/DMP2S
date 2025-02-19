package rest

import (
	"fmt"
	"net/http"

	"DMP2S/api/rest/handlers" // Import handlers
)

// StartServer initializes and starts the HTTP server
func StartServer() error {
	//Better than http.HandleFunc(), gives more control
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/register", handlers.RegisterHandler) //maps /register endpoint to path "handlers.RegisterHandler"
	mux.HandleFunc("/login", handlers.LoginHandler)       //maps /login endpoint to path "handlers.LoginHandler"

	// Start the server
	port := ":8080"
	fmt.Println("Server started on", port)
	return http.ListenAndServe(port, mux)
}
