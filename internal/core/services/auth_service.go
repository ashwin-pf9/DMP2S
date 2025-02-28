package services

import (
	"DMP2S/internal/infrastructure/supabaseclient"
	"context"
	"errors"

	"github.com/nedpals/supabase-go"
)

func RegisterUser(email string, password string, name string, roleID uint) (*supabase.User, error) {
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
	// Use map instead of struct for Data
	// user, err := supabaseClient.Auth.SignUp(context.Background(), supabase.UserCredentials{
	// 	Email:    email,
	// 	Password: password,
	// })

	if err != nil {
		return nil, errors.New("registration failed...: " + err.Error())
	}

	return user, nil
}

// LoginUser handles user login
func LoginUser(email, password string) (*supabase.AuthenticatedDetails, error) {
	supabaseClient := supabaseclient.InitSupabaseClient()
	user, err := supabaseClient.Auth.SignIn(context.Background(), supabase.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, errors.New("login failed: " + err.Error())
	}
	return user, nil
}
