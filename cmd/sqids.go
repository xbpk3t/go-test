/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sqids/sqids-go"
	"time"
)

// sqidsCmd represents the sqids command
var sqidsCmd = &cobra.Command{
	Use:   "sqids",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// s, _ := sqids.New(sqids.Options{
		// 	MinLength: 10,
		// })
		// id, _ := s.Encode([]uint64{1, 2, 3}) // "86Rf07xd4z"
		// numbers := s.Decode(id)              // [1, 2, 3]
		// fmt.Println(id, numbers)

		s, _ := sqids.New()
		id, _ := s.Encode([]uint64{1234567890}) // "PcHfYmv"
		fmt.Println(id)

		start := time.Now().Unix()
		end := time.Now().Add(24 * time.Hour).Unix()
		id, _ = s.Encode([]uint64{uint64(start), uint64(end)}) // "s6eUn008oGU27p"
		fmt.Println(id)

		numbers := s.Decode(id) // [1714879533 1714965933]
		fmt.Println(numbers)
	},
}

func init() {
	rootCmd.AddCommand(sqidsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqidsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqidsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
