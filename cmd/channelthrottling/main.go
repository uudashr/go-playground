package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan string)

	// Producer
	go func() {
		for i := 0; i < 30; i++ {
			ch <- fmt.Sprintf("msg=message-%d timestamp=%s", i, time.Now().Format(time.RFC3339))
			if i == 2 {
				time.Sleep(5 * time.Second)
			}
		}
		close(ch)
	}()

	const rateLimit = 5
	const duration = 1 * time.Second

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	messageCount := 0

	// Consumer
	for msg := range ch {
		log.Println("Got message:", msg)
		messageCount++
		if messageCount >= rateLimit {
			log.Printf("Reached rate limit, waiting for %s\n", duration)
			t := <-ticker.C
			messageCount = 0
			log.Printf("Message counter reset, t=%s\n", t.Format(time.RFC3339))
		}
	}

}
