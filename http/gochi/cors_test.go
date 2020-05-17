package gochi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi"
)

func httpHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://foo.com", "https://bar.com"},
	}))
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		// do nothing
	})
	return r
}

func TestCORS(t *testing.T) {
	testCases := map[string]struct {
		method     string
		path       string
		statusCode int
		origin     string
		retOrigin  string
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
		},
		"With Origin bar.com": {
			method:     http.MethodGet,
			path:       "/hello",
			statusCode: http.StatusOK,
			origin:     "https://bar.com",
			retOrigin:  "https://bar.com",
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

			t.Log("Response header:", rec.Header())
		})
	}
}
