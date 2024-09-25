package cmd

import (
	"github.com/gookit/goutil/dump"
	cregex "github.com/mingrammer/commonregex"
	"github.com/spf13/cobra"
	"mvdan.cc/xurls/v2"
)

// regexCmd represents the regex command
var regexCmd = &cobra.Command{
	Use:   "regex",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// urlRegex := regexp.MustCompile(`(?:(?:(?:https?|ftp):)\/\/)(?:\S+(?::\S*)?@)?(?:(x??!(?:10|127)(?:\.\d{1,3}){3})(x??!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(x??!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z0-9\x{00a1}-\x{ffff}][a-z0-9\x{00a1}-\x{ffff}_-]{0,62})?[a-z0-9\x{00a1}-\x{ffff}]\.)+(?:[a-z\x{00a1}-\x{ffff}]{2,}\.?))(?::\d{2,5})?([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
		// // seperate regex for IP urls since the above did not work
		// ipRegex := regexp.MustCompile(`(?:(?:(?:https?|ftp):)\/\/)\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(:\d{2,5})?`)
		//
		// // find all regex matches and in converted byte data and concat both string slices into single slice
		// textUrls := urlRegex.FindAllString(string(text), -1)
		// ipUrls := ipRegex.FindAllString(string(text), -1)
		//
		// urls := append(textUrls, ipUrls...)
		//
		// // stop reading file
		// // file.Close()
		//
		// urls = utils.RemoveDuplicate(urls)

		text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`

		dateList := cregex.Date(text)
		// ['Jan 9th 2012']
		timeList := cregex.Time(text)
		// ['5:00PM', '4:00']
		linkList := cregex.Links(text)
		// ['www.linkedin.com', 'harold.smith@gmail.com']
		phoneList := cregex.PhonesWithExts(text)
		// ['(519)-236-2723x341']
		emailList := cregex.Emails(text)
		// ['harold.smith@gmail.com']

		dump.Println(dateList, timeList, linkList, phoneList, emailList)

		rxRelaxed := xurls.Relaxed()
		rxRelaxed.FindString("Do gophers live in golang.org?")  // "golang.org"
		rxRelaxed.FindString("This string does not have a URL") // ""

		rxStrict := xurls.Strict()
		rxStrict.FindAllString("must have scheme: http://foo.com/.", -1) // []string{"http://foo.com/"}
		rxStrict.FindAllString("no scheme, no match: foo.com", -1)       // []string{}
	},
}

func init() {
	rootCmd.AddCommand(regexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
