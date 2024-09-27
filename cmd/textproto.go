/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/textproto"

	"github.com/spf13/cobra"
)

// textprotoCmd represents the textproto command
var textprotoCmd = &cobra.Command{
	Use:   "textproto",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("textproto called")
	},
}

type Header map[string][]string

// map作为构造函数
func NewHeader(from, subject string, to ...string) Header {
	headers := Header{}
	headers.SetFrom(from)
	if len(to) > 0 {
		headers.SetTo(to...)
	}
	return headers
}

func (h Header) Add(key, value string) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	h[key] = append(h[key], value)
}

// 不返回参数，只修改属性
func (h Header) Set(key, value string) {
	h[textproto.CanonicalMIMEHeaderKey(key)] = []string{value}
}

func (h Header) SetFrom(mail string) {
	h.Set("From", mail)
}

func (h Header) SetTo(receivers ...string) {
	for _, receiver := range receivers {
		h.Set("Receiver", receiver)
	}
}

func init() {
	rootCmd.AddCommand(textprotoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textprotoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textprotoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
