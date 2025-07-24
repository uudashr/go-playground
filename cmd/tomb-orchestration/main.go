package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/tomb.v2"
)

var errTerminated = errors.New("termination signal received")

func main() {
	logger := slog.Default()

	svc := &service{
		logger: logger,
	}

	if err := svc.run(); err != nil && err != errTerminated {
		logger.Error("Service failed", "error", err)
	}

	logger.Info("Service stopped")
}

type service struct {
	logger *slog.Logger
}

func (svc *service) run() error {
	t, ctx := tomb.WithContext(context.Background())

	t.Go(func() error {
		return svc.signalHandler(ctx)
	})

	t.Go(func() error {
		return svc.service1Handler(ctx)
	})

	t.Go(func() error {
		return svc.service2Handler(ctx)
	})

	return t.Wait()
}

func (svc *service) signalHandler(ctx context.Context) error {
	logger := svc.logger.With("component", "signal")

	ch := make(chan os.Signal, 1)
	defer close(ch)

	logger.InfoContext(ctx, "Listening termination signals")
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.InfoContext(ctx, "Receive done signal, stopping...", "error", ctx.Err())
		return ctx.Err()
	case sig := <-ch:
		logger.InfoContext(ctx, "Received termination signal", "signal", sig)
		return errTerminated
	}
}

func (svc *service) service1Handler(ctx context.Context) error {
	logger := svc.logger.With("component", "service-1")

	for {
		select {
		case <-ctx.Done():
			logger.InfoContext(ctx, "Receive done signal, stopping...", "error", ctx.Err())
			return ctx.Err()
		case <-time.After(1 * time.Second):
			logger.InfoContext(ctx, "Ticking")
		}
	}
}

func (svc *service) service2Handler(ctx context.Context) error {
	logger := svc.logger.With("component", "service-2")

	for {
		select {
		case <-ctx.Done():
			logger.InfoContext(ctx, "Receive done signal, stopping...", "error", ctx.Err())

			for i := 0; i < 5; i++ {
				logger.InfoContext(ctx, "Cleanup", "step", i+1)
				time.Sleep(200 * time.Millisecond)
			}
			logger.InfoContext(ctx, "Cleanup complete")

			return ctx.Err()
		case <-time.After(1 * time.Second):
			logger.InfoContext(ctx, "Ticking")
		}
	}
}
