/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"test/pkg/validates"

	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {
		u := &validates.UserForm{
			Name:   "inhere123",
			Prefix: "123",
			// Age: 20,
			// UpdateAt: time.Now(),
		}
		zhcn.RegisterGlobal()

		// 创建 Validation 实例
		// v := validate.Struct(u)
		// 或者使用
		v := validate.New(u)
		// 启用 ContinueOnError 选项，验证所有字段而不提前终止
		// 注意这个，如果不开启的话，就会在第一个字段验证不通过之后停止，不去验证后面的字段
		v.StopOnError = false
		zhcn.Register(v)

		if v.Validate() { // 验证成功
			// do something ...
			fmt.Println("验证成功")
		} else {
			fmt.Println("所有错误:", v.Errors) // 所有的错误消息
			// 注意：v.Errors.One() 在 ContinueOnError 模式下可能返回空字符串
			// 因为它会返回第一个错误，但 ContinueOnError 会继续收集所有字段的错误
			fmt.Println("随机一条错误:", v.Errors.One())               // 返回随机一条错误消息（第一个错误）
			fmt.Println("Name字段错误:", v.Errors.Field("Name"))     // 返回该字段的错误消息
			fmt.Println("Prefix字段错误:", v.Errors.Field("Prefix")) // 返回该字段的错误消息

			// 让我们检查所有字段的错误
			fmt.Println("\n逐个检查字段错误:")
			fields := []string{"Name", "Email", "Age", "UpdateAt", "Code", "Prefix"}
			for _, field := range fields {
				if err := v.Errors.Field(field); err != nil {
					fmt.Printf("%s 字段错误: %s\n", field, err)
				} else {
					fmt.Printf("%s 字段没有错误\n", field)
				}
			}

			// 显示所有字段的错误，使用 Errors.All() 方法
			fmt.Println("\n使用 Errors.All() 显示所有字段错误:")
			for field, errMsg := range v.Errors.All() {
				fmt.Printf("%s 字段错误: %v\n", field, errMsg)
			}
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
