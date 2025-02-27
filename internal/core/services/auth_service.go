package services

import (
	"DMP2S/internal/infrastructure/supabaseclient"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nedpals/supabase-go"
)

type NewUser struct {
	Name   string `json:"name"`
	RoleID uint   `json:"role_id"`
}

func RegisterUser(email string, password string, data NewUser) (*supabase.User, error) {
	supabaseClient := supabaseclient.InitSupabaseClient()

	// Debug: Print the JSON payload being sent
	payload, _ := json.Marshal(data)
	fmt.Println("Payload sent to Supabase:", string(payload))

	// Use map instead of struct for Data
	user, err := supabaseClient.Auth.SignUp(context.Background(), supabase.UserCredentials{
		Email:    email,
		Password: password,
		Data:     data,
	})

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
