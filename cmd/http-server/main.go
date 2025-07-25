package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	var (
		countDown int
	)

	flag.IntVar(&countDown, "countdown", 0, "Countdown before shutdown.")
	flag.Parse()

	logger := newLogger()
	svc := &service{
		logger:    logger,
		countDown: countDown,
	}

	if err := svc.run(); err != nil && err != errTerminated {
		logger.Error("Service failed", "error", err)
		os.Exit(1)
	}

	logger.Info("Service stopped")
}

func newLogger() *slog.Logger {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	return slog.New(logHandler)
}
