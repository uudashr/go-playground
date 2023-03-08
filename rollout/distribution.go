package rollout

import (
	"errors"
)

var segmentSize = 100000

func distributionIndex(segmentIndex int, distributionRatios []float32) (int, error) {
	if segmentIndex < 0 {
		return -1, errors.New("segmentIndex must be >= 0")
	}

	if segmentIndex > segmentSize {
		return -1, errors.New("segmentIndex must be <= segmentSize")
	}

	if len(distributionRatios) == 0 {
		return -1, errors.New("distributionRatios must have at least 1 element")
	}

	var sumRatio float32
	for _, ratio := range distributionRatios {
		if ratio < 0 || ratio > 1 {
			return -1, errors.New("distributionRatios must be between 0 and 1")
		}

		sumRatio += ratio
	}

	if sumRatio != 1 {
		return -1, errors.New("distributionRatios must sum to 1")
	}

	var runningThreshold float32
	for i, ratio := range distributionRatios[:len(distributionRatios)-1] {
		segmentThreshold := ratio * float32(segmentSize)
		runningThreshold += segmentThreshold
		if float32(segmentIndex) < runningThreshold {
			return i, nil
		}
	}

	return len(distributionRatios) - 1, nil
}
