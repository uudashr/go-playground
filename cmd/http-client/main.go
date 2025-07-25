package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func runMain() error {

	var (
		count      int
		baseURL    string
		serverName string
		verifyCert string
	)

	flag.IntVar(&count, "count", 1, "Number of requests to send")
	flag.StringVar(&baseURL, "base-url", "http://localhost:8080", "Base URL for the HTTP server")
	flag.StringVar(&serverName, "server-name", "localhost", "Server name for TLS verification (leave empty to skip)")
	flag.StringVar(&verifyCert, "verify-cert", "", "Path to CA certificate for TLS verification (leave empty to skip verification)")
	flag.Parse()

	client, err := newHTTPClient(clientConfig{cert: verifyCert, serverName: serverName})
	if err != nil {
		return fmt.Errorf("failed to create HTTP client: %w", err)
	}

	for i := range count {
		fmt.Printf("=== Sending request %d... === \n", i+1)
		if err := sendRequest(client, baseURL); err != nil {
			fmt.Printf("Error sending request %d: %v\n", i+1, err)
		}
	}

	return nil
}

func newHTTPClient(cfg clientConfig) (*http.Client, error) {
	if cfg.cert == "" {
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

func sendRequest(client *http.Client, baseURL string) error {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/hello", nil)
	if err != nil {
		return err
	}

	req.Close = false

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}

	defer res.Body.Close()

	fmt.Printf("Status: %d %s\n", res.StatusCode, res.Status)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}

	fmt.Println(string(b))

	return nil
}

type clientConfig struct {
	cert       string
	serverName string
}

func main() {

	if err := runMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
