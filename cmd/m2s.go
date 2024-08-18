package cmd

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/mitchellh/mapstructure"

	"github.com/spf13/cobra"
)

// User 定义了用户信息的结构
type User struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

// Address 定义了用户的地址信息
type Address struct {
	Street string
	City   string
	Zip    string
}

// m2sCmd represents the m2s command
var m2sCmd = &cobra.Command{
	Use:   "m2s",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// 这是原始的用户信息，存储在map的切片中
		usersMap := []map[string]interface{}{
			{
				"Name":  "John Doe",
				"Age":   28,
				"Email": "john.doe@example.com",
				"Address": map[string]interface{}{
					"Street": "123 Elm St",
					"City":   "Metropolis",
					"Zip":    "12345",
				},
			},
			{
				"Name":  "Jane Smith",
				"Age":   32,
				"Email": "jane.smith@example.com",
				"Address": map[string]interface{}{
					"Street": "456 Oak St",
					"City":   "Smalltown",
					"Zip":    "67890",
				},
			},
		}

		var users []User

		for _, userMap := range usersMap {
			var user User

			// 将Address部分的map单独解码到Address结构体
			if addrMap, ok := userMap["Address"].(map[string]interface{}); ok {
				mapstructure.Decode(addrMap, &user.Address)
			}

			// 将剩余的map解码到User结构体
			err := mapstructure.Decode(userMap, &user)
			if err != nil {
				fmt.Printf("Error decoding user: %s\n", err)
				continue
			}

			// 将解码后的用户信息添加到切片中
			users = append(users, user)
		}

		// 打印转换后的用户信息
		// for _, user := range users {
		// 	fmt.Printf("Name: %s, Age: %d, Email: %s\n", user.Name, user.Age, user.Email)
		// 	fmt.Printf("Address: %s, %s, %s\n", user.Address.Street, user.Address.City, user.Address.Zip)
		// }

		dump.Println(users)
	},
}

func init() {
	rootCmd.AddCommand(m2sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// m2sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// m2sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
