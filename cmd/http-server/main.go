package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := newLogger()
	svc := &service{
		logger: logger,
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
