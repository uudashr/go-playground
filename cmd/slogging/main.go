package main

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog"
)

func main() {
	runZap()
	runZeroLog()
}

func runZap() {
	zapL := zap.Must(zap.NewProduction())

	defer zapL.Sync()

	logger := slog.New(zapslog.NewHandler(zapL.Core()))

	logger.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.String("path", "/api/user"),
		slog.Int("status", 200),
	)
}

func runZeroLog() {
	zerologL := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	logger := slog.New(
		slogzerolog.Option{Logger: &zerologL}.NewZerologHandler(),
	)

	logger.Info(
		"incoming request",
		slog.String("method", "GET"),
		slog.String("path", "/api/user"),
		slog.Int("status", 200),
	)
}
