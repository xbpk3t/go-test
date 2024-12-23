/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/samber/mo"

	"github.com/spf13/cobra"
)

// moCmd represents the mo command
var moCmd = &cobra.Command{
	Use:   "mo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mo called")

		f := mo.Some(11).FlatMap(func(value int) mo.Option[int] {
			return mo.Some(value * 2)
		}).OrElse(1234)
		fmt.Println(f)

		r := mo.Some(42).Map(func(i int) (int, bool) {
			return 1234, true
		})
		fmt.Println(r.IsPresent(), r.OrEmpty())

		fmt.Println(lo.FindUniques([]int{1, 2, 3, 4, 1}))

		// mo.Some([]int{1, 2, 3, 4}).FlatMap(func(val []int) mo.Option[[]int] {
		// })
	},
}

func init() {
	rootCmd.AddCommand(moCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
