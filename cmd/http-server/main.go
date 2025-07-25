package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	var (
		countDown int
		tls       bool
	)

	flag.IntVar(&countDown, "countdown", 0, "Countdown before shutdown")
	flag.BoolVar(&tls, "tls", false, "Enable TLS for the HTTP server")
	flag.Parse()

	logger := newLogger()
	svc := &service{
		logger:    logger,
		countDown: countDown,
		tls:       tls,
	}

	if err := svc.run(); err != nil && err != errTerminated {
		logger.Error("Service failed", "error", err)
		os.Exit(1)
	}

	logger.Info("Service stopped")
}

func newLogger() *slog.Logger {
	var logHandler slog.Handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logHandler = logConnID(logHandler)

	return slog.New(logHandler)
}
