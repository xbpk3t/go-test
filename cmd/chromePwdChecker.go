package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"test/pkg/chromePwdChecker"
)

// chromePwdCheckerCmd represents the chromePwdChecker command
var chromePwdCheckerCmd = &cobra.Command{
	Use:   "chromePwdChecker",
	Short: "检测chrome内置密码管理工具的网站是否能够访问, 适合不使用1password/keypass/lastpass等密码管理工具的用户使用",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chromePwdChecker called")

		checker := chromePwdChecker.NewChromePwdChecker()
		items := checker.ReadCSV("./chrome.csv")
		checker.Checker(items)
	},
}

func init() {
	rootCmd.AddCommand(chromePwdCheckerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chromePwdCheckerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chromePwdCheckerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
