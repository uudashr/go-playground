package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	t := <-ticker.C
	log.Println("Tick", t.Format(time.RFC3339))

	t = <-ticker.C
	log.Println("Tick", t.Format(time.RFC3339))

	time.Sleep(3 * time.Second)
	t = <-ticker.C
	log.Println("Tick", t.Format(time.RFC3339))

	t = <-ticker.C
	log.Println("Tick", t.Format(time.RFC3339))
}
