package stat

import "sort"

func Percentile(arr []int, r float64) float64 {
	a := make([]int, len(arr))
	copy(a, arr)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	count := len(a)
	mid := count / 2
	if count%2 == 0 {
		return float64(a[mid-1]+a[mid]) / 2
	}

	return float64(a[mid])
}

func PercentileScore(arr []int, val int) float64 {
	var count int
	for _, v := range arr {
		if v < val {
			count++
		}
	}
	return float64(count) / float64(len(arr))
}
