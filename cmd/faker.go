/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/spf13/cobra"
)

// fakerCmd represents the faker command
var fakerCmd = &cobra.Command{
	Use:   "faker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("faker called")

		// NewGenerator()

		ft := GetRandValue([]string{"internet"})
		fmt.Println(ft)
	},
}

func init() {
	rootCmd.AddCommand(fakerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fakerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fakerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Generator struct {
	Locale_ string
	Pkg     string
}

func NewGenerator(localVar ...string) Generator {
	newGenerator := Generator{}
	if len(localVar) > 0 {
		newGenerator.Locale_ = localVar[0]
	} else {
		newGenerator.Locale_ = "zh_CN"
	}
	newGenerator.Pkg = reflect.TypeOf(newGenerator).PkgPath()
	return newGenerator
}

func (g *Generator) SetLanguage(localVar string) {
	g.Locale_ = localVar
}

func (g *Generator) Generics(intended string) string {
	return ""
}

// Get Random Value
func GetRandValue(dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	list := Data[dataVal[0]][dataVal[1]]
	return list[rand.Intn(len(list))]
}

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	_, checkOk = Data[dataVal[0]]
	if len(dataVal) == 2 && checkOk {
		_, checkOk = Data[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

var Data = map[string]map[string][]string{
	// "person": Person,
	//	"contact":  Contact,
	//	"address":  Address,
	// "company": Companys,
	//	"lorem":    Lorem,
	"internet": Internet,
	//	"file":     Files,
	// "job":      Jobs,
	// "color":    Colors,
	// "computer": Computer,
	// "gender":   Genders,
	// "areacode": AreaCodes,
	// "phone":    Phones,
	"email": Emails,
	// "airport":  Airports,
	// "flight":   Flights,
	// "train":    Trains,
	// "seat":     Seats,
	// "carbrand": CarBrands,
	//	"payment":  Payment,
	//	"hipster":  Hipster,
	//	"beer":     Beer,
	//	"hacker":   Hacker,
	//	"currency": Currency,
}

// Internet consists of various internet information
var Internet = map[string][]string{
	"browser":       {"firefox", "chrome", "internetExplorer", "opera", "safari"},
	"domain_suffix": {"com", "biz", "info", "name", "net", "org", "io", "live", "tv"},
	"http_method":   {"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
}

var Emails = map[string][]string{
	"postfix": {"@163.com", "@hotmail.com", "@126.cn", "@gmail.com", "@foxmail.com", "@qq.com", "@sdu.edu.cn", "@123.com", "@yahoo.com", "@msn.com", "@github.com", "@ask.com", "@live.com", "@0355.net", "@139.net", "@126.net", "@3721.net", "@yeah.com"},
}
