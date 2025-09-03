package main

import (
	"errors"
	"log/slog"
	"net"
	"os"
)

func main() {
	// slog.New(slog.NewJSONHandler())

	// 除了第一个是msg主体外，后面的kv一一对应即可
	slog.Info("Execute Success", "url", "http://localhost:8080")
	// slog.Error()

	// kv
	// TextHandler和JSONHandler
	h := slog.NewTextHandler(os.Stderr, nil)
	ll := slog.New(h)
	ll.Info("greeting", "name", "xxx")

	h1 := slog.NewJSONHandler(os.Stderr, nil)
	l1 := slog.New(h1)
	l1.Info("greeting", "name", "xxx")

	// 用来做kv复用，通常存context、将程序的进程ID和用于编译的 Go 版本之类的
	l := slog.With("attr1", "attr1_val", "attr2", "attr2_val")
	l.Error("conn server error", "err", net.ErrClosed, "status", 500)
	l.Error("close server error", "err", net.ErrClosed, "status", 501)

	// slog.HandlerOptions
	// 根据应用环境轻松切换足够的处理程序也很容易。例如，您可能更喜欢在开发日志中使用 TextHandler ，因为它更容易阅读，然后在生产环境中切换到 JSONHandler ，以获得更灵活性和与各种日志工具的兼容性。

	// Group形式的日志输出
	slog.Info("group logging", "url", "http://localhost:8080", slog.Group("properties",
		slog.Int("width", 4000),
		slog.Int("height", 3000),
		slog.String("format", "jpeg"),
	))

	// 搭配使用 JSONHandler + Group 更好
	l1.Info("group logging", "url", "http://localhost:8080", slog.Group("properties",
		slog.Int("width", 4000),
		slog.Int("height", 3000),
		slog.String("format", "jpeg"),
	))

	// 动态调整日志级别 slog.AtomicLevel

	// 使用 Slog 进行错误日志记录
	// xerrors
	slog.Error("error msg", slog.Any("error", errors.New("xxx")))

	// 使用 LogValuer 接口隐藏敏感字段

	// User 实现 LogValuer 接口
	// func (u *User) LogValue() slog.Value {
	//    return slog.GroupValue(
	//        slog.String("id", u.ID),
	//        slog.String("name", u.FirstName+" "+u.LastName),
	//    )
	// }

	// logger.Info("info", "user", u)
	// 这样只会在log中有id和name，而不会添加其他如mail之类的字段

	// 13. 编写和存储 Go 日志的最佳实践
	// 13.1. 1. 标准化您的日志接口
	// 13.2. 2. 在错误日志中添加堆栈跟踪
	// 13.3. 3. 对您的 Slog 语句进行检查，以确保一致性
	// 13.4. 4. 集中管理日志，但首先将它们持久化到本地文件
	// 13.5. 5. 采样你的日志
	// 13.6. 6. 使用日志管理服务

	// slog性能比zap还好，
}
