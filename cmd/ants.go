package cmd

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"

	"github.com/spf13/cobra"
)

// antsCmd represents the ants command
var antsCmd = &cobra.Command{
	Use:   "ants",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ants called")
		p, _ := ants.NewPoolWithFunc(10, taskFunc)
		defer p.Release()

	},
}

func init() {
	rootCmd.AddCommand(antsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// antsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// antsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}

	t.wg.Done()
}

func taskFunc(data any) {
	task := data.(*Task)
	task.Do()
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}
