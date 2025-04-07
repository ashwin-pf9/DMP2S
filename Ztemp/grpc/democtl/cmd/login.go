package cmd

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	clipb "github.com/ashwin-pf9/DMP2S/cmd/democtl/proto"
	"github.com/spf13/cobra"
)

// CLI arguments
var email, password string

// loginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the system",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to gRPC Gateway Server
		conn, err := grpc.Dial("http://localhost:50055", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		client := clipb.NewAuthServiceClient(conn)

		// Create login request
		req := &clipb.LoginRequest{
			Email:    email,
			Password: password,
		}

		// Call gRPC method
		resp, err := client.Login(context.Background(), req)
		if err != nil {
			log.Fatalf("Login failed: %v", err)
		}

		fmt.Println("Login successful!")
		fmt.Printf("User ID: %s\n", resp.UserId)
		fmt.Printf("User Name: %s\n", resp.UserName)
		fmt.Printf("Token: %s\n", resp.Token)
	},
}

func init() {
	rootCmd.AddCommand(LoginCmd)
	// Define CLI flags
	LoginCmd.Flags().StringVarP(&email, "email", "e", "", "User email")
	LoginCmd.Flags().StringVarP(&password, "password", "p", "", "User password")
	LoginCmd.MarkFlagRequired("email")
	LoginCmd.MarkFlagRequired("password")
}
