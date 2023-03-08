package rollout

import "hash/fnv"

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
	h := fnv.New32()
	if _, err := h.Write([]byte(id)); err != nil {
		return -1, err
	}

	sum := h.Sum32()
	mod := sum % uint32(segmentSize)

	return distributionIndex(int(mod), ratios)
}
