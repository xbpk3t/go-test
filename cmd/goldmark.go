/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// goldmarkCmd represents the goldmark command
var goldmarkCmd = &cobra.Command{
	Use:   "goldmark",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goldmark called")

		// num := 10
		// qj, _ := gjson.LoadJson(questions)
		// rands := garray.NewArrayFrom(qj.Array()).Shuffle().SubSlice(0, num)
		// rt := ""
		// for _, rand := range rands {
		// 	rt += gconv.String(rand) + "\n"
		// }
		// fmt.Println(rt)

		// lo.Shuffle()
	},
}

func init() {
	rootCmd.AddCommand(goldmarkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goldmarkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goldmarkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
