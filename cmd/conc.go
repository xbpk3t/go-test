/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"test/pkg/concs"
)

// concCmd represents the conc command
var concCmd = &cobra.Command{
	Use:   "conc",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

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

		// p()

		// var wg conc.WaitGroup
		//
		// wg.Go(func() {
		// 	// 模拟任务
		// 	fmt.Println("Task 1")
		// })
		//
		// wg.Go(func() {
		// 	// 模拟任务
		// 	fmt.Println("Task 2")
		// })
		//
		// if err := wg.WaitAndRecover(); err != nil {
		// 	fmt.Println("Error occurred:", err)
		// }

		res := concs.SingleFlight("key", func() any {
			return "result"
		})
		fmt.Println(res)
	},
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
