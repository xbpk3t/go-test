/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/bytedance/sonic"
	"log"

	"github.com/spf13/cobra"
)

// sonicCmd represents the sonic command
var sonicCmd = &cobra.Command{
	Use:   "sonic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sonic called")

		m := map[string]any{
			"name": "z3",
			"age":  20,
		}

		// sonic序列化
		byt, err := sonic.Marshal(&m)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("json: %+v\n", string(byt))

		// sonic反序列化
		um := make(map[string]any)
		err = sonic.Unmarshal(byt, &um)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("unjson: %+v\n", um)
	},
}

func init() {
	rootCmd.AddCommand(sonicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sonicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sonicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
