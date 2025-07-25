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
		h2c       bool
	)

	flag.IntVar(&countDown, "countdown", 0, "Countdown before shutdown")
	flag.BoolVar(&tls, "tls", false, "Enable TLS for the HTTP server (:8443 otherwise :8080)")
	flag.BoolVar(&h2c, "h2c", false, "Enable HTTP/2 support over cleartext (h2c)")
	flag.Parse()

	logger := newLogger()

	svc := &service{
		logger:    logger,
		countDown: countDown,
		tls:       tls,
		h2c:       h2c,
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
