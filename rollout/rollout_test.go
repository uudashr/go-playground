package rollout_test

import (
	"testing"

	"github.com/uudashr/go-playground/rollout"
)

func TestDistribution(t *testing.T) {
	testCases := map[string]struct {
		segmentIndex            int
		distributionPercentages []uint32
		expectDistIndex         int
	}{
		"single distribution - 0 dist index": {
			segmentIndex:            0,
			distributionPercentages: []uint32{100},
			expectDistIndex:         0,
		},
		"single distribution - 4 dist index": {
			segmentIndex:            4,
			distributionPercentages: []uint32{100},
			expectDistIndex:         0,
		},
		"single distribution - 5 dist index": {
			segmentIndex:            5,
			distributionPercentages: []uint32{100},
			expectDistIndex:         0,
		},
		"single distribution - 9 dist index": {
			segmentIndex:            9,
			distributionPercentages: []uint32{100},
			expectDistIndex:         0,
		},

		"multiple distribution - 0 dist index": {
			segmentIndex:            0,
			distributionPercentages: []uint32{50, 50},
			expectDistIndex:         0,
		},
		"multiple distribution - 4 dist index": {
			segmentIndex:            4,
			distributionPercentages: []uint32{50, 50},
			expectDistIndex:         0,
		},
		"multiple distribution - 5 dist index": {
			segmentIndex:            5,
			distributionPercentages: []uint32{50, 50},
			expectDistIndex:         1,
		},
		"multiple distribution - 9 dist index": {
			segmentIndex:            5,
			distributionPercentages: []uint32{50, 50},
			expectDistIndex:         1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			distIndex, err := rollout.Distribution(tc.segmentIndex, tc.distributionPercentages)
			if err != nil {
				t.Fatal("Distribution fail", err)
			}

			if got, want := distIndex, tc.expectDistIndex; got != want {
				t.Fatalf("Distribution index got %d, want %d", got, want)
			}
		})
	}
}
