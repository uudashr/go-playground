package main

import (
	"context"
	"log/slog"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// ratelimit-remaining: 2999
	// ratelimit-reset: 5

	logger := slog.Default()
	l := rate.NewLimiter(rate.Every(5*time.Second), 3000)

	for i := range 3001 {
		iter := i + 1
		logger.Info("Iteration", "iter", iter)
		l.Wait(context.Background())
		logger.Info("Done waiting", "iter", iter)
	}
}
