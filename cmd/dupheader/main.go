package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	req, err := http.NewRequest("POST", "https://example.com", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Type", "application/json")

	out, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
		return
	}

	fmt.Println(string(out))
}
