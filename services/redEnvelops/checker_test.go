package red_envelops

import (
	"math/rand"
	"testing"
	"time"
)

var (
	Money    = 1000
	CountMin = 3
	CountMax = 15
	StatNum  = 1000000
)

// 3到26个人，共24组实验
// 每组实验，金额相同，红包个数不同
// 折线图：x轴是红包个数、y轴是手气最佳的概率
func TestSplit(t *testing.T) {
	res := RedEnvelopStat(Money, StatNum, CountMin, CountMax)
	t.Log(res)
	err := MakeCharts(res)
	if err != nil {
		t.Log(err)
		return
	}
}

// count 红包个数，由
// money 红包总金额
// statNum 试验次数
func RedEnvelopStat(total, statNum, countMin, countMax int) map[int]map[int]int {
	res := make(map[int]map[int]int)

	for count := countMin; count <= countMax; count++ {

		if res[count] == nil {
			res[count] = make(map[int]int)
		}
		for i := 0; i < statNum; i++ {
			rand.Seed(time.Now().UnixNano())
			reward := RedEnvelop{
				Count:       count,
				Money:       total,
				RemainCount: count,
				RemainMoney: total,
			}

			// 打印拆包列表，带手气最佳
			for val := 0; reward.RemainCount > 0; val++ {
				money := GrabReward(&reward)
				if money > reward.BestMoney {
					reward.BestMoneyIndex, reward.BestMoney = val, money
				}
				reward.MoneyList = append(reward.MoneyList, money)
			}

			// fmt.Printf("总个数:%d, 总金额:%.2f\n", reward.Count, float32(reward.Money)/100)
			for key := range reward.MoneyList {
				if reward.BestMoneyIndex == key {
					res[count][key+1] += 1
				}
			}
		}
	}

	return res
}
