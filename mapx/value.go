package mapx

import (
	"fmt"
)

func Value[T any](m map[string]any, key string, validators ...ValidateFunc[T]) (T, error) {
	var zero T

	v, ok := m[key]
	if !ok {
		return zero, fmt.Errorf("key %q not found", key)
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
