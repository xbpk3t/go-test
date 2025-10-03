/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"test/services/api"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Long: `Start the API server with filtering, sorting, and field selection capabilities.
The server will listen on localhost:8080 by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the database
		if err := api.InitDB(); err != nil {
			fmt.Printf("Failed to initialize database: %v\n", err)
			return
		}

		// Register example endpoints
		http.HandleFunc("/api/examples", api.Handle(api.ExampleHandler))
		http.HandleFunc("/api/examples/", api.Handle(api.GetSampleHandler))
		http.HandleFunc("/api/examples/create", api.Handle(api.CreateSampleHandler))

		// Start the server
		fmt.Println("API server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Server failed to start: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
