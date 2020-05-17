package goecho_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func httpHandler() http.Handler {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://foo.com", "https://bar.com"},
	}))
	e.GET("/hello", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	return e
}

func TestCORS(t *testing.T) {
	testCases := map[string]struct {
		method     string
		path       string
		statusCode int
		origin     string
		retOrigin  string
		retVary    string
	}{
		"No Origin": {
			method:     http.MethodGet,
			path:       "/hello",
			statusCode: http.StatusOK,
		},
		"With Origin foo.com": {
			method:     http.MethodGet,
			path:       "/hello",
			statusCode: http.StatusOK,
			origin:     "https://foo.com",
			retOrigin:  "https://foo.com",
			retVary:    "Origin",
		},
		"With Origin bar.com": {
			method:     http.MethodGet,
			path:       "/hello",
			statusCode: http.StatusOK,
			origin:     "https://bar.com",
			retOrigin:  "https://bar.com",
			retVary:    "Origin",
		},
		"With Unallowed Origin": {
			method:     http.MethodGet,
			path:       "/hello",
			statusCode: http.StatusOK,
			origin:     "https://baz.com",
		},
	}

	h := httpHandler()

	for k, tc := range testCases {
		t.Run(k, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, tc.path, nil)
			if tc.origin != "" {
				req.Header.Add("Origin", tc.origin)
			}

			h.ServeHTTP(rec, req)

			if got, want := rec.Code, tc.statusCode; got != want {
				t.Errorf("Status code got: %d, want: %d", got, want)
				return
			}

			if got, want := rec.Header().Get("Access-Control-Allow-Origin"), tc.retOrigin; got != want {
				t.Errorf("Origin got: %q, want: %q", got, want)
			}

			if got, want := rec.Header().Get("Vary"), tc.retVary; got != want {
				t.Errorf("Vary got: %q, want: %q", got, want)
			}

			t.Log("Response header:", rec.Header())
		})
	}
}
