package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// type Result struct {
// 	Error    error
// 	Response *http.Response
// }
//
// // 通过 chan 调用接口
// func main() {
// 	done := make(chan any)
// 	defer close(done)
// 	urls := []string{"https://trello.com", "https://badhost"}
//
// 	for result := range checkStatus(done, urls...) {
// 		if result.Error != nil {
// 			fmt.Printf("error: %v\n", result.Error)
// 			continue
// 		}
// 		fmt.Printf("response: %v\n", result.Response)
// 	}
// }
//
// func checkStatus(done <-chan any, urls ...string) <-chan Result {
// 	results := make(chan Result)
// 	go func() {
// 		defer close(results)
// 		for _, url := range urls {
// 			var result Result
// 			resp, err := http.Get(url)
// 			result = Result{Error: err, Response: resp}
// 			select {
// 			case <-done:
// 				return
// 			case results <- result:
// 			}
// 		}
// 	}()
//
// 	return results
// }

// barrier 模式的使用
// 集中获取所有网站的状态码
func main() {
	barrier([]string{"https://www.baidu.com", "http://www.sina.com", "https://segmentfault.com/"}...)
}

// 合并结果
func barrier(endPoints ...string) {
	requestNumber := len(endPoints)
	in := make(chan barrierResp, requestNumber)
	response := make([]barrierResp, requestNumber)

	defer close(in)

	for _, point := range endPoints {
		go makeRequest(in, point)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Status)
		}
	}
}

type barrierResp struct {
	Err    error
	Resp   string
	Status int
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}

	client := http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}
