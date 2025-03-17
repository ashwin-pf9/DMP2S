package service

import (
	"authservice/supabaseclient"
	"context"
	"errors"

	"github.com/nedpals/supabase-go"
)

type AuthenticatedUser struct {
	UserID   string `json:"user_id"`
	UserName string `json:"name"`
	Token    string `json:"token"`
}

type AuthService struct{}

// NewAuthService returns a new instance of AuthService.
func NewAuthService() *AuthService {
	return &AuthService{}
}

// func (s *AuthService) Login(email, password string) (*AuthenticatedUser, error)
func (s *AuthService) Register(email string, password string, name string, roleID uint) (*supabase.User, error) {
	supabaseClient := supabaseclient.InitSupabaseClient()

	ctx := context.Background()

	// User registration credentials
	signUpData := supabase.UserCredentials{
		Email:    email,
		Password: password,
		Data: map[string]interface{}{
			"name":    name,
			"role_id": roleID,
		},
	}
	user, err := supabaseClient.Auth.SignUp(ctx, signUpData)

	if err != nil {
		return nil, errors.New("registration failed...: " + err.Error())
	}

	return user, nil
}

// Login authenticates a user through Supabase.
func (s *AuthService) Login(email, password string) (*AuthenticatedUser, error) {
	supabaseClient := supabaseclient.InitSupabaseClient()

	user, err := supabaseClient.Auth.SignIn(context.Background(), supabase.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, errors.New("login failed: " + err.Error())
	}

	var profile []struct {
		UserName string `json:"name"`
	}

	err = supabaseClient.DB.
		From("profiles").
		Select("name").
		Eq("id", user.User.ID).
		Execute(&profile)

	if err != nil {
		return nil, errors.New("failed to fetch user profile: " + err.Error())
	}

	return &AuthenticatedUser{
		UserID:   user.User.ID,
		UserName: profile[0].UserName,
		Token:    user.AccessToken,
	}, nil
}
