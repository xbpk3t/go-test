package cmd

import (
	"fmt"
	"github.com/google/go-cmp/cmp"

	"github.com/spf13/cobra"
)

// cmpCmd represents the cmp command
var cmpCmd = &cobra.Command{
	Use:   "cmp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmp called")

		var m map[string]int
		var n map[string]int

		// isEqual := cmp.Compare(m, n)
		isEqual := cmp.Equal(m, n)
		fmt.Print(isEqual)
	},
}

func init() {
	rootCmd.AddCommand(cmpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
