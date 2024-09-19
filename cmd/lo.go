package cmd

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// loCmd represents the lo command
var loCmd = &cobra.Command{
	Use:   "lo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lo called")
		ints := []int{1, 2, 3, 1, 2}
		res := lo.FindUniques(ints)
		dump.Println(res)

		res = lo.Drop(ints, 2)
		dump.Println(res)

	},
}

func init() {
	rootCmd.AddCommand(loCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type ConfigRepos []ConfigRepo

type ConfigRepo struct {
	Repos []Repository `yaml:"repo"`
	Qs    []string     `yaml:"qs"`
	Md    bool         `yaml:"md,omitempty"`
}

type Repository struct {
	URL  string   `yaml:"url"`
	Name string   `yaml:"name,omitempty"`
	Qs   []string `yaml:"qs,omitempty"`
}

func RemoveElements() {
	var configRepos ConfigRepos
	configRepos = append(configRepos, ConfigRepo{
		Repos: []Repository{
			{URL: "https://mp.weixin.qq.com/", Name: "weixin", Qs: []string{"wqs1", "wqs2", "wqs3"}},
		}, Qs: []string{"qs1", "qs2", "qs3"}, Md: true,
	}, ConfigRepo{
		Repos: []Repository{
			{URL: "https://mp.weixin.qq.com/", Name: "weixin", Qs: []string{"wqs1", "wqs2", "wqs3"}},
		}, Qs: []string{"qs1", "qs2", "qs3"}, Md: false,
	}, ConfigRepo{
		Repos: []Repository{
			{URL: "https://www.cnblogs.com/", Name: "cnblogs"},
		}, Md: true,
	})

	xs := configRepos.FilterMD()
	fmt.Println(xs)
}

// FilterMD Remove element from slice of struct.
func (cr *ConfigRepos) FilterMD() ConfigRepos {
	var filteredConfig ConfigRepos
	for _, crv := range *cr {
		if crv.Md {
			var filteredRepos []Repository
			for _, repo := range crv.Repos {
				if repo.Qs != nil {
					filteredRepos = append(filteredRepos, repo)
				}
			}
			crv.Repos = filteredRepos
			filteredConfig = append(filteredConfig, crv)
		}
	}
	return filteredConfig
}
