package cmd

import (
	"bytes"
	"fmt"
	"github.com/RoaringBitmap/roaring/v2"

	"github.com/spf13/cobra"
)

// bitmapCmd represents the bitmap command
var bitmapCmd = &cobra.Command{
	Use:   "bitmap",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bitmap called")

		fmt.Println("==roaring==")
		rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
		fmt.Println(rb1.String())

		rb2 := roaring.BitmapOf(3, 4, 1000)
		fmt.Println(rb2.String())

		rb3 := roaring.New()
		fmt.Println(rb3.String())

		fmt.Println("Cardinality: ", rb1.GetCardinality())

		fmt.Println("Contains 3? ", rb1.Contains(3))

		rb1.And(rb2)

		rb3.Add(1)
		rb3.Add(5)

		rb3.Or(rb1)

		// computes union of the three bitmaps in parallel using 4 workers
		roaring.ParOr(4, rb1, rb2, rb3)
		// computes intersection of the three bitmaps in parallel using 4 workers
		roaring.ParAnd(4, rb1, rb2, rb3)

		// prints 1, 3, 4, 5, 1000
		i := rb3.Iterator()
		for i.HasNext() {
			fmt.Println(i.Next())
		}
		fmt.Println()

		// next we include an example of serialization
		buf := new(bytes.Buffer)
		rb1.WriteTo(buf) // we omit error handling
		newrb := roaring.New()
		newrb.ReadFrom(buf)
		if rb1.Equals(newrb) {
			fmt.Println("I wrote the content to a byte stream and read it back.")
		}

		// [漫画：什么是Bitmap算法？1. 给定长度是10的bitmap，每一个bit位分别对应着从0到9的10个整型数。此时b - 掘金](https://juejin.cn/post/6844903769201704973)

	},
}

func init() {
	rootCmd.AddCommand(bitmapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bitmapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bitmapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
