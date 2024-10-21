package carbons

import (
	"errors"
	"fmt"
	"time"
)

// BeforeNowFormat timeType是年月日
func BeforeNowFormat(timeType string, ti int, timeFormat string) string {
	now := time.Now()

	if timeType == "year" {
		return now.AddDate(-ti, 0, 0).Format(timeFormat)
	}
	if timeType == "month" {
		return now.AddDate(0, -ti, 0).Format(timeFormat)
	}
	if timeType == "minute" {
		// int要用time.Duration()转一下
		return now.Add(-time.Minute * time.Duration(ti)).Format(timeFormat)
	}

	return now.AddDate(0, 0, -ti).Format(timeFormat)
}

// AfterNowFormat 现在之后的格式化时间
func AfterNowFormat(timeType string, ti int, timeFormat string) string {
	now := time.Now()
	if timeType == "year" {
		return now.AddDate(ti, 0, 0).Format(timeFormat)
	}
	if timeType == "month" {
		return now.AddDate(0, ti, 0).Format(timeFormat)
	}
	if timeType == "minute" {
		return now.Add(time.Minute * time.Duration(ti)).Format(timeFormat)
	}

	return now.AddDate(0, 0, ti).Format(timeFormat)
}

// BeforeNowTimestamp
func BeforeNowTimestamp(timeType string, ti int) int64 {
	now := time.Now()
	if timeType == "year" {
		return now.AddDate(-ti, 0, 0).Unix()
	}
	if timeType == "month" {
		return now.AddDate(0, -ti, 0).Unix()
	}
	if timeType == "minute" {
		// int要用time.Duration()转一下
		return now.Add(-time.Minute * time.Duration(ti)).Unix()
	}

	return now.AddDate(0, 0, ti).Unix()
}

// AfterNowTimestamp 现在之后的时间戳
func AfterNowTimestamp(timeType string, ti int) int64 {
	now := time.Now()
	if timeType == "year" {
		return now.AddDate(ti, 0, 0).Unix()
	}
	if timeType == "month" {
		return now.AddDate(0, ti, 0).Unix()
	}
	if timeType == "minute" {
		return now.Add(time.Minute * time.Duration(ti)).Unix()
	}

	return now.AddDate(0, 0, ti).Unix()
}

// BeforeDaysFormatList 列出之前n天的格式化时间，返回数组（包括当天）
func BeforeDaysFormatList(days int, prefix string) []string {
	var dayArr []string
	for i := 0; i < days; i++ {
		dayArr = append(dayArr, fmt.Sprintf("%s%s", prefix, BeforeNowFormat("day", i, "2006-01-02")))
	}

	return dayArr
}

// GetDiffTime 获得两个时间之间的差值，返回秒（t1-t2）
func GetDiffTime(ti1 string, ti2 string) (int, error) {
	SHORT_TIME := "2006-01-02"
	t1, _ := time.Parse(SHORT_TIME, ti1)
	t2, _ := time.Parse(SHORT_TIME, ti2)

	diff := t1.Sub(t2)
	if diff <= time.Duration(0) {
		return 0, errors.New("时间错误")
	}

	return int(diff), nil
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}
