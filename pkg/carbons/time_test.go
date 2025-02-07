package carbons

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon(t *testing.T) {

	t.Run("carbon和time.Time 互相转换", func(t *testing.T) {
		assert.Equal(t, carbon.Now().StdTime(), time.Now())
		assert.Equal(t, carbon.CreateFromStdTime(time.Now()), carbon.Now())
	})

	t.Run("NowFormat()", func(t *testing.T) {
		x := carbon.Now().String()
		fmt.Println(x)
		assert.Equal(t, NowFormat(), x)
	})

	t.Run("Compare", func(t *testing.T) {
		carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15")) // true
	})

	t.Run("StartOfDay()", func(t *testing.T) {
		x := carbon.Now().StartOfDay()
		fmt.Println(x)
		// assert.Equal(t, GetToday().String(), x.String())
	})

	// fmt.Println(TimeToStr(time.Now(), YYYY_MM))
	// fmt.Println(TimeToStr(time.Now(), YYYY_MM_DD_HH_MM))

	// assert.Equal(t, MsToTime(), carbon)

	// t.Run("MsToTime()", func(t *testing.T) {
	// 	toTime, err := MsToTime(strconv.Itoa(carbon.Now().Millisecond()))
	// 	if err != nil {
	// 		return
	// 	}
	// 	t.Log(toTime.String())
	// })

	t.Run("Format()", func(t *testing.T) {
		t.Log(Format(time.Now(), "yyyy-MM-dd"))
		assert.Equal(t, Format(time.Now(), "yyyy-MM-dd"), carbon.Now().ToDateString())
		assert.Equal(t, Format(time.Now(), "yyyy-MM-dd HH:mm:ss"), carbon.Now().ToDateTimeString())
		// ...
	})

	t.Run("TimeToStr()", func(t *testing.T) {
		t.Log(TimeToStr(time.Now(), "yyyy-MM-dd"))
		assert.Equal(t, TimeToStr(time.Now(), "yyyy-MM-dd"), carbon.Now().ToDateString())
	})

	// t.Run("BeforeNowFormat()", func(t *testing.T) {
	// 	BeforeNowFormat("year")
	// })
}

func TestBeforeNowFormat(t *testing.T) {
	t.Run("SubYear()", func(t *testing.T) {
		// 测试年之前的时间
		yearTime := BeforeNowFormat("year", 1, "2006-01-02")
		expectedYear := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
		t.Log(yearTime)
		if yearTime != expectedYear {
			t.Errorf("Expected year time to be %s, but got %s", expectedYear, yearTime)
		}
		ct := carbon.Now().SubYear().ToDateString()
		if yearTime != ct {
			t.Errorf("Expected year time to be %s, but got %s", expectedYear, ct)
		}
	})

	t.Run("SubMonth()", func(t *testing.T) {
		// 测试月之前的时间
		monthTime := BeforeNowFormat("month", 1, "2006-01-02")
		expectedMonth := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
		t.Log(monthTime)
		if monthTime != expectedMonth {
			t.Errorf("Expected month time to be %s, but got %s", expectedMonth, monthTime)
		}
		ct := carbon.Now().SubMonth().ToDateString()
		if monthTime != ct {
			t.Errorf("Expected month time to be %s, but got %s", expectedMonth, ct)
		}
	})

	t.Run("SubMinute()", func(t *testing.T) {
		// 测试分钟之前的时间
		minuteTime := BeforeNowFormat("minute", 5, "2006-01-02 15:04:05")
		expectedMinute := time.Now().Add(-time.Minute * 5).Format("2006-01-02 15:04:05")
		t.Log(minuteTime)
		if minuteTime != expectedMinute {
			t.Errorf("Expected minute time to be %s, but got %s", expectedMinute, minuteTime)
		}
		ct := carbon.Now().SubMinutes(5).ToDateTimeString()
		if minuteTime != ct {
			t.Errorf("Expected minute time to be %s, but got %s", expectedMinute, ct)
		}
	})

	t.Run("SubDay()", func(t *testing.T) {
		// 测试日之前的时间
		dayTime := BeforeNowFormat("day", 1, "2006-01-02")
		expectedDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		t.Log(dayTime)
		if dayTime != expectedDay {
			t.Errorf("Expected day time to be %s, but got %s", expectedDay, dayTime)
		}
		ct := carbon.Now().SubDay().ToDateString()
		if dayTime != ct {
			t.Errorf("Expected day time to be %s, but got %s", expectedDay, ct)
		}
	})

	t.Run("", func(t *testing.T) {
		cts := carbon.Now().SubDay().Timestamp()
		t.Logf("subday ts: %d", cts)
		ct := carbon.Now().SubDay().ToDateTimeString()
		ctFromTs := carbon.CreateFromTimestamp(cts).ToDateTimeString()
		t.Log(ct)
		t.Log(ctFromTs)
		assert.Equal(t, ct, ctFromTs)
	})

	t.Run("GetDiffTime()", func(t *testing.T) {
		// GetDiffTime()
		t.Logf("???: %d", carbon.Now().AddMinutes(10).DiffInMinutes())
		t.Logf("Abs: %d", carbon.Now().AddMinutes(10).DiffAbsInMinutes())

		carbon.Now().AddMinutes(10).DiffInSeconds()
	})

	t.Run("", func(t *testing.T) {

		date := "2025-01-18"
		parsedDate := carbon.ParseByLayout(date, "2006-01-02")
		fmt.Println(parsedDate)

		parsedDate2 := carbon.ParseByFormat(date, "yyyy-MM-dd")
		fmt.Println(parsedDate2)

		fmt.Println(subYear("2020-01-01"))
		fmt.Println(subDays("2020-01-01", 20))
	})

}

// subYear 为了处理同比数据
// func subYear(d string) string {
// 	return subDays(d, carbon.DaysPerNormalYear)
// }
//
// // subDays
// func subDays(d string, days int) string {
// 	return carbon.ParseByLayout(d, "yyyy-MM-dd").SubDays(days).String()
// }

// subYear 为了处理同比数据
func subYear(d string) string {
	parsedDate := carbon.ParseByLayout(d, "2006-01-02")
	return parsedDate.SubYears(1).Format("2006-01-02")
}

// subDays 减去指定的天数
func subDays(d string, days int) string {
	parsedDate := carbon.ParseByLayout(d, "2006-01-02")
	return parsedDate.SubDays(days).Format("2006-01-02")
}
