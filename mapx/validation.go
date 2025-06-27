package mapx

import "fmt"

type LenCapable interface {
	~string | ~[]int | ~[]string | ~map[string]int | ~chan int
}

type ValidateFunc[T any] func(T) error

func NotEmpty[T LenCapable](v T) error {
	if len(v) == 0 {
		return fmt.Errorf("value is empty")
	}

	return nil
}

func NotZero[T comparable](v T) error {
	var zero T

	if v == zero {
		return fmt.Errorf("value is zero")
	}

	return nil
}
