package rollout

import (
	"errors"
	"hash/fnv"
)

func ShouldRollout(id string, ratio float32) (bool, error) {
	h := fnv.New32()
	if _, err := h.Write([]byte(id)); err != nil {
		return false, err
	}

	mod := h.Sum32() % 100
	if float32(mod) > ratio*100 {
		return false, nil
	}

	return true, nil
}

func DistributionIndex(id string, ratios []float32) (int, error) {
	if len(ratios) == 0 {
		return -1, errors.New("ratios is empty")
	}

	if err := validateRatios(ratios); err != nil {
		return -1, err
	}

	var sumRatio float32
	for _, ratio := range ratios {
		if ratio < 0 || ratio > 1 {
			return -1, errors.New("ratio must be in [0, 1]")
		}
		sumRatio += ratio
	}

	h := fnv.New32()
	if _, err := h.Write([]byte(id)); err != nil {
		return -1, err
	}

	hashVal := h.Sum32()
	segmentIndex := int(hashVal % 100) // put into segments [0 - 99]

	var limit int
	for i, ratio := range ratios[:len(ratios)-1] {
		limit += int(ratio * 100)
		if segmentIndex <= limit {
			return i, nil
		}
	}

	return len(ratios) - 1, nil
}

func validateRatios(ratios []float32) error {
	if len(ratios) == 0 {
		return errors.New("ratios is empty")
	}

	var sumRatio float32
	for _, ratio := range ratios {
		if ratio < 0 || ratio > 1 {
			return errors.New("ratio must be in [0, 1]")
		}
		sumRatio += ratio
	}

	if sumRatio != 1 {
		return errors.New("sum of ratios must be 1")
	}

	return nil
}
