package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// structCmd represents the struct command
var structCmd = &cobra.Command{
	Use:   "struct",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tag()

		omitempty()
	},
}

func init() {
	rootCmd.AddCommand(structCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// structCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// structCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 使用 omitempty
// 可以看到，这次输出的 json 中只有 Birth 字段了，string、int、对象类型的字段，都因为没有赋值，默认是零值，所以被忽略，对于日期时间类型，由于不可以设置为零值，也就是 0000-00-00 00:00:00，不会被忽略。
// 需要注意这样的情况：如果一个人的年龄是 0（对于刚出生的婴儿，这个值是合理的），刚好是 int 字段的零值，在添加 omitempty tag 的情况下，年龄字段会被忽略。
func omitempty() {
	type PersonAllowEmpty struct {
		Name     string             `json:",omitempty"`
		Age      int64              `json:",omitempty"`
		Birth    time.Time          `json:",omitempty"`
		Children []PersonAllowEmpty `json:",omitempty"`
	}

	person := PersonAllowEmpty{}
	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes)) // {"Birth":"0001-01-01T00:00:00Z"}
}

// 可以看到，使用 json:"-" 标签的字段都被忽略了
func tag() {
	type Person struct {
		Name     string    `json:"-"`
		Age      int64     `json:"-"`
		Birth    time.Time `json:"-"`
		Children []string  `json:"-"`
	}

	birth, _ := time.Parse(time.RFC3339, "1988-12-02T15:04:27+08:00")
	person := Person{
		Name:     "Wang Wu",
		Age:      30,
		Birth:    birth,
		Children: []string{},
	}

	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes)) // {}
}

func x() {
	type MyStruct struct {
		Field1 *string `json:"field1,omitempty"`
		Field2 *int    `json:"field2,omitempty"`
	}

	field1 := ""
	field2 := 0

	myStruct := MyStruct{
		Field1: &field1, // Non-empty string
		Field2: &field2, // Zero integer
	}

	jsonData, _ := json.Marshal(myStruct)
	fmt.Println(string(jsonData)) // Output: {"field1":null,"field2":null}
}

func z() {
	type User struct {
		Name   string
		Mobile string
		URL    url.URL
	}

	// s := struct {
	// 	*xxx
	// 	Password string
	// }{
	// 	Name: "xxx",
	// }

	s, _ := json.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User:     &User{Name: "xxx", Mobile: "2024/4/12", URL: url.URL{Scheme: "https", Host: "www.google.com"}},
		Password: true,
	})
	fmt.Println(string(s))

	// 用字符串传递数字
	type TestObject struct {
		Field1 int `json:",string"`
	}
	object := TestObject{
		Field1: 1,
	}
	fmt.Println(object)
	os, _ := json.Marshal(object)
	fmt.Println(string(os))

	// f, _, _ := big.ParseFloat("1", 10, 64, big.ToZero)
	// val := map[*big.Float]string{f: "2"}
	// str, err := MarshalToString(val)

	type TestObject2 struct {
		Field1 string
		Field2 json.RawMessage
	}
	var data TestObject2
	json.Unmarshal([]byte(`{"field1": "hello", "field2": [1,2,3]}`), &data)
	fmt.Println(string(data.Field2))

	fmt.Println(slices.Contains([]string{"hello", "world"}, data.Field1))

	fmt.Println(strings.NewReader("xxx").Len())
}

// [Go json 反序列化 interface{}对 int64 处理 | Daryl's Blog](https://darylliu.github.io/archives/a7f0b68f.html)
func v() {
	request := `{"id":7044144249855934983,"name":"demo"}`
	var test interface{}
	decoder := json.NewDecoder(strings.NewReader(request))
	decoder.UseNumber()

	err := decoder.Decode(&test)
	if err != nil {
		return
	}

	res, err := json.Marshal(test)
	if err != nil {
		return
	}
	// {"id":7044144249855934983,"name":"demo"}
	fmt.Println(string(res))
}

// 使用 jsoniter 处理int64
// func r() {
// 	request := `{"id":7044144249855934983,"name":"demo"}`
// 	var test interface{}
// 	decoder := jsoniter.NewDecoder(strings.NewReader(request))
// 	decoder.UseNumber()
//
// 	err := decoder.Decode(&test)
// 	if err != nil {
// 		return
// 	}
//
// 	res, err := jsoniter.Marshal(test)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(string(res))
// }
