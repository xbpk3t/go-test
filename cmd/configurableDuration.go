/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// [用Bitmap与AST做一个配置化时长系统](https://mp.weixin.qq.com/s?__biz=MzU3NzEwNjI5OA==&mid=2247484732&idx=1&sn=9ec2035f53e6d77b47499a430d2eebfc)
// configurableDurationCmd represents the configurableDuration command
var configurableDurationCmd = &cobra.Command{
	Use:   "configurableDuration",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configurableDuration called")

		expr := `1 + 4 - 2 + 100 - 20 + 12 `
		// expr := ` 1 + 4 `

		ast := getAst(expr)
		result := getResult(ast)

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(configurableDurationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configurableDurationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configurableDurationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	Number   = 0
	Operator = 1
)

type Node struct {
	Type  int
	Value string
	Left  *Node
	Right *Node
}

// input: 1 + 4 - 2
// result:
//
//	 -
//	    /   \
//	 +  2
//	  /   \
//	1  4
func getAst(expr string) *Node {

	operator := make(map[string]int)
	operator["+"] = Operator
	operator["-"] = Operator

	nodeList := make([]Node, 0)
	var root *Node

	expr = strings.Trim(expr, " ")
	words := strings.Split(expr, " ")
	for _, word := range words {

		var node Node

		if _, ok := operator[word]; ok {
			node.Type = Operator
		} else {
			node.Type = Number
		}
		node.Value = word
		nodeList = append(nodeList, node)
	}

	for i := 0; i < len(nodeList); i++ {
		if root == nil {
			root = &nodeList[i]
			continue
		}

		switch nodeList[i].Type {
		case Operator:
			nodeList[i].Left = root
			root = &nodeList[i]
		case Number:
			root.Right = &nodeList[i]
		}
	}

	return root
}

func getResult(node *Node) string {
	switch node.Type {
	case Number:
		return node.Value
	case Operator:
		return calc(getResult(node.Left), getResult(node.Right), node.Value)
	}
	return ""
}

func calc(left, right string, operator string) string {
	leftVal, _ := TransToInt(left)
	rightVal, _ := TransToInt(right)
	val := 0
	switch operator {
	case "+":
		val = leftVal + rightVal
	case "-":
		val = leftVal - rightVal
	}
	return TransToString(val)
}

func TransToString(data interface{}) (res string) {
	val := reflect.ValueOf(data)
	return strconv.FormatInt(val.Int(), 10)
}

func TransToInt(data interface{}) (res int, err error) {
	return strconv.Atoi(strings.TrimSpace(data.(string)))
}
