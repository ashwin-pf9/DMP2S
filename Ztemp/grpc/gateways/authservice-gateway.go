package gateways

import (
	"context"
	clipb "grpc/proto"
)

// AuthServiceGateway will act as a proxy
type AuthServiceGateway struct {
	clipb.UnimplementedAuthServiceServer
	AuthClient clipb.AuthServiceClient // gRPC client for actual Auth microservice
}

// Implement Login method to forward request to Auth microservice
func (a *AuthServiceGateway) Login(ctx context.Context, req *clipb.LoginRequest) (*clipb.LoginResponse, error) {
	// Forward request to actual AuthService microservice
	return a.AuthClient.Login(ctx, req)
}

// Implement Register method to forward request
func (a *AuthServiceGateway) Register(ctx context.Context, req *clipb.RegisterRequest) (*clipb.RegisterResponse, error) {
	return a.AuthClient.Register(ctx, req)
}
