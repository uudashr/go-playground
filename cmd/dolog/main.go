package main

import (
	"errors"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	// slog.Debug("Hello, debug!")
	// slog.Info("Hello, info!")
	// slog.Warn("Hello, warn!")
	// slog.Error("Hello, error!", errors.New("Opps"))

	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// handler := slog.NewTextHandler(os.Stderr)
	handler := opts.NewJSONHandler(os.Stderr)

	logger := slog.New(handler)
	logger.Debug("Hello, debug!")
	logger.Info("Hello, info!")
	logger.Warn("Hello, warn!")
	logger.Error("Hello, error!", errors.New("Opps"))

	logger.Info("finished", slog.Group("result", slog.String("status", "ok")))

}
