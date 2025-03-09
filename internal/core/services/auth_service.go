package services

import (
	"DMP2S/internal/infrastructure/supabaseclient"
	"context"
	"errors"
	"log"

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

type AuthenticatedUser struct {
	UserID   string `json:"user_id"`
	UserName string `json:"name"`
	Token    string `json:"token"`
}

// LoginUser handles user login
func LoginUser(email, password string) (*AuthenticatedUser, error) {
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

	log.Println("in login user function")
	if err != nil {
		return nil, errors.New("failed to fetch user profile: " + err.Error())
	}

	return &AuthenticatedUser{
		UserID:   user.User.ID,
		UserName: profile[0].UserName,
		Token:    user.AccessToken,
	}, nil
}
