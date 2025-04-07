/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
// main.go
package main

import (
	"log"
	"os"

	"github.com/ashwin-pf9/DMP2S/cmd/democtl/cmd"
	"github.com/spf13/cobra"
)

func main() {
	// Create root command and add subcommands
	rootCmd := &cobra.Command{Use: "democtl"}

	// Add the subcommands
	rootCmd.AddCommand(cmd.LoginCmd)
	rootCmd.AddCommand(cmd.RegisterCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
