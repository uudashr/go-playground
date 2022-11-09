package slices_test

import (
	"fmt"
	"strings"

	"github.com/uudashr/go-playground/slices"
)

func ExampleMap() {
	out := slices.Map([]int{1, 2, 3}, func(i int) int {
		return i * 2
	})

	// Output: [2 4 6]
	fmt.Println(out)
}

func ExampleMap_split() {
	splits := strings.Split("hello, world", ",")
	out := slices.Map(splits, strings.TrimSpace)

	// Output: [hello world]
	fmt.Println(out)
}
