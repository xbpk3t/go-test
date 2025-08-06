/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// haxmapCmd represents the haxmap command
var haxmapCmd = &cobra.Command{
	Use:   "haxmap",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("haxmap called")
	},
}

func init() {
	rootCmd.AddCommand(haxmapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// haxmapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// haxmapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
