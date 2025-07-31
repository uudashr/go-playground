package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	var (
		countDown           int
		maxConcurrentStream int
		secure              bool
		h2c                 bool
	)

	flag.IntVar(&countDown, "countdown", 0, "Countdown before shutdown")
	flag.BoolVar(&secure, "secure", false, "Enable TLS for the HTTP server (:8443 otherwise insecure at :8080)")
	flag.IntVar(&maxConcurrentStream, "max-concurrent-streams", 0, "Max concurrent streams for HTTP/2 server (0 for default)")
	flag.BoolVar(&h2c, "h2c", false, "Enable HTTP/2 support over cleartext (h2c)")
	flag.Parse()

	logger := newLogger()

	svc := &service{
		logger:              logger,
		countDown:           countDown,
		secure:              secure,
		maxConcurrentStream: uint32(maxConcurrentStream),
		h2c:                 h2c,
	}

	if err := svc.run(); err != nil && err != errTerminated {
		logger.Error("Service failed", "error", err)
		os.Exit(1)
	}

	logger.Info("Service stopped")
}

func newLogger() *slog.Logger {
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	handler = logConnID(handler)

	return slog.New(handler)
}
