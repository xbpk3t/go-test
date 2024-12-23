/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cobraCmd represents the cobra command
var cobraCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A brief description of your command",
	// Args:  cobra.ExactArgs(1),
	// Args:  cobra.NoArgs,
	// Args:  cobra.RangeArgs(0, 2),
	// Args:  cobra.MinimumNArgs(1),
	// Args:  cobra.MaximumNArgs(1),
	// Args: cobra.OnlyValidArgs,
	// Args: cobra.ArbitraryArgs,

	ValidArgs: []string{"x", "y"},
	Args:      cobra.OnlyValidArgs,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobra called")
	},
}

func init() {
	rootCmd.AddCommand(cobraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cobraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cobraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
