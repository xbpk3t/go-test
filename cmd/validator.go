// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

//
// import (
// 	"errors"
// 	"fmt"
// 	CN_ZH "github.com/go-playground/locales/zh"
// 	ut "github.com/go-playground/universal-translator"
// 	"github.com/go-playground/validator/v10"
// 	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
// 	"github.com/spf13/cobra"
// 	"reflect"
// )
//
// type UserInfo struct {
// 	FirstName string `validate:"required"`
// 	LastName  string `validate:"required"`
// 	Age       uint8  `validate:"gte=0,lte=130" label:"年龄"`
// }
//
// var validate *validator.Validate
//
// // Validate/v10 全局验证器
// var trans ut.Translator
//
// // 检验并返回检验错误信息
// func Translate(err error) (errMsg string) {
// 	var errs validator.ValidationErrors
// 	errors.As(err, &errs)
// 	for _, err := range errs {
// 		errMsg = err.Translate(trans)
// 	}
// 	return
// }
//
// func validateStruct() {
// 	user := &UserInfo{
// 		FirstName: "Badger",
// 		LastName:  "Smith",
// 		Age:       135,
// 	}
//
// 	if err := validate.Struct(user); err != nil {
// 		fmt.Println(Translate(err))
// 	}
// }
//
// // validatorCmd represents the validator command
// var validatorCmd = &cobra.Command{
// 	Use:   "validator",
// 	Short: "A brief description of your command",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("validator called")
// 		validate = validator.New()
// 		validateStruct()
// 	},
// }
//
// // 初始化Validate/v10国际化
// func init() {
// 	rootCmd.AddCommand(validatorCmd)
//
// 	// Here you will define your flags and configuration settings.
//
// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// validatorCmd.PersistentFlags().String("foo", "", "A help for foo")
//
// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// validatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//
// 	zh := CN_ZH.New()
// 	uni := ut.New(zh, zh)
// 	trans, _ = uni.GetTranslator("zh")
//
// 	validate = validator.New()
//
// 	// 通过label标签返回自定义错误内容
// 	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
// 		label := field.Tag.Get("label")
// 		if label == "" {
// 			return field.Name
// 		}
// 		return label
// 	})
// 	zhTranslations.RegisterDefaultTranslations(validate, trans)
//
// 	// 自定义required_if错误内容
// 	validate.RegisterTranslation("required_if", trans, func(ut ut.Translator) error {
// 		return ut.Add("required_if", "{0}为必填字段!", false) // see universal-translator for details
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("required_if", fe.Field())
// 		return t
// 	})
//
// 	validate.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
// 		return ut.Add("lte", "{0}该值错误!", false) // see universal-translator for details
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("lte", fe.Field())
// 		return t
// 	})
// }
