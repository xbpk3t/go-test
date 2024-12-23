package red_envelops

import (
	"math/rand"
	"time"
)

type RedEnvelop struct {
	Count          int   // 个数
	Money          int   // 总金额(分)
	RemainCount    int   // 剩余个数
	RemainMoney    int   // 剩余金额(分)
	BestMoney      int   // 手气最佳金额
	BestMoneyIndex int   // 手气最佳序号
	MoneyList      []int // 拆分列表
}

// 返回红包列表(实际业务是动态抢红包，不需要该方法)
// count 红包个数
// money 红包总金额
func RedEnvelopList(count, total int) []int {
	rand.Seed(time.Now().UnixNano())
	reward := RedEnvelop{
		Count:       count,
		Money:       total,
		RemainCount: count,
		RemainMoney: total,
	}

	// 打印拆包列表，带手气最佳
	for i := 0; reward.RemainCount > 0; i++ {
		money := GrabReward(&reward)
		if money > reward.BestMoney {
			reward.BestMoneyIndex, reward.BestMoney = i, money
		}
		reward.MoneyList = append(reward.MoneyList, money)
	}

	// fmt.Printf("总个数:%d, 总金额:%.2f", reward.Count, float32(reward.Money)/100)
	// for i := range reward.MoneyList {
	//	money := reward.MoneyList[i]
	//	isBest := ""
	//	if reward.BestMoneyIndex == i {
	//		isBest = " ** 手气最佳"
	//	}
	//	fmt.Printf("money_%d : %d %s\n", i+1, money, isBest)
	// }
	// fmt.Println("-------")
	return reward.MoneyList
}

// 分配给某个红包的金额
// todo 目前红包金额均为int，需要调整为float
func GrabReward(reward *RedEnvelop) int {
	if reward.RemainCount <= 0 {
		panic("RemainCount <= 0")
	}
	// 最后一个
	if reward.RemainCount-1 == 0 {
		money := reward.RemainMoney
		reward.RemainCount = 0
		reward.RemainMoney = 0
		return money
	}
	// 是否可以直接0.01
	if (reward.RemainMoney / reward.RemainCount) == 1 {
		money := 1
		reward.RemainMoney -= money
		reward.RemainCount--
		return money
	}

	// 红包算法参考 https://www.zhihu.com/question/22625187
	// 最大可领金额 = 剩余金额的平均值x2 = (剩余金额 / 剩余数量) * 2
	// 领取金额范围 = 0.01 ~ 最大可领金额
	maxMoney := int(reward.RemainMoney/reward.RemainCount) * 2
	rand.Seed(time.Now().UnixNano())
	money := rand.Intn(maxMoney)
	// 防止零
	for money == 0 {
		money = rand.Intn(maxMoney)
	}
	reward.RemainMoney -= money
	// 防止剩余金额负数
	if reward.RemainMoney < 0 {
		money += reward.RemainMoney
		reward.RemainMoney = 0
		reward.RemainCount = 0
	} else {
		reward.RemainCount--
	}
	return money
}
