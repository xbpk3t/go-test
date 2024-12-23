package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

// pbCmd represents the pb command
var pbCmd = &cobra.Command{
	Use:   "pb",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		// progress := pterm.DefaultProgressbar.WithTotal(100).WithTitle("Processing")
		// for i := 0; i <= 100; i++ {
		// 	progress.Increment()
		// 	time.Sleep(time.Millisecond * 50) // Simulating some work
		// }

		var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
			"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

		// Create progressbar as fork from the default progressbar.
		p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

		for i := 0; i < p.Total; i++ {
			if i == 6 {
				time.Sleep(time.Second * 3) // Simulate a slow download.
			}
			p.UpdateTitle("Downloading " + fakeInstallList[i])         // Update the title of the progressbar.
			pterm.Success.Println("Downloading " + fakeInstallList[i]) // If a progressbar is running, each print will be printed above the progressbar.
			p.Increment()                                              // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
			time.Sleep(time.Millisecond * 350)                         // Sleep 350 milliseconds.
		}
	},
}

func init() {
	rootCmd.AddCommand(pbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
