package gorillamux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func httpHandler(origins ...string) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// do nothing
	})

	var h http.Handler
	h = r
	h = handlers.CORS(
		handlers.AllowedOrigins(origins),
	)(h)
	return h
}

func TestCORS(t *testing.T) {
	testCases := map[string]struct {
		allowedOrigins []string
		method         string
		path           string
		statusCode     int
		origin         string
		retOrigin      string
		retVary        string
	}{
		"No Origin": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
		},
		"With Origin foo.com": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://foo.com",
			retOrigin:      "https://foo.com",
			retVary:        "Origin",
		},
		"With Origin bar.com": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://bar.com",
			retOrigin:      "https://bar.com",
			retVary:        "Origin",
		},
		"With Unallowed Origin": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://baz.com",
		},
	}

	for k, tc := range testCases {
		t.Run(k, func(t *testing.T) {
			h := httpHandler(tc.allowedOrigins...)

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
