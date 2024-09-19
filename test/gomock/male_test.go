package gomock

import (
	"testing"

	"github.com/golang/mock/gomock"
)

// [1.4 使用 Gomock 进行单元测试 - 跟煎鱼学 Go](https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/gomock)
func TestMale_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := NewMockMale(ctl)
	gomock.InOrder(mockMale.EXPECT().Get(id).Return(nil))

	user := NewUser(mockMale)
	err := user.GetUserInfo(id)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}
