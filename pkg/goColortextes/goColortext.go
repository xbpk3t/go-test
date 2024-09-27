package goColortextes

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
)

type (
	LogPriority int
)

const (
	LOG_EMERG LogPriority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERROR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

var Color = struct {
	Emerg  ct.Color
	Alert  ct.Color
	Crit   ct.Color
	Error  ct.Color
	Warn   ct.Color
	Notice ct.Color
	Info   ct.Color
	Debug  ct.Color
	Output ct.Color
}{
	ct.Red,
	ct.Red,
	ct.Red,
	ct.Red,
	ct.Yellow,
	ct.Magenta,
	ct.Green,
	ct.Cyan,
	ct.Cyan,
}

var Prefix = struct {
	Emerg  string
	Alert  string
	Crit   string
	Error  string
	Warn   string
	Notice string
	Info   string
	Debug  string
	Output string
}{
	"[EMERG] ",
	"[ALERT] ",
	"[CRIT] ",
	"[ERR] ",
	"[WARN] ",
	"[NOTICE] ",
	"[INFO] ",
	"[DEBUG] ",
	">> ",
}

var PrintPriority = LOG_DEBUG

// 根据log的等级，在控制台打印不同颜色
func PrintLog(priority LogPriority, ft string, args ...interface{}) {
	if PrintPriority < priority {
		return
	}
	switch priority {
	case LOG_EMERG:
		PrintEmerg(ft, args...)
	case LOG_ALERT:
		PrintAlert(ft, args...)
	case LOG_CRIT:
		PrintCrit(ft, args...)
	case LOG_ERROR:
		PrintError(ft, args...)
	case LOG_WARNING:
		PrintWarn(ft, args...)
	case LOG_NOTICE:
		PrintNotice(ft, args...)
	case LOG_INFO:
		PrintInfo(ft, args...)
	default:
		PrintDebug(ft, args...)
	}
}

func Println(color ct.Color, prefix string, ft string, args ...interface{}) {
	ct.Foreground(color, false)
	fmt.Print(prefix)
	ct.ResetColor()
	fmt.Printf(ft, args...)
	fmt.Println()
}

func PrintEmerg(ft string, args ...interface{}) {
	Println(Color.Emerg, Prefix.Emerg, ft, args...)
}

func PrintAlert(ft string, args ...interface{}) {
	Println(Color.Alert, Prefix.Alert, ft, args...)
}

func PrintCrit(ft string, args ...interface{}) {
	Println(Color.Crit, Prefix.Crit, ft, args...)
}

func PrintError(ft string, args ...interface{}) {
	Println(Color.Error, Prefix.Error, ft, args...)
}

func PrintWarn(ft string, args ...interface{}) {
	Println(Color.Warn, Prefix.Warn, ft, args...)
}

func PrintNotice(ft string, args ...interface{}) {
	Println(Color.Notice, Prefix.Notice, ft, args...)
}

func PrintInfo(ft string, args ...interface{}) {
	Println(Color.Info, Prefix.Info, ft, args...)
}

func PrintDebug(ft string, args ...interface{}) {
	Println(Color.Debug, Prefix.Debug, ft, args...)
}

func PrintOutput(ft string, args ...interface{}) {
	Println(Color.Output, Prefix.Output, ft, args...)
}
