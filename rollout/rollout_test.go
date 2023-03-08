package rollout_test

import (
	"fmt"
	"testing"

	"github.com/rs/xid"
	"github.com/uudashr/go-playground/rollout"
)

func ExampleShouldRollout() {
	id := "0998ce83-551b-4091-bcc2-c50c19c08c9c"

	ok, _ := rollout.ShouldRollout(id, 0.1)
	fmt.Println("ShouldRollout:", ok)
}

func TestShouldRollout(t *testing.T) {
	testCases := map[string]struct {
		id    string
		ratio float32
		ok    bool
	}{
		"always ok": {
			id:    "0998ce83-551b-4091-bcc2-c50c19c08c9c",
			ratio: 0,
			ok:    false,
		},
		"always nok": {
			id:    "0998ce83-551b-4091-bcc2-c50c19c08c9c",
			ratio: 1,
			ok:    true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ok, err := rollout.ShouldRollout(tc.id, tc.ratio)
			if err != nil {
				t.Fatal("ShouldRollout fail:", err)
			}

			if got, want := ok, tc.ok; got != want {
				t.Errorf("ShouldRollout got %v, want %v", got, want)
			}
		})
	}
}

func TestShouldRollout_increasedRollout(t *testing.T) {
	testCases := map[string]struct {
		ratios         []float32
		iterationCount int
	}{
		"increased rollout - 2 steps": {
			ratios:         []float32{0.2, 0.5},
			iterationCount: 500,
		},
		"increased rollout - 3 steps": {
			ratios:         []float32{0.3, 0.5, 0.9},
			iterationCount: 500,
		},
		"increased rollout - 4 steps": {
			ratios:         []float32{0.2, 0.3, 0.5, 0.7},
			iterationCount: 500,
		},
		"increased rollout - 5 steps": {
			ratios:         []float32{0.2, 0.3, 0.5, 0.9},
			iterationCount: 500,
		},
	}

	for name, tc := range testCases {
		results := make(map[string]bool)

		t.Run(name, func(t *testing.T) {
			for i := 0; i < tc.iterationCount; i++ {
				id := xid.New().String()
				ok, err := rollout.ShouldRollout(id, tc.ratios[0])
				if err != nil {
					t.Fatal("ShouldRollout fail:", err)
				}

				results[id] = ok
			}

			for _, ratio := range tc.ratios[1:] {
				for id, prevOK := range results {
					ok, err := rollout.ShouldRollout(id, ratio)
					if err != nil {
						t.Fatal("ShouldRollout fail:", err)
					}

					if prevOK && !ok {
						t.Errorf("ShouldRollout %q initialOK=%t, ok=%t", id, prevOK, ok)
					}

					results[id] = ok
				}
			}
		})
	}
}
