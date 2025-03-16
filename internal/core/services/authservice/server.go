package main

import (
	"DMP2S/internal/core/services/authservice/service"
	"DMP2S/internal/protobuffs/authpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create a new instance of AuthService
	authService := service.NewAuthService()

	// Create a new instance of AuthServer and pass authService to it
	authServer := service.NewAuthServer(authService)

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the AuthServer with the gRPC server
	authpb.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
