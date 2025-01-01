package gq

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	jsoniter "github.com/json-iterator/go"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

type Movie struct {
	MovieName    string
	Score        float64  // 电影评分
	Categories   []string // 电影标签
	Pic          string   // 电影海报
	Place        string   // 首映地
	PublishTime  string   // 电影上映时间
	PlayDuration int      // 片长，单位分钟
}

const (
	BaseURL = "https://ssr1.scrape.center/page/%d"
)

// @url https://ssr1.scrape.center/
func TestLx1(t *testing.T) {
	var wg sync.WaitGroup

	ret := make([]Movie, 0)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			page := fmt.Sprintf(BaseURL, i)
			doc := gq.FetchHTML(page)
			doc.Find(".el-card").Each(func(i int, sel *goquery.Selection) {

				categories := []string{}
				sel.Find(".category").Each(func(i int, s *goquery.Selection) {
					categories = append(categories, s.Find("span").Text())
				})
				// 电影海报
				cover, isCoverExist := sel.Find(".cover").Attr("src")
				if !isCoverExist {
					cover = ""
				}
				// 片长
				duration, err := strconv.Atoi(sel.Find(".m-v-sm").First().Find("span").Eq(2).Text())
				if err != nil {
					duration = 120
				}
				// 发布时间
				pt := sel.Find(".m-v-sm").Eq(1).Find("span").Text()
				pubTime := ""
				if pt != "" {
					pubTime = pt[0 : len(pt)-7]
				}

				ret = append(ret, Movie{
					MovieName:    sel.Find(".m-b-sm").Text(),
					Categories:   categories,
					Score:        gconv.Float64(strings.TrimSpace(sel.Find(".m-b-n-sm").Text())),
					Pic:          cover,
					Place:        sel.Find(".m-v-sm").First().Find("span").First().Text(),
					PlayDuration: duration,
					PublishTime:  pubTime,
				})
			})
		}(i)

		wg.Wait()
	}

	// slice转json
	mb, err := jsoniter.Marshal(ret)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("./result/lx1.json", mb, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestStrip(t *testing.T) {
	pt := "2002-01-26 上映"
	pt = pt[0 : len(pt)-7]
	p, err := gtime.StrToTimeFormat(pt, "Y-m-d")
	if err != nil {
		return
	}
	fmt.Println(p)
}
