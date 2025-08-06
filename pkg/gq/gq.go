package gq

import (
	"fmt"
	query "github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

// FetchHTML 获取网页
func FetchHTML(url string) *query.Document {
	resp, err := ghttp.NewClient().Get(url)
	if err != nil {
		return &query.Document{}
	}

	// defer func(Body io.ReadCloser) {
	// 	if err := Body.Close(); err != nil {
	// 		logrus.WithFields(log.Text(url, nil)).Error("http close failed")
	// 	}
	// }(resp.Body)

	return DocQuery(resp.Response)
}

// PostHTML 发送表单请求
func PostHTML(url string, m map[string]interface{}) *query.Document {
	resp, err := ghttp.NewClient().Post(url, m)
	if err != nil {
		return nil
	}
	// defer func(Body io.ReadCloser) {
	// 	if err := Body.Close(); err != nil {
	// 		logrus.WithFields(log.Text(url, nil)).Error("http close failed")
	// 	}
	// }(resp.Response.Body)

	return DocQuery(resp.Response)
}

// 请求goquery
func DocQuery(resp *http.Response) *query.Document {
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		return &query.Document{}
	}

	return doc
}

// connection reset by peer
func GetWithCookie(url string, cookie string) *query.Document {

	resp, err := ghttp.NewClient().SetCookie("", cookie).Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// resp.RawDump()

	return DocQuery(resp.Response)
}

// func GetWithCookie(url string, cookie string) *query.Document {
//
// 	client := resty.New().SetDebug(true)
//
// 	tr := &http.Transport{
// 		MaxIdleConns:        20,
// 		MaxIdleConnsPerHost: 20,
// 		MaxConnsPerHost:     50,
// 	}
// 	// cookieList := CookieList()
// 	client.SetTransport(tr).SetHeader("Connection", "close").SetCookies(cookieList)
// 	resp, err := client.R().Get(url)
//
// 	if err != nil {
// 		return &query.Document{}
// 	}
//
// 	return DocQuery(resp.RawResponse)
// }
//
// func CookieList(cookie string) []*http.Cookie {
// 	// 字符串转json再转slice
//
// }
