/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"time"

	"github.com/spf13/cobra"
)

// moCmd represents the mo command
var moCmd = &cobra.Command{
	Use:   "mo",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mo called")

		f := mo.Some(11).FlatMap(func(value int) mo.Option[int] {
			return mo.Some(value * 2)
		}).OrElse(1234)
		fmt.Println(f)

		r := mo.Some(42).Map(func(i int) (int, bool) {
			return 1234, true
		})
		fmt.Println(r.IsPresent(), r.OrEmpty())

		fmt.Println(lo.FindUniques([]int{1, 2, 3, 4, 1}))

		// mo.Some([]int{1, 2, 3, 4}).FlatMap(func(val []int) mo.Option[[]int] {
		// })

		res := mo.Ok(42)
		res.Match(func(i int) (int, error) {
			return i * 2, nil
		}, func(err error) (int, error) {
			return 0, err
		})
		fmt.Println(res)

		rs := divide(42, 1)
		rs.Match(func(val int) (int, error) {
			return val, nil
		}, func(err error) (int, error) {
			return 0, err
		})
		fmt.Println(rs.Get())

		fmt.Println(UseFindValue())
	},
}

type clearTarget struct {
	name     string
	argLong  string
	argShort mo.Option[string]
	location func() string
}

//	func divide(a, b int) (int, error) {
//	   if b == 0 {
//	       return 0, errors.New("division by zero")
//	   }
//	   return a / b, nil
//	}
func divide(a, b int) mo.Result[int] {
	if b == 0 {
		return mo.Err[int](fmt.Errorf("divide by zero"))
	}
	return mo.Ok(a / b)
}

//	func findValue(val int) *int {
//	   if val > 0 {
//	       return &val
//	   }
//	   return nil
//	}
func findValue(val int) mo.Option[int] {
	if val > 0 {
		return mo.Some(val)
	}
	return mo.None[int]()
}

func UseFindValue() int {
	return findValue(42).Match(func(value int) (int, bool) {
		return value, true
	}, func() (int, bool) {
		return 0, false
	}).MustGet()
}

//	func fileExists(path string) (bool, error) {
//	   _, err := os.Stat(path)
//	   if errors.Is(err, os.ErrNotExist) {
//	       return false, nil
//	   } else if err != nil {
//	       return false, err
//	   }
//	   return true, nil
//	}
// func fileExists(path string) mo.IOEither[bool] {
// 	return mo.NewIOEither1(func(path string) (bool, error) {
// 		_, err := os.Stat(path)
// 		if errors.Is(err, os.ErrNotExist) {
// 			return false, nil
// 		} else if err != nil {
// 			return false, err
// 		}
// 		return true, nil
// 	})
// }

// func UseFileExists() {
// 	io := fileExists("./io_either.go")
// 	either := io.Run("./io_either.go")
// 	either.Match(
// 		func(err error) mo.Either[bool, error] {
// 			fmt.Println("Error:", err)
// 			return mo.Left[bool, error](false)
// 		},
// 		func(exist bool) mo.Either[bool, error] {
// 			fmt.Println("File exists:", exist)
// 			return mo.Right[bool, error](exist)
// 		},
// 	)
// }

//	func asyncOperation() (string, error) {
//	   time.Sleep(2 * time.Second)
//	   return "Completed", nil
//	}
func asyncOperation() *mo.Future[string] {
	return mo.NewFuture(func(resolve func(string), reject func(error)) {
		time.Sleep(2 * time.Second)
		resolve("Completed")
	})
}

func UseAsyncOperation() {
	f := asyncOperation()
	f.Result().Match(func(value string) (string, error) {
		return value, nil
	}, func(err error) (string, error) {
		return "", err
	})
}

// var clearTargets = []clearTarget{
// 	{"cache directory", "cache", mo.Some("c"), where.Cache},
// 	{"history file", "history", mo.Some("s"), where.History},
// 	{"anilist binds", "anilist", mo.Some("a"), where.AnilistBinds},
// 	{"queries history", "queries", mo.Some("q"), where.Queries},
// }

func init() {
	rootCmd.AddCommand(moCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
