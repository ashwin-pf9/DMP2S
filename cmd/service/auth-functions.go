package service

import (
	"context"
	"fmt"
	"log"

	"DMP2S/cmd/utils"
	grpcauthpb "DMP2S/global-protos/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Global variable to hold the gRPC client
var authClient grpcauthpb.AuthServiceClient

// initGRPCClient initializes the connection to the gRPC server
func InitAuthClient(url string) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	authClient = grpcauthpb.NewAuthServiceClient(conn)
}

// registerUser makes a gRPC request to the Register service
func RegisterUser(email, password, name string, roleID int) {
	req := &grpcauthpb.RegisterRequest{
		Email:    email,
		Password: password,
		Name:     name,
		RoleId:   int32(roleID),
	}

	resp, err := authClient.Register(context.Background(), req)
	if err != nil {
		log.Fatalf("Error registering user: %v", err)
	}
	fmt.Printf("User registered! UserID: %s, Email: %s\n", resp.UserId, resp.Email)
}

// loginUser makes a gRPC request to the Login service
func LoginUser(email, password string) {
	req := &grpcauthpb.LoginRequest{
		Email:    email,
		Password: password,
	}

	resp, err := authClient.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Error logging in: %v", err)
	}

	fmt.Printf("Login successful! UserID: %s, UserName: %s", resp.UserId, resp.UserName)
	utils.SaveToken(resp.Token)
}
