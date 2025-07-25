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
		http2     bool
	)

	flag.IntVar(&countDown, "countdown", 0, "Countdown before shutdown")
	flag.BoolVar(&tls, "tls", false, "Enable TLS for the HTTP server (:8443 otherwise :8080)")
	flag.BoolVar(&http2, "http2", false, "Enable HTTP/2 support")
	flag.Parse()

	logger := newLogger()

	svc := &service{
		logger:    logger,
		countDown: countDown,
		tls:       tls,
		http2:     http2,
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
