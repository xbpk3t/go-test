/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go.deanishe.net/fuzzy"

	"github.com/spf13/cobra"
)

// gofuzzyCmd represents the gofuzzy command
var gofuzzyCmd = &cobra.Command{
	Use:   "gofuzzy",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		actors := []string{"Tommy Lee Jones", "James Earl Jones", "Keanu Reeves"}
		fuzzy.SortStrings(actors, "jej")
		fmt.Println(actors[0])
		// -> James Earl Jones
	},
}

func init() {
	rootCmd.AddCommand(gofuzzyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gofuzzyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gofuzzyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
