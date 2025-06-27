package mapx

import "fmt"

func OptValue[T any](m map[string]any, key string, validators ...ValidateFunc[T]) (T, error) {
	var zero T

	return OptValueWithDefault(m, key, zero, validators...)
}

func OptValueWithDefault[T any](m map[string]any, key string, defaultVal T, validators ...ValidateFunc[T]) (T, error) {
	v, ok := m[key]
	if !ok {
		return defaultVal, nil
	}

	out, ok := v.(T)
	if !ok {
		return defaultVal, fmt.Errorf("value for key %q is %T, not %T", key, v, defaultVal)
	}

	for _, validate := range validators {
		if err := validate(out); err != nil {
			return defaultVal, fmt.Errorf("validation failed for key %q: %w", key, err)
		}
	}

	return out, nil
}
