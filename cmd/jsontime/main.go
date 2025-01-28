package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Activity struct {
	Name      string    `json:"name"`
	OccuredAt time.Time `json:"occured_at"`
}

func encode() {
	act := Activity{
		Name:      "swimming",
		OccuredAt: time.Now(),
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ") // pretty print

	if err := enc.Encode(act); err != nil {
		panic(err)
	}
}

func main() {
	// Create a time in a specific timezone
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Date(2025, 1, 24, 15, 0, 0, 0, loc) // 3 PM in Tokyo timezone

	// Serialize to JSON
	jsonData, _ := json.Marshal(t)
	fmt.Println("Serialized JSON:", string(jsonData))

	// Deserialize back
	var parsedTime time.Time
	_ = json.Unmarshal(jsonData, &parsedTime)
	fmt.Println("Parsed Time:", parsedTime)
}
