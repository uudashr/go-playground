package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moby/moby/pkg/namesgenerator"
	"golang.org/x/sync/errgroup"
)

var errTerminated = errors.New("termination signal received")

type service struct {
	logger    *slog.Logger
	countDown int
}

func (svc *service) run() error {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return svc.signalListener(ctx)
	})

	g.Go(func() error {
		return svc.httpServer(ctx)
	})

	if svc.countDown > 0 {
		g.Go(func() error {
			return svc.countdownHandler(ctx, svc.countDown)
		})
	}

	return g.Wait()
}

func (svc *service) countdownHandler(ctx context.Context, count int) error {
	logger := svc.logger.With("component", "countdown")

	logger.InfoContext(ctx, "Starting countdown", "count", count)
	for i := count; i >= 0; i-- {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(1 * time.Second):
			logger.InfoContext(ctx, "Counting down", "step", i)
		}
	}

	return errors.New("explode error")
}

func (svc *service) signalListener(ctx context.Context) error {
	logger := svc.logger.With("component", "signal-listener")

	ch := make(chan os.Signal, 1)
	defer close(ch)

	logger.InfoContext(ctx, "Listening termination signals")
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.InfoContext(ctx, "Receive done signal, stopping...", "error", ctx.Err(), "cause", context.Cause(ctx))
	case sig := <-ch:
		logger.InfoContext(ctx, "Received termination signal", "signal", sig)
		return errTerminated
	}

	return nil
}

func (svc *service) httpServer(ctx context.Context) error {
	logger := svc.logger.With("component", "http-server")

	h := newHTTPHandler(logger)

	svr := &http.Server{
		Addr:    ":8080",
		Handler: h,
		BaseContext: func(net.Listener) context.Context {
			logger.DebugContext(ctx, "HTTP server base context created")
			return ctx
		},
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			name := namesgenerator.GetRandomName(0)

			ctx = contextWithConnID(ctx, name)
			logger.DebugContext(ctx, "HTTP server connection context created", "localAddr", c.LocalAddr())
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		logger.InfoContext(ctx, "Receive done signal, shutting down HTTP server...", "error", ctx.Err(), "cause", context.Cause(ctx))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := svr.Shutdown(ctx); err != nil {
			logger.WarnContext(ctx, "HTTP server shutdown failed", "error", err)
		}
	}()

	logger.InfoContext(ctx, "Starting HTTP server", "addr", svr.Addr)
	if err := svr.ListenAndServe(); err != http.ErrServerClosed {
		logger.WarnContext(ctx, "HTTP server fail", "error", err)
		return err
	}

	logger.InfoContext(ctx, "HTTP server stopped")

	return nil
}
