package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type connectionIDKey struct{}

func requestLogAttrs(r *http.Request) []any {
	attrs := []any{
		"method", r.Method,
		"url", r.URL.String(),
	}

	if val := r.Header.Get("Connection"); val != "" {
		attrs = append(attrs, "connection", val)
	}

	return attrs
}

type httpHandler struct {
	logger   *slog.Logger
	mux      *http.ServeMux
	shutdown chan struct{}
}

func newHTTPHandler(logger *slog.Logger) *httpHandler {
	h := &httpHandler{
		logger:   logger,
		shutdown: make(chan struct{}),
	}
	h.initialize()

	return h
}

func (h *httpHandler) initialize() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shutdown", h.shutdownHandler)
	mux.HandleFunc("/", h.defaultHandler)
	h.mux = mux
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *httpHandler) shutdownHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := h.logger

	logger.InfoContext(ctx, "Shutdown request received")
	close(h.shutdown)
}

func (h *httpHandler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := h.logger

	logAttrs := requestLogAttrs(r)
	logger.InfoContext(ctx, "HTTP request received", logAttrs...)

	err := h.delayReq(r)
	if errors.Is(err, context.Canceled) {
		if err := context.Cause(ctx); errors.Is(err, context.Canceled) {
			// Client disconnection
			logger.WarnContext(ctx, "Delay canceled due to client disconnection", "error", err)
			return
		}

		logger.WarnContext(ctx, "Delay canceled due to termination", "error", err)
		http.Error(w, "Server is shutting down", http.StatusServiceUnavailable)
		return
	}

	if errors.Is(err, context.DeadlineExceeded) {
		logger.WarnContext(ctx, "Failed to delay due to exceeded deadline", "error", err, "cause", context.Cause(ctx))
		http.Error(w, "Request deadline exceeded", http.StatusGatewayTimeout)
		return
	}

	if err != nil {
		logger.ErrorContext(ctx, "Failed to handle request", "error", err)
		http.Error(w, "Failed to handle request", http.StatusInternalServerError)
		return
	}

	if connID := connIDFromContext(ctx); connID != "" {
		w.Header().Set("X-Connection-ID", connID)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

func (h *httpHandler) delayReq(r *http.Request) error {
	ctx := r.Context()

	if val := r.URL.Query().Get("delay"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			return &invalidDelayError{val: val, cause: err}
		}

		h.delay(ctx, d)
	}

	return nil
}

func (h *httpHandler) delay(ctx context.Context, d time.Duration) error {
	logger := h.logger

	logger.InfoContext(ctx, "Delaying response", "duration", d)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		logger.InfoContext(ctx, "Delay completed, sending response", "duration", d)
	}

	return nil
}

type invalidDelayError struct {
	val   string
	cause error
}

func (e *invalidDelayError) Error() string {
	return fmt.Sprintf("invalid delay parameter %q", e.val)
}

func (e *invalidDelayError) Unwrap() error {
	return e.cause
}

func contextWithConnID(ctx context.Context, connID string) context.Context {
	return context.WithValue(ctx, connectionIDKey{}, connID)
}

func connIDFromContext(ctx context.Context) string {
	if val, ok := ctx.Value(connectionIDKey{}).(string); ok {
		return val
	}

	return ""
}

type connIDLogHandler struct {
	slog.Handler
}

func logConnID(next slog.Handler) slog.Handler {
	return &connIDLogHandler{Handler: next}
}

func (h *connIDLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if connID := connIDFromContext(ctx); connID != "" {
		r.Add(slog.String("connID", connID))
	}

	return h.Handler.Handle(ctx, r)
}

func (h *connIDLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return logConnID(h.Handler.WithAttrs(attrs))
}

func (h *connIDLogHandler) WithGroup(name string) slog.Handler {
	return logConnID(h.Handler.WithGroup(name))
}
