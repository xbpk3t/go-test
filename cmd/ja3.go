/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ja3Cmd represents the ja3 command
var ja3Cmd = &cobra.Command{
	Use:   "ja3",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ja3 called")
	},
}

func init() {
	rootCmd.AddCommand(ja3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ja3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ja3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
