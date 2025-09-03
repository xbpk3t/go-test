package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"sync"
	"testing"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"

	"github.com/bwmarrin/snowflake"
)

// 创建高并发场景下订单号

type OrderNo struct {
	TimeStamp string // 时间戳，精确到毫秒
	Merchant  string // 商户号，3位
	RandStr   string // 随机数，2位
}

var (
	orderNo OrderNo
	wg      sync.WaitGroup
)

// 复现一下高并发下生成重复uuid的情况
// 随机数2位：总数: 94745 重复: 84749 去重后: 9996
// 随机数5位：总数: 49643 重复: 31917 去重后: 17726
// 随机数10位：总数: 92624 重复: 74023 去重后: 18601
// 显而易见，这种传统方法在高并发场景下，订单号很容易重复
func BenchmarkOrder(b *testing.B) {
	orderNo.Merchant = grand.Digits(3)

	ss := []string{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			no := fmt.Sprintf("%s%s%s", orderNo.timestamp(), orderNo.Merchant, grand.Digits(10))
			ss = append(ss, no)
			defer wg.Done()
		}()
	}
	ssLen := len(ss)
	szLen := garray.NewStrArrayFrom(ss).Unique().Len()

	b.Logf("总数: %d \n重复: %d \n去重后: %d \n", ssLen, ssLen-szLen, szLen)

	wg.Wait()
}

func (OrderNo) timestamp() string {
	return gtime.Now().Format("YYYYMMDDHHMMSSSSS")
}
func (OrderNew) timestamp() string {
	return gtime.Now().Format("YYYYMMDDHHMMSS")
}
func (OrderUser) timestamp() string {
	return gtime.Now().Format("YYYYMMDDHHMMSS")
}

// 毫秒仅保留三位
// 加上用户id，但是用户顺序id会很长，且会暴露用户id
type OrderNew struct {
	Timestamp string
	RandStr   string
}

var orderNew OrderNew

// 总数: 671642
// 重复: 0
// 去重后: 671642
// 时间戳+snowflake
// 支持订单号唯一和高并发，但是订单号太长，14+19位，共33位
func BenchmarkOrderNew(b *testing.B) {
	ss := []string{}

	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		go func(node *snowflake.Node) {
			n := node.Generate().String()

			// 注意发号器要写到外面，循环里取号，否则就有问题
			// no := fmt.Sprintf("%s%s%s", orderNew.Channel, orderNew.timestamp(), orderNew.rand())
			no := fmt.Sprintf("%s%s", orderNew.timestamp(), n)
			ss = append(ss, no)
			defer wg.Done()
		}(node)
	}

	ssLen := len(ss)

	szLen := garray.NewStrArrayFrom(ss).Unique().Len()

	b.Logf("总数: %d \n重复: %d \n去重后: %d \n", ssLen, ssLen-szLen, szLen)

	wg.Wait()
}

func (OrderNew) rand() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	n := node.Generate().String()
	return n
}

type OrderUser struct {
	Timestamp string
	UserHash  string
}

var orderUser OrderUser

// 时间戳+用户id的方案也不可行，因为取用户id前n位的碰撞率其实很高
// 只能说把完整用户id按照算法转换之后，再写入订单号
// 总数: 10000000
// 重复: 212
// 去重后: 9999788
// 碰撞率很低
func BenchmarkOrderUser(b *testing.B) {
	list := orderUser.userIdList()

	ss := []string{}

	for _, v := range list {

		no := fmt.Sprintf("%s%s", orderUser.timestamp(), orderUser.rand(v))
		ss = append(ss, no)
	}

	ssLen := len(ss)
	szLen := garray.NewStrArrayFrom(ss).Unique().Len()

	b.Logf("总数: %d \n重复: %d \n去重后: %d \n", ssLen, ssLen-szLen, szLen)
}

func (OrderUser) userIdList() []string {
	ss := []string{}
	sid := ""
	for i := 1; i <= 1000000; i++ {
		id := strconv.Itoa(i)
		// 如果id长度大于4，取后4位
		if len(id) <= 4 {
			sid = fmt.Sprintf("%07s", id)
		} else {
			sid = id
		}
		// 如果id长度小于4，补全4位
		ss = append(ss, sid)
	}

	return ss
}

// 只使用fnv算法，长度最长10位，不满10位的前面用0填充，是可以的。也可以用位运算转为固定长度
// [FNV哈希算法 - Hotsum - 博客园](https://www.cnblogs.com/zr520/p/5308353.html)
func (OrderUser) rand(v string) string {
	ik := strconv.Itoa(int(hash(v)))
	iki, err := strconv.Atoi(ik)
	if err != nil {
		return ""
	}
	i := (iki >> 24) ^ (iki & 0xFFFFFF)

	return string(rune(i))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func TestW(t *testing.T) {
	// s := big.NewInt(202105101455400001111111)
	s := 20210510145540000
	// t.Log(len(s))

	i := (s >> 24) ^ (s & 0xFFFFFF)
	t.Log(i)
}
