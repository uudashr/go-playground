package main

import (
	"fmt"
	"time"
)

func main() {
	expiresAt := time.Now().Add(2 * time.Hour).Add(30 * time.Minute).Add(10 * time.Second)

	// Print Unix timestamp in milliseconds in UTC.
	fmt.Println(expiresAt.UnixMilli())
}
