package carbons

import (
	"strconv"
	"strings"
	"time"
)

type DateStyle string

const (
	YYYYMM            = "yyyyMM"
	YYYYMMDD          = "yyyyMMdd"
	YYYYMMDDHHMMSS    = "yyyyMMddHHmmss"
	YYYYMMDDHHMMSSSSS = "yyyyMMddHHmmssSSS"
	YYYYMMDDHHMM      = "yyyyMMddHHmm"
	YYYYMMDDHH        = "yyyyMMddHH"
	// not used
	// YYMMDDHHMM        = "yyMMddHHmm"

	MM_DD                   = "MM-dd"
	YYYY_MM                 = "yyyy-MM"
	YYYY_MM_DD              = "yyyy-MM-dd"
	MM_DD_HH_MM             = "MM-dd HH:mm"
	MM_DD_HH_MM_SS          = "MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM        = "yyyy-MM-dd HH:mm"
	YYYY_MM_DD_HH_MM_SS     = "yyyy-MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS = "yyyy-MM-dd HH:mm:ss:SSS"

	MM_DD_EN                   = "MM/dd"
	YYYY_MM_EN                 = "yyyy/MM"
	YYYY_MM_DD_EN              = "yyyy/MM/dd"
	MM_DD_HH_MM_EN             = "MM/dd HH:mm"
	MM_DD_HH_MM_SS_EN          = "MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_EN        = "yyyy/MM/dd HH:mm"
	YYYY_MM_DD_HH_MM_SS_EN     = "yyyy/MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS_EN = "yyyy/MM/dd HH:mm:ss.SSS"

	MM_DD_CN               = "MM月dd日"
	YYYY_MM_CN             = "yyyy年MM月"
	YYYY_MM_DD_CN          = "yyyy年MM月dd日"
	MM_DD_HH_MM_CN         = "MM月dd日 HH:mm"
	MM_DD_HH_MM_SS_CN      = "MM月dd日 HH:mm:ss"
	YYYY_MM_DD_HH_MM_CN    = "yyyy年MM月dd日 HH:mm"
	YYYY_MM_DD_HH_MM_SS_CN = "yyyy年MM月dd日 HH:mm:ss"

	HH_MM       = "HH:mm"
	HH_MM_SS    = "HH:mm:ss"
	HH_MM_SS_MS = "HH:mm:ss.SSS"
)

// NowFormat 当前时间戳字符串
func NowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetToday() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
	// return time.Now().Round(24 * time.Hour).Truncate(24 * time.Hour)
}

// MsToTime 毫秒字符串转time.Time
func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	return tm, nil
}

// Format 格式化后的字符串
// format time like java, such as: yyyy-MM-dd HH:mm:ss
// t：时间
// format：格式化字符串
// 返回值：
func Format(t time.Time, format string) string {
	// year
	if strings.ContainsAny(format, "y") {
		year := strconv.Itoa(t.Year())

		if strings.Count(format, "yy") == 1 && strings.Count(format, "y") == 2 {
			format = strings.Replace(format, "yy", year[2:], 1)
		} else if strings.Count(format, "yyyy") == 1 && strings.Count(format, "y") == 4 {
			format = strings.Replace(format, "yyyy", year, 1)
		} else {
			panic("format year error! please 'yyyy' or 'yy'")
		}
	}

	// month
	if strings.ContainsAny(format, "M") {
		var month string

		if int(t.Month()) < 10 {
			month = "0" + strconv.Itoa(int(t.Month()))
		} else {
			month = strconv.Itoa(int(t.Month()))
		}

		if strings.Count(format, "MM") == 1 && strings.Count(format, "M") == 2 {
			format = strings.Replace(format, "MM", month, 1)
		} else {
			panic("format month error! please 'MM'")
		}
	}

	// day
	if strings.ContainsAny(format, "d") {
		var day string

		if t.Day() < 10 {
			day = "0" + strconv.Itoa(t.Day())
		} else {
			day = strconv.Itoa(t.Day())
		}

		if strings.Count(format, "dd") == 1 && strings.Count(format, "d") == 2 {
			format = strings.Replace(format, "dd", day, 1)
		} else {
			panic("format day error! please 'dd'")
		}
	}

	// hour
	if strings.ContainsAny(format, "H") {
		var hour string

		if t.Hour() < 10 {
			hour = "0" + strconv.Itoa(t.Hour())
		} else {
			hour = strconv.Itoa(t.Hour())
		}

		if strings.Count(format, "HH") == 1 && strings.Count(format, "H") == 2 {
			format = strings.Replace(format, "HH", hour, 1)
		} else {
			panic("format hour error! please 'HH'")
		}
	}

	// minute
	if strings.ContainsAny(format, "m") {
		var minute string

		if t.Minute() < 10 {
			minute = "0" + strconv.Itoa(t.Minute())
		} else {
			minute = strconv.Itoa(t.Minute())
		}
		if strings.Count(format, "mm") == 1 && strings.Count(format, "m") == 2 {
			format = strings.Replace(format, "mm", minute, 1)
		} else {
			panic("format minute error! please 'mm'")
		}
	}

	// second
	if strings.ContainsAny(format, "s") {
		var second string

		if t.Second() < 10 {
			second = "0" + strconv.Itoa(t.Second())
		} else {
			second = strconv.Itoa(t.Second())
		}

		if strings.Count(format, "ss") == 1 && strings.Count(format, "s") == 2 {
			format = strings.Replace(format, "ss", second, 1)
		} else {
			panic("format second error! please 'ss'")
		}
	}

	return format
}

// TimeToStr time.Time转字符串
// 其他语言通用的日期格式
func TimeToStr(date time.Time, dateStyle DateStyle) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)

	return date.Format(layout)
}

// func TestFormatDate(t *testing.T) {
// 	fixedTime := "2000-01-02 03:04:05"
// 	trans, _ := gtime.StrToTimeFormat(fixedTime, "Y-m-d H:i:s")
// 	transTime := trans.Time
//
// 	type args struct {
// 		date      time.Time
// 		dateStyle DateStyle
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{"", args{date: transTime, dateStyle: HH_MM}, "03:04"},
// 		{"", args{date: transTime, dateStyle: HH_MM_SS}, "03:04:05"},
// 		{"", args{date: transTime, dateStyle: HH_MM_SS_MS}, "03:04:05.000"},
//
// 		{"", args{date: transTime, dateStyle: MM_DD}, "01-02"},
// 		{"", args{date: transTime, dateStyle: MM_DD_CN}, "01月02日"},
// 		{"", args{date: transTime, dateStyle: MM_DD_EN}, "01/02"},
// 		{"", args{date: transTime, dateStyle: MM_DD_HH_MM}, "01-02 03:04"},
// 		{"", args{date: transTime, dateStyle: MM_DD_HH_MM_CN}, "01月02日 03:04"},
// 		{"", args{date: transTime, dateStyle: MM_DD_HH_MM_EN}, "01/02 03:04"},
// 		{"", args{date: transTime, dateStyle: MM_DD_HH_MM_SS}, "01-02 03:04:05"},
// 		{"", args{date: transTime, dateStyle: MM_DD_HH_MM_SS_CN}, "01月02日 03:04:05"},
// 		{"", args{date: transTime, dateStyle: MM_DD_HH_MM_SS_EN}, "01/02 03:04:05"},
//
// 		{"", args{date: transTime, dateStyle: YYYYMM}, "200001"},
// 		{"", args{date: transTime, dateStyle: YYYYMMDD}, "20000102"},
// 		{"", args{date: transTime, dateStyle: YYYYMMDDHH}, "2000010203"},
// 		{"", args{date: transTime, dateStyle: YYYYMMDDHHMM}, "200001020304"},
// 		{"", args{date: transTime, dateStyle: YYYYMMDDHHMMSS}, "20000102030405"},
//
// 		{"", args{date: transTime, dateStyle: YYYY_MM}, "2000-01"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_CN}, "2000年01月"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD}, "2000-01-02"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_CN}, "2000年01月02日"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_EN}, "2000/01/02"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM}, "2000-01-02 03:04"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_CN}, "2000年01月02日 03:04"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_EN}, "2000/01/02 03:04"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_SS}, "2000-01-02 03:04:05"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_SS_CN}, "2000年01月02日 03:04:05"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_SS_EN}, "2000/01/02 03:04:05"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_SS_SSS}, "2000-01-02 03:04:05:000"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD_HH_MM_SS_SSS_EN}, "2000/01/02 03:04:05.000"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_EN}, "2000/01"},
// 		{"", args{date: transTime, dateStyle: YYYY_MM_DD}, "2000-01-02"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := TimeToStr(tt.args.date, tt.args.dateStyle); got != tt.want {
// 				t.Errorf("FormatDate() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
