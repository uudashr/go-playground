package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func runMain() error {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello?delay=10s", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer res.Body.Close()

	fmt.Printf("%s %s\n", res.Status, res.Proto)
	return nil
}

func runRetryable() error {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// 1. Possible client.Transport is nil, the it will use http.DefaultTransport

	rc := retryablehttp.NewClient()
	rc.RetryMax = 5
	client.Transport = rc.StandardClient().Transport
	rc.Logger = nil

	// 2. Possible it is a *retryablehttp.RoundTripper

	fmt.Printf("%T\n", client.Transport)
	rt := client.Transport.(*retryablehttp.RoundTripper)
	fmt.Printf("%T\n", rt.Client.HTTPClient.Transport)

	transport := rt.Client.HTTPClient.Transport.(*http.Transport)
	transport.HTTP2 = &http.HTTP2Config{
		MaxConcurrentStreams: 100,
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello?delay=10s", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer res.Body.Close()

	fmt.Printf("%s %s\n", res.Status, res.Proto)
	return nil

}

func main() {
	if err := runRetryable(); err != nil {
		fmt.Fprintf(os.Stderr, "Error in main: %v\n", err)
		os.Exit(1)
	}
}
