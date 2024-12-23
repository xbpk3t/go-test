package red_envelops

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 多图
func MakeCharts(res map[int]map[int]int) error {
	page := components.NewPage()

	for key, val := range res {
		page.AddCharts(MakeLine(key, val))
	}
	f, err := os.Create("charts.html")
	if err != nil {
		panic(err)
	}

	err = page.Render(io.MultiWriter(f))
	if err != nil {
		return err
	}
	return nil
}

// 绘制单个折线图
func MakeLine(num int, val map[int]int) *charts.Line {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: strconv.Itoa(num),
		}))
	line.SetXAxis(GetRange(1, num, 1)).
		AddSeries("Category A", generateLineItems2(val)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false}), charts.WithLabelOpts(opts.Label{Show: true}))

	return line
}

// generate random data for line chart
func generateLineItems2(val map[int]int) []opts.LineData {
	items := make([]opts.LineData, 0)

	for k, v := range val {
		fmt.Println(k, v, "===")
		total := MapSum(val)
		ratio := float64(v) / float64(total)
		items = append(items, opts.LineData{Value: ratio})
	}
	return items
}

func MapSum(input map[int]int) (res int) {
	for _, v := range input {
		res += v
	}
	return res
}

// 获取数字区间的所有数字
func GetRange(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}

	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}
