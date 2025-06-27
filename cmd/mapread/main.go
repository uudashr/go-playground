package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[string]any{
		"timestamp": time.Now().UnixMilli(),
	}

	i, err := mapValue[int64](m, "timestamp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Timestamp:", i)

	s, err := mapValue[string](m, "timestamp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Timestamp as string:", s)
}

func mapValue[T any](m map[string]any, key string, validators ...ValidateFunc[T]) (T, error) {
	var zero T

	v, ok := m[key]
	if !ok {
		return zero, fmt.Errorf("key %q not found in map", key)
	}

	out, ok := v.(T)
	if !ok {
		return zero, fmt.Errorf("value for key %q is %T, not %T", key, v, zero)
	}

	for _, validate := range validators {
		if err := validate(out); err != nil {
			return zero, fmt.Errorf("validation failed for key %q: %w", key, err)
		}
	}

	return out, nil
}

func optMapValue[T any](m map[string]any, key string) (T, error) {
	var zero T
	return optMapValueWithDefault(m, key, zero)
}

func optMapValueWithDefault[T any](m map[string]any, key string, defVal T, validators ...ValidateFunc[T]) (T, error) {
	v, ok := m[key]
	if !ok {
		return defVal, nil
	}

	out, ok := v.(T)
	if !ok {
		return defVal, fmt.Errorf("value for key %q is %T, not %T", key, v, defVal)
	}

	for _, validate := range validators {
		if err := validate(out); err != nil {
			return defVal, fmt.Errorf("validation failed for key %q: %w", key, err)
		}
	}

	return out, nil
}

type ValidateFunc[T any] func(T) error

func NotEmpty(s string) error {
	if s == "" {
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
