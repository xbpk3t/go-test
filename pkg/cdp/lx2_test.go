package glidedsky

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gs/util/gq"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

// @url http://www.glidedsky.com/level/web/crawler-basic-2
func TestLx2(t *testing.T) {

	sum := 0
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(i, sum int) {
			defer wg.Done()
			page := fmt.Sprintf("http://www.glidedsky.com/level/web/crawler-basic-2?page=%d", i)
			doc := gq.GetWithCookie(page, ReadCookie())
			doc.Find(".col-md-1").Each(func(i int, selection *goquery.Selection) {
				num, _ := strconv.Atoi(strings.TrimSpace(selection.Text()))
				sum += num
			})
			time.Sleep(time.Second)
		}(i, sum)
	}

	wg.Wait()
	fmt.Println(sum)
}
