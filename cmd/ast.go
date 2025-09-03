/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var shortNameMap = map[string]string{
	"error":         "err",
	"package":       "pkg",
	"request":       "req",
	"configuration": "config",
}

// astCmd represents the ast command
var astCmd = &cobra.Command{
	Use:   "ast",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ast called")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		filePath := filepath.Join(homeDir, "Desktop/test/cmd/ast_test.go")

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, filePath, nil, 0)
		if err != nil {
			log.Fatal(err)
		}

		ast.Inspect(node, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.Ident:
				for longName, shortName := range shortNameMap {
					if x.Name == longName {
						log.Printf("Suggest replacing '%s' with '%s' at %v", longName, shortName, fset.Position(x.Pos()))
					}
				}
			case *ast.FuncDecl:
				for longName, shortName := range shortNameMap {
					// 检查函数名是否包含长名称
					if strings.Contains(strings.ToLower(x.Name.Name), longName) {
						log.Printf("Suggest replacing '%s' with '%s' in function name '%s' at %v",
							longName, shortName, x.Name.Name, fset.Position(x.Name.NamePos))
					}
				}
			}
			return true
		})
	},
}

func init() {
	rootCmd.AddCommand(astCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// astCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// astCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
