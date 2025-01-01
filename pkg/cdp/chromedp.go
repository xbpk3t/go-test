package cdp

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	Username       = "yyzw@live.com"
	Password       = "F9ezJngtjWMBNuZ"
	HomeURL        = "http://www.glidedsky.com/"
	Lx1URL         = "http://www.glidedsky.com/level/web/crawler-basic-1"
	CookieFilename = "cookie"
)

// @url
func TestLx1(t *testing.T) {

	cookie := ReadCookie()
	sum := GetSum(cookie)

	assert.Equal(t, 246553, sum)
}

func ReadCookie() string {
	_, err := os.Stat(CookieFilename)
	if os.IsNotExist(err) {
		MockLogin()
	}
	cookie, err := os.ReadFile(CookieFilename)
	if err != nil {
		return ""
	}
	return string(cookie)
}

// 模拟登录, 并写入本地
func MockLogin() string {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	cookie := ""

	err := chromedp.Run(ctx,
		chromedp.Tasks{
			// 打开导航
			chromedp.Navigate(HomeURL),
			// 等待元素加载完成
			// chromedp.WaitVisible("body", chromedp.ByQuery),
			// 输入chromedp
			chromedp.SendKeys(`#email`, Username, chromedp.NodeVisible),
			chromedp.SendKeys(`#password`, Password, chromedp.NodeVisible),
			// 提交
			// chromedp.Submit(".gLFyf.gsfi", chromedp.ByQuery),
			chromedp.Click(`//*[@id="app"]/main/div/div/div/div/div[2]/form/div[4]/div/button`, chromedp.NodeVisible),
			chromedp.Sleep(3 * time.Second),
			// 获取cookie
			chromedp.ActionFunc(func(ctx context.Context) error {
				cookies, err := network.GetCookies().Do(ctx)
				if err != nil {
					return err
				}
				for i, v := range cookies {
					cookie += v.Name + "=" + v.Value
					if i != len(cookies)-1 {
						cookie += "; "
					}
				}
				return nil
			}),
		},
	)

	if os.WriteFile(CookieFilename, []byte(cookie), 0644) != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	return cookie
}

// 获取所有数字之和
func GetSum(cookies string) int {

	doc := gq.GetWithCookie(Lx1URL, cookies)
	sum := 0
	doc.Find(".col-md-1").Each(func(i int, selection *goquery.Selection) {
		num, err := strconv.Atoi(strings.TrimSpace(selection.Text()))
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	})
	return sum
}

// 获取网站上爬取的数据
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		log.Printf("Run err : %v\n", err)
		return "", err
	}
	//log.Println(htmlContent)

	return htmlContent, nil
}

func ScreenShot() {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var (
		buf   []byte
		value string
	)

	err := chromedp.Run(ctx,
		chromedp.Tasks{
			// 打开导航
			chromedp.Navigate("https://google.com/"),
			// 等待元素加载完成
			chromedp.WaitVisible("body", chromedp.ByQuery),
			// 输入chromedp
			chromedp.SendKeys(`.gLFyf.gsfi`, "chromedp", chromedp.NodeVisible),
			// 打印输入框的值
			chromedp.Value(`.gLFyf.gsfi`, &value),
			// 提交
			chromedp.Submit(".gLFyf.gsfi", chromedp.ByQuery),
			chromedp.Sleep(3 * time.Second),
			// 截图
			chromedp.CaptureScreenshot(&buf),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("value: ", value)
	if err := os.WriteFile("fullScreenshot.png", buf, 0644); err != nil {
		fmt.Println(err)
	}
}
