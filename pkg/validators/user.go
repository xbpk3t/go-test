package validators

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

type Student struct {
	Name  string `validate:required`
	Email string `validate:"email"`
	Age   int    `validate:"max=30,min=12"`
}

// [golang常用库：字段参数验证库-validator使用 - 九卷 - 博客园](https://www.cnblogs.com/jiujuan/p/13823864.html#224880901)
func V() {
	en := en.New() // 英文翻译器
	zh := zh.New() // 中文翻译器

	// 第一个参数是必填，如果没有其他的语言设置，就用这第一个
	// 后面的参数是支持多语言环境（
	// uni := ut.New(en, en) 也是可以的
	// uni := ut.New(en, zh, tw)
	uni := ut.New(en, zh)
	trans, _ := uni.GetTranslator("zh") // 获取需要的语言

	student := Student{
		Name:  "tom",
		Email: "testemal",
		Age:   40,
	}
	validate := validator.New()

	zhtrans.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(student)
	if err != nil {
		// fmt.Println(err)

		// errs := err.(validator.ValidationErrors)
		// fmt.Println(errs.Translate(trans))
		// fmt.Println(removeStructName(errs.Translate(trans)))

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
		}
	}
}

func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}
