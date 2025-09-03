// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"
	"test/pkg/validators"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

type UserInfo struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130" label:"年龄"`
}

type MyStruct struct {
	String string `validate:"is-awesome"`
}

// validatorCmd represents the validator command
var validatorCmd = &cobra.Command{
	Use:   "validator",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validator called")
		validators.V()

		var validate *validator.Validate

		validate = validator.New()
		validate.RegisterValidation("is-awesome", ValidateMyVal)

		s := MyStruct{String: "not awesome"}
		err := validate.Struct(s)
		if err != nil {
			fmt.Printf("%v", err)
		}
	},
}

// 初始化Validate/v10国际化
func init() {
	rootCmd.AddCommand(validatorCmd)
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
