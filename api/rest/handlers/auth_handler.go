package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	authpb "github.com/ashwin-pf9/DMP2S/internal/protobuffs/auth"
	"google.golang.org/grpc"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	RoleID   uint   `json:"role_id"`
}

// Define gRPC auth service client
var authClient authpb.AuthServiceClient

// Initialize gRPC client (Call this in main.go)
func InitAuthClient() {
	// Establish a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	authClient = authpb.NewAuthServiceClient(conn)
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

	// fmt.Println("entered name : ", creds.Email)
	// fmt.Println("entered role_id : ", creds.Password)
	// fmt.Println("entered name : ", creds.Name)
	// fmt.Println("entered role_id : ", creds.RoleID)

	// Create gRPC context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	InitAuthClient()

	user, err := authClient.Register(ctx, &authpb.RegisterRequest{
		Email:    creds.Email,
		Password: creds.Password,
		Name:     creds.Name,
		RoleId:   int32(creds.RoleID),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Sending response back to the client
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

	// Create gRPC context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	InitAuthClient()

	user, err := authClient.Login(ctx, &authpb.LoginRequest{
		Email:    creds.Email,
		Password: creds.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"token":     user.Token, // Ensure `user` has a valid Token field
		"user_name": user.UserName,
		"user_id":   user.UserId,
	})
}
