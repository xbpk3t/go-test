/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Employee struct {
	Name     string
	Position string
	Salary   int
	Joined   time.Time
}

// goPrettyCmd represents the goPretty command
var goPrettyCmd = &cobra.Command{
	Use:   "goPretty",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 创建表格写入器
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("员工薪资报表")

		// 2. 设置表格样式 (内置6种主题)
		t.SetStyle(table.StyleRounded)

		// 3. 添加表头
		t.AppendHeader(table.Row{"ID", "姓名", "职位", "薪资", "入职时间", "状态"})

		// 4. 准备数据 (结构体切片)
		employees := []Employee{
			{"张三", "高级工程师", 25000, time.Date(2020, 3, 15, 0, 0, 0, 0, time.UTC)},
			{"李四", "产品经理", 32000, time.Date(2019, 8, 22, 0, 0, 0, 0, time.UTC)},
			{"王五", "UI设计师", 18000, time.Date(2022, 1, 10, 0, 0, 0, 0, time.UTC)},
		}

		// 5. 添加行数据 (自动转换结构体)
		for i, emp := range employees {
			status := text.FgGreen.Sprint("在职")
			if i == 1 {
				status = text.FgYellow.Sprint("休假中")
			}

			t.AppendRow(table.Row{
				i + 1,
				emp.Name,
				emp.Position,
				fmt.Sprintf("¥%d", emp.Salary), // 格式化数字
				emp.Joined.Format("2006-01-02"),
				status, // 彩色状态
			})
		}

		// 6. 添加汇总行 (自动计算)
		t.AppendFooter(table.Row{"", "", "总计", calculateTotal(employees), "", ""})

		// 7. 添加分页标记
		t.SetPageSize(5) // 每页显示行数

		// 8. 设置列配置
		t.Style().Options.SeparateRows = true
		t.Style().Format.Header = text.FormatUpper
		t.SetColumnConfigs([]table.ColumnConfig{
			{Name: "薪资", Align: text.AlignRight, AlignFooter: text.AlignRight},
			{Name: "入职时间", Align: text.AlignCenter},
		})

		// 9. 渲染表格 (支持多种输出)
		fmt.Println("\n===== 终端表格输出 =====")
		t.Render() // 默认终端输出

		fmt.Println("\n===== Markdown 格式 =====")
		t.RenderMarkdown() // Markdown格式输出
	},
}

func calculateTotal(emps []Employee) string {
	total := 0
	for _, emp := range emps {
		total += emp.Salary
	}
	return fmt.Sprintf("¥%d", total)
}

func init() {
	rootCmd.AddCommand(goPrettyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goPrettyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goPrettyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
