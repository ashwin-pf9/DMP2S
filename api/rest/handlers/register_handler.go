package handlers

import (
	"encoding/json"
	"net/http"

	"DMP2S/internal/core/services" // Import the auth service
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler processes user sign-ups
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.RegisterUser(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
