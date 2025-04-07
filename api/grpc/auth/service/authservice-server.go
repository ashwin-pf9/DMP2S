package service

import (
	"context"
	grpcauthpb "grpc/auth/proto"
)

type AuthServiceServer struct {
	grpcauthpb.UnimplementedAuthServiceServer
	AuthClient grpcauthpb.AuthServiceClient
}

func (s *AuthServiceServer) Login(ctx context.Context, req *grpcauthpb.LoginRequest) (*grpcauthpb.LoginResponse, error) {
	return s.AuthClient.Login(ctx, req)
}

func (s *AuthServiceServer) Register(ctx context.Context, req *grpcauthpb.RegisterRequest) (*grpcauthpb.RegisterResponse, error) {
	return s.AuthClient.Register(ctx, req)
}
