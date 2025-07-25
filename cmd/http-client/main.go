package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

func runMain() error {

	var (
		count      int
		baseURL    string
		serverName string
		cert       string
		delay      time.Duration
		concurrent bool
		h2c        bool
	)

	flag.IntVar(&count, "count", 1, "Number of requests to send")
	flag.StringVar(&baseURL, "base-url", "http://localhost:8080", "Base URL for the HTTP server")
	flag.StringVar(&serverName, "server-name", "localhost", "Server name for TLS verification (leave empty to skip)")
	flag.StringVar(&cert, "cert", "", "Path to CA certificate for TLS verification (leave empty to skip verification)")
	flag.BoolVar(&concurrent, "concurrent", false, "Send requests concurrently")
	flag.DurationVar(&delay, "delay", 0, "Delaying response (e.g., 1s, 500ms)")
	flag.BoolVar(&h2c, "h2c", false, "Enable HTTP/2 support over cleartext (h2c)")

	flag.Parse()

	cfg := clientConfig{
		cert:       cert,
		serverName: serverName,
		h2c:        h2c,
	}

	client, err := newHTTPClient(cfg)
	if err != nil {
		return fmt.Errorf("failed to create HTTP client: %w", err)
	}

	ha := &httpAgent{
		baseURL: baseURL,
		client:  client,
		delay:   delay,
	}

	if concurrent {
		return sendConcurrently(ha, count)
	}

	return sendSequentially(ha, count)
}

func sendConcurrently(ha *httpAgent, count int) error {
	var wg sync.WaitGroup
	wg.Add(count)

	for i := range count {
		go func(num int) {
			defer wg.Done()
			if err := ha.sendRequest(i); err != nil {
				fmt.Printf("Error sending request req-%d: %v\n", num, err)
			}
		}(i + 1)
	}

	wg.Wait()

	return nil
}

func sendSequentially(ha *httpAgent, count int) error {
	for i := range count {
		iter := i + 1
		if err := ha.sendRequest(i); err != nil {
			fmt.Printf("Error sending request req-%0d: %v\n", iter, err)
		}
	}

	return nil
}

func newPriorKnowledgeClient() *http.Client {
	transport := &http2.Transport{
		// Allow HTTP/2 over plaintext connections
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			// Return a regular TCP connection instead of TLS
			return nil, fmt.Errorf("use DialTLSContext")
		},
		DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
			// This is a bit of a hack - we're using the TLS dial but without actual TLS
			d := &net.Dialer{}
			return d.DialContext(ctx, network, addr)
		},
	}

	return &http.Client{
		Transport: transport,
	}
}

func newHTTPClient(cfg clientConfig) (*http.Client, error) {
	if cfg.cert == "" {
		// Insecure mode, skip TLS verification

		if cfg.h2c {
			// Enable HTTP/2 support over cleartext
			return newPriorKnowledgeClient(), nil
		}

		return &http.Client{}, nil
	}

	// Load certificate
	// PEM file (Privacy-Enhanced Mail) is a base64 encoded certificate
	certPEM, err := os.ReadFile(cfg.cert)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}

	// Add pem to root CA pool
	// x509 is a standard library package for handling X.509 certificates
	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(certPEM); !ok {
		return nil, fmt.Errorf("failed to add certificate %s to pool", cfg.cert)
	}

	tlsConfig := &tls.Config{
		RootCAs:    rootCAs,
		ServerName: cfg.serverName,
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}, nil
}

type httpAgent struct {
	baseURL string
	client  *http.Client
	delay   time.Duration
}

func (ha *httpAgent) sendRequest(i int) error {
	num := i + 1

	req, err := http.NewRequest(http.MethodGet, ha.baseURL+"/hello", nil)
	if err != nil {
		return err
	}

	if ha.delay > 0 {
		d := time.Duration(int64(ha.delay) / int64(num))
		assignDelay(req, d)
	}

	fmt.Printf("==> Send req-%03d %s\n", num, req.URL.String())
	res, err := ha.client.Do(req)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}

	defer res.Body.Close()

	fmt.Printf("==== Response req-%03d ==============\n", num)
	fmt.Printf("Status -> %d %s\n", res.StatusCode, res.Status)
	fmt.Printf("Proto -> %s\n", res.Proto)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}

	fmt.Printf("---- Body req-%03d ------------------\n", num)
	fmt.Println(string(b))
	fmt.Println("____________________________________")

	return nil
}

func assignDelay(r *http.Request, d time.Duration) {
	vals := r.URL.Query()
	vals.Set("delay", d.String())
	r.URL.RawQuery = vals.Encode()
}

type clientConfig struct {
	// The cert used for TLS verification CA certificate. Non empty value means
	// TLS verification is enabled.
	cert       string
	serverName string

	// h2c enable HTTP/2 support over cleartext (h2c)
	h2c bool
}

func main() {
	if err := runMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
