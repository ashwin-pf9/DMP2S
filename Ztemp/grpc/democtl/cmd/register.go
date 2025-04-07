/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	clipb "github.com/ashwin-pf9/DMP2S/cmd/democtl/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var name, regEmail, regPassword string
var role_id int32

// RegisterCmd represents the register command
var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to gRPC Gateway Server
		conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		client := clipb.NewAuthServiceClient(conn)

		// Create registration request
		req := &clipb.RegisterRequest{
			Name:     name,
			Email:    regEmail,
			Password: regPassword,
			RoleId:   role_id,
		}

		// Call gRPC method
		resp, err := client.Register(context.Background(), req)
		if err != nil {
			log.Fatalf("Registration failed: %v", err)
		}

		fmt.Println("Registration successful!")
		fmt.Printf("User ID: %s\n", resp.UserId)
		fmt.Printf("Email: %s\n", resp.Email)
	},
}

func init() {
	rootCmd.AddCommand(RegisterCmd)
	RegisterCmd.Flags().StringVarP(&name, "name", "n", "John", "Full name")
	RegisterCmd.Flags().StringVarP(&regEmail, "email", "e", "default@domain.com", "Email")
	RegisterCmd.Flags().StringVarP(&regPassword, "password", "p", "", "Password")
	RegisterCmd.Flags().Int32VarP(&role_id, "role_id", "r", 0, "Role Id")
	RegisterCmd.MarkFlagRequired("name")
	RegisterCmd.MarkFlagRequired("email")
	RegisterCmd.MarkFlagRequired("password")
	RegisterCmd.MarkFlagRequired("role_id")
}
