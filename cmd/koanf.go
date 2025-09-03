/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// koanfCmd represents the koanf command
var koanfCmd = &cobra.Command{
	Use:   "koanf",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("koanf called")
	},
}

func init() {
	rootCmd.AddCommand(koanfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// koanfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// koanfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
