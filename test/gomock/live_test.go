package gomock

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

// [Go单测从零到溜系列—3.接口测试](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651450075&idx=3&sn=975dd50a7fe97cdfd74674a71dbe5a73)
//
// [如何有效地测试Go代码](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651445612&idx=4&sn=9da9662a5751d23fd903602e38e69ac7)
//
// [golang 打桩，mock 数据怎么玩？ - 掘金](https://juejin.cn/post/7024781846793682981)
//
// [Go Mock (gomock)简明教程 | 快速入门 | 极客兔兔](https://geektutu.com/post/quick-gomock.html#2-%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E7%9A%84-Demo)

// [Golang后台单元测试实践 - 掘金](https://juejin.cn/post/6917956015132672007#heading-28)

// [白话Golang单元测试 | 火丁笔记](https://blog.huoding.com/2021/11/28/968)
func TestLive(t *testing.T) {
	ctrl := gomock.NewController(t)
	life := NewMockLife(ctrl)
	handler := func(money int64) error {
		if money <= 0 {
			return errors.New("error")
		}
		return nil
	}

	life.EXPECT().GoodGoodStudy(gomock.Any()).AnyTimes().DoAndReturn(handler)
	life.EXPECT().BuyHouse(gomock.Any()).AnyTimes().DoAndReturn(handler)
	life.EXPECT().Marry(gomock.Any()).AnyTimes().DoAndReturn(handler)

	Convey("live", t, func() {
		person := &Person{
			Life: life,
		}
		Convey("GoodGoodStudy error", func() {
			So(person.Live(0, 100, 100), ShouldBeError)
		})
		Convey("GoodGoodStudy success", func() {
			Convey("BuyHouse error", func() {
				So(person.Live(100, 0, 100), ShouldBeError)
			})
			Convey("BuyHouse success", func() {
				Convey("Marry error", func() {
					So(person.Live(100, 100, 0), ShouldBeError)
				})
				Convey("Marry ok", func() {
					So(person.Live(100, 100, 100), ShouldBeNil)
				})
			})
		})
	})
}

func TestGoodGoodStudy(t *testing.T) {
	ctrl := gomock.NewController(t)
	life := NewMockLife(ctrl)
	handler := func(money int64) error {
		if money <= 0 {
			return errors.New("error")
		}
		return nil
	}
	life.EXPECT().GoodGoodStudy(gomock.Any()).AnyTimes().DoAndReturn(handler)
	Convey("GoodGoodStudy", t, func() {
		person := &Person{
			Life: life,
		}
		Convey("", func() {
			So(person.GoodGoodStudy(0), ShouldBeError)
		})
	})
}

func TestPerson_BuyHouse(t *testing.T) {
	ctrl := gomock.NewController(t)
	life := NewMockLife(ctrl)
	handler := func(money int64) error {
		if money <= 0 {
			return errors.New("error")
		}
		return nil
	}
	life.EXPECT().BuyHouse(gomock.Any()).AnyTimes().DoAndReturn(handler)
	Convey("BuyHouse", t, func() {
		person := &Person{
			Life: life,
		}
		Convey("Error", func() {
			So(person.BuyHouse(100), ShouldBeError)
		})
	})
}
