package validates

import (
	"time"

	"github.com/gookit/validate"
)

// UserForm struct
type UserForm struct {
	Name     string    `validate:"required|minLen:7"`
	Email    string    `validate:"email"`
	Age      int       `validate:"required|int|min:1|max:99"`
	Safe     int       `validate:"-"`
	CreateAt int       `validate:"min:1"`
	UpdateAt time.Time `validate:"required"`
	Code     string    `validate:"customValidator"` // 使用自定义验证器
	Prefix   string    `validate:"required|regex:^\\d{4,6}$"`
	// 结构体嵌套
	ExtInfo struct {
		Homepage string `validate:"required"`
		CityName string
	} `validate:"required"`
}

// CustomValidator 定义在结构体中的自定义验证器
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// ConfigValidation 配置验证
// - 定义验证场景
// - 也可以添加验证设置
func (f UserForm) ConfigValidation(v *validate.Validation) {
	// v.StringRule()

	v.WithScenes(validate.SValues{
		"add":    []string{"ExtInfo.Homepage", "Name", "Code"},
		"update": []string{"ExtInfo.CityName", "Name"},
	})
}

// Messages 您可以自定义验证器错误消息
// func (f UserForm) Messages() map[string]string {
// 	return validate.MS{
// 		"required":      "oh! the {field} is required",
// 		"Name.required": "message for special field",
// 		"minLen":        "{field} 的长度不能小于 {param}",
// 	}
// }

// Translates 你可以自定义字段翻译
func (f UserForm) Translates() map[string]string {
	return validate.MS{
		"Name":             "用户名称",
		"Email":            "用户邮箱",
		"ExtInfo.Homepage": "用户主页",
		"Prefix":           "前缀test",
	}
}
