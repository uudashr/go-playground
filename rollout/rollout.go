package rollout

import (
	"errors"
	"hash/fnv"
)

var segmentSize = uint32(10)

func DistributionString(s string, distributionPercentages []uint32) (distributionIndex int, err error) {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		return -1, err
	}

	sum := h.Sum32()
	segmentIndex := sum % segmentSize

	return DistributionUint32(segmentIndex, distributionPercentages)
}

func Distribution(segmentIndex int, distributionPercentages []uint32) (distributionIndex int, err error) {
	if segmentIndex < 0 {
		return -1, errors.New("segment index cannot be less than 0")
	}

	return DistributionUint32(uint32(segmentIndex), distributionPercentages)
}

func DistributionUint32(segmentIndex uint32, distributionPercentages []uint32) (distributionIndex int, err error) {
	if segmentIndex >= segmentSize {
		return -1, errors.New("segment index cannot be greater than segment size")
	}

	if len(distributionPercentages) == 0 {
		return -1, errors.New("distribution percentages cannot be empty")
	}

	var sum uint32
	for _, percentage := range distributionPercentages {
		if percentage > 100 {
			return -1, errors.New("percentage cannot be greater than 100")
		}

		sum += percentage
	}

	if sum > 100 {
		return -1, errors.New("total percentages cannot be greater than 100")
	}

	if len(distributionPercentages) == 1 {
		return 0, nil
	}

	var (
		incrementalPercentage float64
		segmentIndexFloat64   = float64(segmentIndex)
	)

	for i, percentage := range distributionPercentages {
		r := float64(percentage) / 100
		r += incrementalPercentage
		segmentLimit := float64(segmentSize) * r
		if segmentIndexFloat64 < segmentLimit {
			return i, nil
		}

		incrementalPercentage += r
	}

	panic("should not reach here")
}
