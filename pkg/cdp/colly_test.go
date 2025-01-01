package glidedsky

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"testing"
)

func TestColly(t *testing.T) {
	mUrl := "http://www.ifeng.com/"
	//colly的主体是Collector对象，管理网络通信和负责在作业运行时执行附加的回掉函数
	c := colly.NewCollector(
		// 开启本机debug
		colly.Debugger(&debug.LogDebugger{}),
	)
	//发送请求之前的执行函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("这里是发送之前执行的函数")
	})
	//发送请求错误被回调
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Print(err)
	})

	//响应请求之后被回调
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response body length：", len(r.Body))
	})
	//response之后会调用该函数，分析页面数据
	c.OnHTML("div#newsList h1 a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	//在OnHTML之后被调用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	//这里是执行访问url
	c.Visit(mUrl)
}
