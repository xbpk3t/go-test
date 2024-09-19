package httptest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helloHandler(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"base case", `{"name": "liwenzhou"}`, "hello liwenzhou"},
		{"bad case", "", "we need a name"},
	}

	r := SetupRouter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest("POST", "/hello", strings.NewReader(tt.param))
			// mock一个响应记录器
			w := httptest.NewRecorder()
			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			// 断言状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)
			// 解析并断言
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
