package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"DMP2S/internal/core/services" // Import the auth service
)

type Credentials struct {
	Email    string           `json:"email"`
	Password string           `json:"password"`
	Data     services.NewUser `json:"data"`
}

// RegisterHandler processes user sign-ups
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 	return
	// }

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Debug: Print the JSON payload being sent
	payload, _ := json.Marshal(creds.Data)
	fmt.Println("Payload sent to Supabase:", string(payload))

	user, err := services.RegisterUser(creds.Email, creds.Password, creds.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Sending response back to the client
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// LoginHandler processes user logins
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.LoginUser(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
