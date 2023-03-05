package main

import (
	"fmt"
	"hash/fnv"

	"github.com/rs/xid"
)

// ref: https://docs.launchdarkly.com/home/flags/rollouts
// flipt: https://github.com/flipt-io/flipt/blob/6ded9f8ee4bec234fe6a8960da4e534b249db5df/internal/server/evaluator.go
// flagd: https://github.com/open-feature/flagd/blob/34aca79e6ec9876a6cced0fe49e1ceea34d83696/pkg/eval/fractional_evaluation.go

func main() {
	// 1 2 3 4 5 6 7 8 9 10
	// 0 1 2 3 4 5 6 7 8 9
	// 50% = 5

	rolloutPercentage := 80
	rolloutRatio := float32(rolloutPercentage) / 100
	totalSegments := 100000
	segmentLimit := float32(totalSegments) * rolloutRatio

	fmt.Println("Total segments", totalSegments)
	fmt.Printf("Rollout percentage %d%%\n", rolloutPercentage)
	fmt.Println("Rollout ratio", rolloutRatio)
	fmt.Println("Segment limit", segmentLimit)
	fmt.Println()

	totalIDs := 500000
	totalIncluded := 0
	for i := 0; i < totalIDs; i++ {
		id := xid.New().String()
		h := fnv.New32a()
		h.Write([]byte(id))
		sum := h.Sum32()

		mod := sum % uint32(totalSegments)
		includeRollout := float32(mod) < segmentLimit
		if includeRollout {
			totalIncluded++
		}
		// fmt.Printf("%q\t%d\t%d\t%t\n", id, sum, mod, includeRollout)
	}
	fmt.Println()
	percentageIncluded := float32(totalIncluded) / float32(totalIDs) * 100
	fmt.Printf("Total included: %d (%f%%)\n", totalIncluded, percentageIncluded)
}

type User struct {
	ID    string
	Attrs map[string]string
}
