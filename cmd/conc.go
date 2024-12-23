/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/stream"
	"time"

	"github.com/spf13/cobra"
)

// concCmd represents the conc command
var concCmd = &cobra.Command{
	Use:   "conc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("conc called")

		// var wg conc.WaitGroup
		// defer wg.Wait()
		//
		// // startTheThing(&wg)
		//
		// for i := 0; i < 100; i++ {
		// 	wg.Go(func() {
		// 		fmt.Println("Hello World")
		// 		time.Sleep(1 * time.Second)
		// 	})
		// }

		p()
	},
}

func startTheThing(wg *conc.WaitGroup) {
	wg.Go(func() {
		fmt.Println("Starting Thing...")
	})
}

// func process(values []int)  {
// 	iter.ForEach(values, func)
// }

func mapStream(
	in chan int,
	out chan int,
	f func(int) int,
) {
	s := stream.New().WithMaxGoroutines(10)
	for elem := range in {
		elem := elem
		s.Go(func() stream.Callback {
			res := f(elem)
			return func() { out <- res }
		})
	}
	s.Wait()
}

// 用conc实现多线程轮流打印
func p() {
	const (
		numGoroutines = 4
		totalPrints   = 100
	)

	// 创建一个channel用于控制打印顺序
	ch := make(chan int, 1)
	wg := conc.NewWaitGroup()

	// 创建4个协程
	for i := 0; i < numGoroutines; i++ {
		id := i + 1
		wg.Go(func() {
			for j := id; j <= totalPrints; j += numGoroutines {
				// 等待轮到自己的回合
				current := <-ch
				if current%numGoroutines+1 == id {
					fmt.Println(id)
					time.Sleep(time.Second)
					// 通知下一个协程
					ch <- current + 1
				}
			}
		})
	}

	// 启动第一次打印
	ch <- 0

	// 等待所有协程完成
	wg.Wait()
}

func init() {
	rootCmd.AddCommand(concCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// concCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// concCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
