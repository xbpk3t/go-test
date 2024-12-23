package red_envelops

import "fmt"

const (
	RedEnvelopNormal       = iota // 红包正常，未过期也未抢完
	RedEnvelopFinish              // 红包已抢完
	RedEnvelopNotExist            // 红包不存在
	RedEnvelopExpired             // 红包已过期
	RedEnvelopRepeatUnpack        // 红包不可重复拆
)

var RedEnvelopState = map[int]string{
	RedEnvelopFinish:       "红包已抢完",
	RedEnvelopNotExist:     "红包不存在",
	RedEnvelopRepeatUnpack: "红包不可重复拆",
	RedEnvelopExpired:      "红包已过期",
	RedEnvelopNormal:       "红包正常",
}

func RedEnvelopStateMap(state int) string {
	if desc, exist := RedEnvelopState[state]; exist {
		return desc
	}
	return fmt.Sprintf("error code: %d", state)
}
