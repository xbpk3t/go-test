package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"

	"github.com/spf13/cobra"
)

// restyCmd represents the resty command
var restyCmd = &cobra.Command{
	Use:   "resty",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resty called")
		restyX()
	},
}

func init() {
	rootCmd.AddCommand(restyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 用resty请求接口，对比一下用net/http的方案
// resty会自动关闭conn，除了SetDoNotParseResponse方法
func restyX() {
	client := resty.New()

	resp, err := client.R().EnableTrace().Get("https://api.myrightone.com/api/feed/moment-list?last_object_id=&num=20&sign=4d5d88e2ab4790282da36627ec638b85f996279f&start=0&timestamp=1570804569&type=recommend&user_id=8443055")

	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
}

// [Go 每日一库之 resty](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651450184&idx=3&sn=ed4a6c84e6adfd8a18cc42810420829d)
func restyZ() {
	client := resty.New()
	libraries := &Libraries{}
	client.R().SetResult(libraries).Get("https://api.cdnjs.com/libraries")

	fmt.Printf("%d libraries\n", len(libraries.Results))

	for _, lib := range libraries.Results {
		fmt.Println("first library:")
		fmt.Printf("name:%s latest:%s\n", lib.Name, lib.Latest)
		break
	}
}

type Library struct {
	Name   string
	Latest string
}

type Libraries struct {
	// slice要用复数
	Results []*Library
}
