package tatanas

import (
	"log/slog"
	"math"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/katana/pkg/engine/standard"
	"github.com/projectdiscovery/katana/pkg/output"
	"github.com/projectdiscovery/katana/pkg/types"
)

func main() {
	options := &types.Options{
		MaxDepth:     3,             // 最大爬取深度
		FieldScope:   "rdn",         // 爬取范围字段
		BodyReadSize: math.MaxInt,   // 最大响应大小
		Timeout:      10,            // 请求超时时间（秒）
		Concurrency:  10,            // 并发数
		Parallelism:  10,            // URL处理并行数
		Delay:        0,             // 每次爬取请求的延迟时间（秒）
		RateLimit:    150,           // 每秒最大请求数
		Strategy:     "depth-first", // 爬取策略（深度优先或广度优先）
		OnResult: func(result output.Result) { // 处理结果的回调函数
			gologger.Info().Msg(result.Request.URL)
		},
	}
	crawlerOptions, err := types.NewCrawlerOptions(options)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	defer crawlerOptions.Close()
	crawler, err := standard.New(crawlerOptions)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	defer crawler.Close()
	var input = "https://www.hackerone.com"
	err = crawler.Crawl(input)
	if err != nil {
		gologger.Warning().Msgf("Could not crawl %s: %s", input, err.Error())
	}

	slog.Error("")
}
