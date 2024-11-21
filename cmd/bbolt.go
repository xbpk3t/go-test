package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// bboltCmd represents the bbolt command
var bboltCmd = &cobra.Command{
	Use:   "bbolt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bbolt called")
		// 我们的大柜子
		db, err := bolt.Open("./my.db", 0600, nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		// 往db里面插入数据
		err = db.Update(func(tx *bolt.Tx) error {
			// 我们的小柜子
			bucket, err := tx.CreateBucketIfNotExists([]byte("user"))
			if err != nil {
				log.Fatalf("CreateBucketIfNotExists err:%s", err.Error())
				return err
			}
			// 放入东西
			if err = bucket.Put([]byte("hello"), []byte("world")); err != nil {
				log.Fatalf("bucket Put err:%s", err.Error())
				return err
			}
			return nil
		})
		if err != nil {
			log.Fatalf("db.Update err:%s", err.Error())
		}

		// 从db里面读取数据
		err = db.View(func(tx *bolt.Tx) error {
			// 找到柜子
			bucket := tx.Bucket([]byte("user"))
			// 找东西
			val := bucket.Get([]byte("hello"))
			log.Printf("the get val:%s", val)
			val = bucket.Get([]byte("hello2"))
			log.Printf("the get val2:%s", val)
			return nil
		})
		if err != nil {
			log.Fatalf("db.View err:%s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(bboltCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bboltCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bboltCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
