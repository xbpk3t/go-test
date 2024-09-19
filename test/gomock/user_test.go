package gomock

import (
	"errors"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
)

// [搞定Go单元测试（二）—— mock框架(gomock) - 掘金](https://juejin.cn/post/6844903853532381198)
// 静态设置返回值
func TestReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserInterface(ctrl)
	repo.EXPECT().FindOne(1).Return(&UserStruct{Name: "张三"}, nil)
	repo.EXPECT().FindOne(2).Return(&UserStruct{Name: "李四"}, nil)
	repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))

	// 验证一下结果
	// log.Println(repo.FindOne(1)) // 这是张三
	// log.Println(repo.FindOne(2)) // 这是李四
	// log.Println(repo.FindOne(3)) // user not found
	// log.Println(repo.FindOne(4)) // 没有设置4的返回值，却执行了调用，测试不通过
}

// 动态设置返回值
func TestReturnDynamic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserInterface(ctrl)
	repo.EXPECT().FindOne(gomock.Any()).DoAndReturn(func(i int) (*UserStruct, error) {
		if i == 0 {
			return nil, errors.New("")
		}
		if i < 100 {
			return &UserStruct{
				Name: "小于100",
			}, nil
		} else {
			return &UserStruct{
				Name: "大于100",
			}, nil
		}
	})

	// log.Println(repo.FindOne(120))
	// log.Println(repo.FindOne(66))
	// log.Println(repo.FindOne(0))
}

func TestTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserInterface(ctrl)
	repo.EXPECT().FindOne(1).Return(&UserStruct{Name: "张三"}, nil)
	repo.EXPECT().FindOne(2).Return(&UserStruct{Name: "李四"}, nil).Times(2)
	repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found")).AnyTimes()

	// 验证一下结果
	// log.Println(repo.FindOne(1)) // 这是张三
	// log.Println(repo.FindOne(2)) // 这是李四
	// log.Println(repo.FindOne(2)) // FindOne(2) 需调用两次,注释本行代码将导致测试不通过
	// log.Println(repo.FindOne(3)) // user not found, 不限调用次数，注释掉本行也能通过测试
}

func TestOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserInterface(ctrl)
	o1 := repo.EXPECT().FindOne(1).Return(&UserStruct{Name: "张三"}, nil)
	o2 := repo.EXPECT().FindOne(2).Return(&UserStruct{Name: "李四"}, nil)
	o3 := repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))

	gomock.InOrder(o1, o2, o3)

	log.Println(repo.FindOne(1)) // 这是张三
	log.Println(repo.FindOne(2)) // 这是李四
	log.Println(repo.FindOne(3)) // user not found
}
