package chromePwdChecker

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/modood/table"
)

type ChromePwdItem struct {
	Name, URL, Username, Password string
	IsAccess                      bool
	wg                            sync.WaitGroup
}

func NewChromePwdChecker() *ChromePwdItem {
	return &ChromePwdItem{}
}

// 读取csv文件
func (*ChromePwdItem) ReadCSV(filename string) []*ChromePwdItem {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	reader := csv.NewReader(bufio.NewReader(file))
	var items []*ChromePwdItem
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errors.New("parse error")
		}
		items = append(items, &ChromePwdItem{
			Name:     line[0],
			URL:      line[1],
			Username: line[2],
			Password: line[3],
		})
	}

	return items
}

func (cpi *ChromePwdItem) Checker(items []*ChromePwdItem) {
	for _, item := range items {
		cpi.wg.Add(1)
		go func(item *ChromePwdItem) {
			isAccess := cpi.CheckIsAccess(item.URL)
			item.IsAccess = isAccess
			defer cpi.wg.Done()
		}(item)
	}
	cpi.wg.Wait()

	j := 0
	// 移除所有access=true的数据
	for _, item := range items {
		if !item.IsAccess {
			items[j] = item
			j++
		}
	}

	table.Output(items[:j])
}

// CheckIsAccess 查看网站是否能够打开.
func (cpi *ChromePwdItem) CheckIsAccess(url string) bool {
	client := resty.New()

	response, err := client.R().Get(url)
	if err != nil {
		return false
	}
	return response.IsSuccess()
}
