package service

import (
	"DMP2S/internal/protobuffs/authpb"
	"context"
)

// AuthServer struct
type AuthServer struct {
	authpb.UnimplementedAuthServiceServer              // Embeds the gRPC server interface
	authService                           *AuthService // Auth service instance
}

// NewAuthServer initializes AuthServer
func NewAuthServer(authService *AuthService) *AuthServer {
	return &AuthServer{authService: authService}
}

// Implements the Login gRPC method
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user, err := s.authService.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &authpb.LoginResponse{
		UserId:   user.UserID,   // Returning user ID
		UserName: user.UserName, //// Returning user Name
		Token:    user.Token,    // Returning access token
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	user, err := s.authService.Register(req.Email, req.Password, req.Name, uint(req.RoleId))
	if err != nil {
		return nil, err
	}
	return &authpb.RegisterResponse{
		UserId: user.ID,
		Email:  user.Email,
	}, nil
}
