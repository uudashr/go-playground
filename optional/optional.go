package optional

type Value[T any] struct {
	value T
	ok    bool
}

func ValueOf[T any](value T) Value[T] {
	return Value[T]{value: value, ok: true}
}

func (v Value[T]) Get() (T, bool) {
	return v.value, v.ok
}

func (v Value[T]) Empty() bool {
	return !v.ok
}

func (v Value[T]) Default(defaultValue T) T {
	if !v.Empty() {
		return v.value
	}

	return defaultValue
}
