package rollout

import "testing"

func TestDistribution(t *testing.T) {
	// input: segmentIndex, distributionRatios
	// output: distributionIndex

	testCases := map[string]struct {
		segmentIndex       int
		distributionRatios []float32
		expectIndex        int
	}{
		"1 distribution - segmentIndex 0": {
			segmentIndex:       0,
			distributionRatios: []float32{1},
			expectIndex:        0,
		},
		"1 distribution - segmentIndex 1": {
			segmentIndex:       1,
			distributionRatios: []float32{1},
			expectIndex:        0,
		},
		"1 distribution - segmentIndex 2": {
			segmentIndex:       2,
			distributionRatios: []float32{1},
			expectIndex:        0,
		},
		"2 distributions - segmentIndex 0": {
			segmentIndex:       0,
			distributionRatios: []float32{0.5, 0.5},
			expectIndex:        0,
		},
		"2 distributions - segmentIndex 49999": {
			segmentIndex:       49999,
			distributionRatios: []float32{0.5, 0.5},
			expectIndex:        0,
		},
		"2 distributions - segmentIndex 50000": {
			segmentIndex:       50000,
			distributionRatios: []float32{0.5, 0.5},
			expectIndex:        1,
		},
		"2 distributions - segmentIndex 99999": {
			segmentIndex:       99999,
			distributionRatios: []float32{0.5, 0.5},
			expectIndex:        1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			index, err := distributionIndex(tc.segmentIndex, tc.distributionRatios)
			if err != nil {
				t.Fatal("distributionIndex fail:", err)
			}

			if got, want := index, tc.expectIndex; got != want {
				t.Errorf("distributionIndex got=%d, want=%d", got, want)
			}
		})
	}

}
