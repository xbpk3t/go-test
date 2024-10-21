/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/resend/resend-go/v2"
	"github.com/spf13/cobra"
)

// mailCmd represents the mail command
var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		client := resend.NewClient("")

		params := &resend.SendEmailRequest{
			From:    "Acme <onboarding@resend.dev>",
			To:      []string{"jeffcottlu@gmail.com"},
			Subject: "hello world",
			Html:    "<p>it works!</p>",
		}

		sent, err := client.Emails.SendWithContext(ctx, params)

		if err != nil {
			panic(err)
		}
		fmt.Println(sent.Id)
	},
}

func init() {
	rootCmd.AddCommand(mailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
