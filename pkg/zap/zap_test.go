package zap

import (
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// BenchmarkZap 测试 zap 日志库的性能
func BenchmarkZap(b *testing.B) {
	// 创建一个丢弃所有日志输出的 zap logger
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	))
	defer logger.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test zap",
			zap.String("msg", "success"),
			zap.Int("count", i),
			zap.Time("timestamp", time.Now()),
			zap.Any("test", []string{"why", "not", "use", "more", "fields"}))
	}
}

// BenchmarkSlog 测试 slog 日志库的性能
func BenchmarkSlog(b *testing.B) {
	// 创建一个丢弃所有日志输出的 slog logger
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test slog",
			slog.String("msg", "success"),
			slog.Int("count", i),
			slog.Time("timestamp", time.Now()),
			slog.Any("test", []string{"why", "not", "use", "more", "fields"}))
	}
}

// BenchmarkZero 测试 zerolog 日志库的性能
func BenchmarkZero(b *testing.B) {
	// 创建一个丢弃所有日志输出的 zerolog logger
	logger := zerolog.New(io.Discard).Level(zerolog.InfoLevel)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("msg", "success").
			Int("count", i).
			Time("timestamp", time.Now()).
			Any("test", []string{"why", "not", "use", "more", "fields"}).
			Msg("test zerolog")
	}
}

// BenchmarkZapDevelopment 测试 zap 开发模式下的性能
func BenchmarkZapDevelopment(b *testing.B) {
	// 创建一个丢弃所有日志输出的 zap development logger
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(io.Discard),
		zapcore.DebugLevel,
	))
	defer logger.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test zap development",
			zap.String("msg", "success"),
			zap.Int("count", i),
			zap.Time("timestamp", time.Now()),
			zap.Any("test", []string{"why", "not", "use", "more", "fields"}))
	}
}

// BenchmarkZapSugar 测试 zap sugared logger 的性能
func BenchmarkZapSugar(b *testing.B) {
	// 创建一个丢弃所有日志输出的 zap sugared logger
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	))
	sugar := logger.Sugar()
	defer sugar.Sync()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar.Infow("test zap sugar",
			"msg", "success",
			"count", i,
			"timestamp", time.Now(),
			"test", []string{"why", "not", "use", "more", "fields"})
	}
}

// BenchmarkSlogJSON 测试 slog 使用 JSON 处理器的性能
func BenchmarkSlogJSON(b *testing.B) {
	// 修复空指针异常：使用 io.Discard 而不是 nil
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test slog json",
			slog.String("msg", "success"),
			slog.Int("count", i),
			slog.Time("timestamp", time.Now()),
			slog.Any("test", []string{"why", "not", "use", "more", "fields"}))
	}
}

// BenchmarkZeroDisabled 测试 zerolog 在禁用日志级别时的性能
func BenchmarkZeroDisabled(b *testing.B) {
	// 创建一个丢弃所有日志输出的 zerolog logger
	logger := zerolog.New(io.Discard).Level(zerolog.Disabled)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().
			Str("msg", "success").
			Int("count", i).
			Time("timestamp", time.Now()).
			Any("test", []string{"why", "not", "use", "more", "fields"}).
			Msg("test zerolog disabled")
	}
}
