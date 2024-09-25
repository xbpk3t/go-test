// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"test/pkg/validators"
)

type UserInfo struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130" label:"年龄"`
}

// validatorCmd represents the validator command
var validatorCmd = &cobra.Command{
	Use:   "validator",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validator called")
		validators.V()
	},
}

// 初始化Validate/v10国际化
func init() {
	rootCmd.AddCommand(validatorCmd)
}
