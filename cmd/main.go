package main

import (
	"DMP2S/cmd/service"
	"log"

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

	var createPipelineCmd = &cobra.Command{
		Use:   "createpipeline",
		Short: "Create Single Pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			user_id, _ := cmd.Flags().GetString("user_id")
			name, _ := cmd.Flags().GetString("name")

			service.CreatePipeline(user_id, name)
		},
	}
	createPipelineCmd.Flags().String("user_id", "", "User ID")
	createPipelineCmd.Flags().String("name", "", "Pipeline Name")

	var getUserPipelines = &cobra.Command{
		Use:   "getPipelines",
		Short: "Fetch User Pipelines",
		Run: func(cmd *cobra.Command, args []string) {
			user_id, _ := cmd.Flags().GetString("user_id")

			service.GetUserPipelines(user_id)
		},
	}
	getUserPipelines.Flags().String("user_id", "", "User ID")

	var getPipelineStages = &cobra.Command{
		Use:   "getStages",
		Short: "Fetch Pipeline Stages",
		Run: func(cmd *cobra.Command, args []string) {
			pipeline_id, _ := cmd.Flags().GetString("pipeline_id")

			service.GetPipelineStages(pipeline_id)
		},
	}
	getPipelineStages.Flags().String("pipeline_id", "", "Pipeline ID")

	var executePipeline = &cobra.Command{
		Use:   "start",
		Short: "Start Execution of Pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			pipeline_id, _ := cmd.Flags().GetString("pipeline_id")

			service.ExecutePipeline(pipeline_id)
		},
	}

	executePipeline.Flags().String("pipeline_id", "", "Pipeline ID")

	var deletePipeline = &cobra.Command{
		Use:   "delete",
		Short: "Delete Pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			pipeline_id, _ := cmd.Flags().GetString("pipeline_id")

			service.DeletePipeline(pipeline_id)
		},
	}
	deletePipeline.Flags().String("pipeline_id", "", "Pipeline ID")

	// Add commands to root
	rootCmd.AddCommand(registerCmd, loginCmd, createPipelineCmd, getUserPipelines, getPipelineStages, executePipeline, deletePipeline)

	// Initialize gRPC auth client
	service.InitAuthClient("localhost:30080")

	// Initialize gRPC crud client
	service.InitCrudClient("localhost:30080")

	// Initialize gRPC orch client
	service.InitOrchClient("localhost:30080")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
