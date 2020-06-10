package stat_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uudashr/go-playground/stat"
)

func ExamplePercentileScore() {
	arr := []int{1, 2, 4, 5}
	score := stat.PercentileScore(arr, 3)
	fmt.Println(score)
	// Output: 0.5
}

func ExamplePercentile() {
	arr := []int{1, 2, 3, 4, 5}
	p := stat.Percentile(arr, 0.5)
	fmt.Println(p)
	// Output: 3
}

func TestPercentileScore(t *testing.T) {
	testCases := map[string]struct {
		arr   []int
		val   int
		score float64
	}{
		"Simmple sorted data": {
			arr:   []int{1, 2, 4, 5},
			val:   3,
			score: 0.5,
		},
		"Simmple sorted data, none": {
			arr:   []int{1, 2, 4, 5},
			val:   1,
			score: 0,
		},
		"Simmple sorted data, all": {
			arr:   []int{1, 2, 4, 5},
			val:   6,
			score: 1,
		},
		"Simmple sorted data, one not included": {
			arr:   []int{1, 2, 4, 5},
			val:   5,
			score: 0.75,
		},
		"Simmple sorted data, only one included": {
			arr:   []int{1, 2, 4, 5},
			val:   2,
			score: 0.25,
		},
	}

	for k, tc := range testCases {
		t.Run(k, func(t *testing.T) {
			score := stat.PercentileScore(tc.arr, tc.val)
			if got, want := score, tc.score; got != want {
				t.Errorf("Score got: %f, want: %f", got, want)
			}
		})
	}
}

func TestPercentile(t *testing.T) {
	testCases := map[string]struct {
		arr []int
		r   float64
		p   float64
	}{
		"Data with center element": {
			arr: []int{1, 2, 3, 4, 5},
			r:   0.5,
			p:   3,
		},
		"Data with no center element": {
			arr: []int{1, 2, 3, 4},
			r:   0.5,
			p:   2.5,
		},
		"Data with center element, unsorted": {
			arr: []int{3, 1, 5, 4, 2},
			r:   0.5,
			p:   3,
		},
		"Data with no center element, unsorted": {
			arr: []int{2, 1, 4, 3},
			r:   0.5,
			p:   2.5,
		},
	}

	for k, tc := range testCases {
		t.Run(k, func(t *testing.T) {
			p := stat.Percentile(tc.arr, tc.r)
			if got, want := p, tc.p; got != want {
				t.Errorf("Percentile got: %f, want: %f", got, want)
			}
		})
	}
}

func TestArrNotChanged(t *testing.T) {
	arr := []int{2, 1, 3}
	expectArr := []int{2, 1, 3}
	_ = stat.Percentile(arr, 0.5)

	if got, want := arr, expectArr; !reflect.DeepEqual(got, want) {
		t.Errorf("Arr got: %v, want: %v", got, want)
	}
}
