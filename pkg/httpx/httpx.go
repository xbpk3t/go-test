package httpx

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func x() {
	var HTTPTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时时间
			KeepAlive: 60 * time.Second, // 保持长连接的时间
		}).DialContext, // 设置连接的参数
		MaxIdleConns:          500,              // 最大空闲连接
		IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
		ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
		MaxIdleConnsPerHost:   100,              // 每个host保持的空闲连接数
	}

	start2 := time.Now()
	client2 := http.Client{Transport: HTTPTransport} // 初始化一个带有transport的http的client

	fmt.Println(start2.String())
	fmt.Println(client2)
}
