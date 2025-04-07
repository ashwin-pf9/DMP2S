package main

import (
	authpb "authservice/proto"
	"authservice/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	/*-- Enable Reflection --*/
	reflection.Register(grpcServer)

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
