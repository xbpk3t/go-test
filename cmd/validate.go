/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
	"test/pkg/validates"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {
		u := &validates.UserForm{
			Name: "inhere",
		}
		zhcn.RegisterGlobal()

		// 创建 Validation 实例
		// v := validate.Struct(u)
		// 或者使用
		v := validate.New(u)
		zhcn.Register(v)

		if v.Validate() { // 验证成功
			// do something ...
		} else {
			fmt.Println(v.Errors)               // 所有的错误消息
			fmt.Println(v.Errors.One())         // 返回随机一条错误消息
			fmt.Println(v.Errors.Field("Name")) // 返回该字段的错误消息
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
