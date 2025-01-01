/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jxskiss/base62"
	"github.com/sony/sonyflake"
	"sync"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

// snoyflakeCmd represents the snoyflake command
var snoyflakeCmd = &cobra.Command{
	Use:   "sonyflake",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var st sonyflake.Settings // 创建setting实例

		sf, err := sonyflake.New(st) // 创建snoyflake实例

		if err != nil {
			return
		}

		for i := 0; i < 10; i++ {
			wg.Add(1)         // 启动一个goroutine就登记+1
			go generateID(sf) // 使用goroutine来并发执行
		}
		wg.Wait() // 等待所有登记的goroutine都结束（否则会因为main的主协程退出了，其他协程也推出导致不会产生任何id，因此要等全部协程全部执行完）
	},
}

func generateID(sf *sonyflake.Sonyflake) (uid uint64) {

	defer wg.Done() // goroutine结束就登记-1

	uid, err := sf.NextID() // 根据当前sf的情况产生id

	if err != nil {
		return
	}

	ik := base62.EncodeToString(base62.FormatUint(uid))
	fmt.Println(uid, ik)

	return uid
}

func init() {
	rootCmd.AddCommand(snoyflakeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snoyflakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snoyflakeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
