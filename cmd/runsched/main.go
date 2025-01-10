package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	slog.Info("Run scheduler", "time", time.Now())

	fn := func() {
		slog.Info("Run", "time", time.Now())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if err := runTickerJob(ctx, fn); err != nil {
		panic(err)
	}

	fmt.Println("Done")
}

func runCronJob(ctx context.Context, fn func()) error {
	jobdef := gocron.CronJob("* * * * *", true)

	sched, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		sched.Shutdown()
	}()

	sched.NewJob(jobdef, gocron.NewTask(fn))

	sched.Start()

	return nil
}

func runTickerJob(ctx context.Context, fn func()) error {
	now := time.Now()
	nextMinute := now.Truncate(time.Minute).Add(1 * time.Minute)
	initialDelay := nextMinute.Sub(now)

	time.Sleep(initialDelay)

	ticker := time.NewTicker(1 * time.Minute)

	defer ticker.Stop()

	fn()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			fn()
		}
	}
}
