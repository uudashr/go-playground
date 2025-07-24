package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func runMain() error {

	for i := range 2 {
		fmt.Printf("=== Sending request %d... === \n", i+1)
		if err := sendRequest(); err != nil {
			fmt.Printf("Error sending request %d: %v\n", i+1, err)
		}
	}

	return nil
}

func sendRequest() error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	if err != nil {
		return err
	}

	req.Close = false

	client := &http.Client{}
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

func main() {

	if err := runMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
