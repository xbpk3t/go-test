package gock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

// [Go单测从零到溜系列—1.mock网络测试](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651450021&idx=3&sn=0e167476b501cc0b26a625485cc0e805)
func TestGetResultByAPI(t *testing.T) {
	defer gock.Off()

	t.Run("", func(t *testing.T) {
		gock.New("http://your-api.com").Post("/post").MatchType("json").
			JSON(map[string]int{"x": 1}).Reply(200).JSON(map[string]int{"value": 100})

		res := GetResultByAPI(1, 1)
		assert.Equal(t, res, 101)
	})

	t.Run("", func(t *testing.T) {
		// mock 请求外部api时传参x=2返回200
		gock.New("http://your-api.com").
			Post("/post").
			MatchType("json").
			JSON(map[string]int{"x": 2}).
			Reply(200).
			JSON(map[string]int{"value": 200})

		// 调用我们的业务函数
		res := GetResultByAPI(2, 2)
		// 校验返回结果是否符合预期
		assert.Equal(t, res, 202)

		assert.True(t, gock.IsDone()) // 断言mock被触发
	})
}
