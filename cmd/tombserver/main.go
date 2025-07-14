package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/tomb.v2"
)

var ErrTerminanated = errors.New("termination signal received")

func main() {
	logger := slog.Default()

	svc := &service{
		logger: logger,
	}

	if err := svc.run(); err != ErrTerminanated {
		logger.Error("Application failed", "error", err)
		os.Exit(1)
	}

	logger.Info("Application stopped")
}

type service struct {
	logger *slog.Logger
}

func (s *service) run() error {
	t, ctx := tomb.WithContext(context.Background())

	t.Go(func() error {
		return s.signalHandler(ctx)
	})

	t.Go(func() error {
		return s.httpServer(ctx)
	})

	return t.Wait()
}

func (s *service) signalHandler(ctx context.Context) error {
	logger := s.logger
	logger = logger.With("component", "signal")

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	logger.InfoContext(ctx, "Waiting for termination signal")
	select {
	case <-ch:
		logger.InfoContext(ctx, "Received termination signal")
		return ErrTerminanated
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *service) httpServer(ctx context.Context) error {
	logger := s.logger
	logger = logger.With("component", "http")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		latency := r.Header.Get("X-Latency")
		if latency != "" {
			duration, err := time.ParseDuration(latency)
			if err != nil {
				logger.ErrorContext(ctx, "Invalid latency header", "error", err)
				http.Error(w, "Invalid latency header", http.StatusBadRequest)
				return
			}

			logger.InfoContext(ctx, "Simulating latency", "duration", duration)
			time.Sleep(duration)
			logger.InfoContext(ctx, "Latency simulation complete")
		}

		w.Write([]byte("Hello, World!"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	errCh := make(chan error, 1)
	go func() {
		defer close(errCh)

		logger.InfoContext(ctx, "Starting server", "address", server.Addr)
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			errCh <- err
		}

		logger.InfoContext(ctx, "Server stopped")
	}()

	select {
	case err := <-errCh:
		logger.ErrorContext(ctx, "Server failed", "error", err)
		return err
	case <-ctx.Done():
		logger.InfoContext(ctx, "Shutting down server")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.ErrorContext(ctx, "Server shutdown failed", "error", err)
			return err
		}

		return ctx.Err()
	}
}
