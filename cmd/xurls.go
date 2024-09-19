/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mvdan.cc/xurls/v2"

	"github.com/spf13/cobra"
)

// xurlsCmd represents the xurls command
var xurlsCmd = &cobra.Command{
	Use:   "xurls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("xurls called")

		rxRelaxed := xurls.Relaxed()
		rxRelaxed.FindString("Do gophers live in golang.org?")  // "golang.org"
		rxRelaxed.FindString("This string does not have a URL") // ""

		rxStrict := xurls.Strict()
		rxStrict.FindAllString("must have scheme: http://foo.com/.", -1) // []string{"http://foo.com/"}
		rxStrict.FindAllString("no scheme, no match: foo.com", -1)       // []string{}

	},
}

func init() {
	rootCmd.AddCommand(xurlsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// xurlsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// xurlsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
