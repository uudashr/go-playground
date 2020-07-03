package gorillamux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func httpHandler(allowCreds bool, origins ...string) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// do nothing
	})

	var h http.Handler
	h = r
	opts := []handlers.CORSOption{
		handlers.AllowedOrigins(origins),
	}

	if allowCreds {
		opts = append(opts, handlers.AllowCredentials())
	}

	h = handlers.CORS(
		opts...,
	)(h)
	return h
}

func TestCORS(t *testing.T) {
	testCases := map[string]struct {
		allowedOrigins []string
		credsAllowed   bool
		method         string
		path           string
		origin         string
		statusCode     int
		retOrigin      string
		retVary        string
		retAllowCreds  string
	}{
		// --- Allowed Origins: "*"
		"No Origin sent, Allowed Origins *": {
			allowedOrigins: []string{"*"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
		},
		"With Origin baz.com sent, Allowed Origins *": {
			allowedOrigins: []string{"*"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://baz.com",
			retOrigin:      "*",
		},

		// --- Allowed Origins: "https://foo.com", "https://bar.com"
		"No Origin sent, Allowed Origins set": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
		},
		"With Origin foo.com sent, Allowed Origins set": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://foo.com",
			retOrigin:      "https://foo.com",
			retVary:        "Origin",
		},
		"With Origin bar.com sent, Allowed Origins set": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://bar.com",
			retOrigin:      "https://bar.com",
			retVary:        "Origin",
		},
		"With Unallowed Origin sent, Allowed Origins set": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://baz.com",
		},

		// Creds Allowed
		"With Creds Allowed, Origin foo.com sent, Allowed Origins set, return allow creds": {
			allowedOrigins: []string{"https://foo.com", "https://bar.com"},
			credsAllowed:   true,
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
			origin:         "https://foo.com",
			retOrigin:      "https://foo.com",
			retAllowCreds:  "true",
			retVary:        "Origin",
		},
		"With Creds Allowed, No Origin sent, Allowed Origins *, return no allow creds": {
			allowedOrigins: []string{"*"},
			credsAllowed:   true,
			method:         http.MethodGet,
			path:           "/hello",
			statusCode:     http.StatusOK,
		},
	}

	for k, tc := range testCases {
		t.Run(k, func(t *testing.T) {
			h := httpHandler(tc.credsAllowed, tc.allowedOrigins...)

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

			if got, want := rec.Header().Get("Access-Control-Allow-Credentials"), tc.retAllowCreds; got != want {
				t.Errorf("Allow-Credentials got: %q, want: %q", got, want)
			}

			if got, want := rec.Header().Get("Vary"), tc.retVary; got != want {
				t.Errorf("Vary got: %q, want: %q", got, want)
			}

			t.Log("Response header:", rec.Header())
		})
	}
}
