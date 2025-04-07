package main

import (
	"log"

	"github.com/ashwin-pf9/DMP2S/cmd/service"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "democtl",
		Short: "A CLI tool for registering and logging in with the auth service",
	}

	// Register command
	var registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			name, _ := cmd.Flags().GetString("name")
			roleID, _ := cmd.Flags().GetInt("role_id")

			// Make Register request to gRPC server
			service.RegisterUser(email, password, name, roleID)
		},
	}
	registerCmd.Flags().String("email", "", "Email address")
	registerCmd.Flags().String("password", "", "Password")
	registerCmd.Flags().String("name", "", "Name of the user")
	registerCmd.Flags().Int("role_id", 0, "Role ID")

	// Login command
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login with existing user",
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")

			// Make Login request to gRPC server
			service.LoginUser(email, password)
		},
	}
	loginCmd.Flags().String("email", "", "Email address")
	loginCmd.Flags().String("password", "", "Password")

	// Add commands to root
	rootCmd.AddCommand(registerCmd, loginCmd)

	// Initialize gRPC client
	service.InitGRPCClient()

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
