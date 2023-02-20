// Package specification implement the specification pattern.
// Implementation can be in form of Hardcoded specification,
// Parameterized specification or Composite specification.
// Based on https://martinfowler.com/apsupp/spec.pdf
package spec

// Spec is the generic interface of specification.
type Spec[T any] interface {
	SatisfiedBy(T) bool
}

type Func[T any] func(T) bool

func (f Func[T]) SatisfiedBy(v T) bool {
	return f(v)
}

func OfFunc[T any](f func(T) bool) Spec[T] {
	return Func[T](f)
}

func Not[T any](s Spec[T]) Spec[T] {
	return Func[T](func(v T) bool {
		return !s.SatisfiedBy(v)
	})
}

func And[T any](s1, s2 Spec[T]) Spec[T] {
	return Func[T](func(v T) bool {
		return s1.SatisfiedBy(v) && s2.SatisfiedBy(v)
	})
}

func AndNot[T any](s1, s2 Spec[T]) Spec[T] {
	return And(s1, Not(s2))
}

func Or[T any](s1, s2 Spec[T]) Spec[T] {
	return Func[T](func(v T) bool {
		return s1.SatisfiedBy(v) || s2.SatisfiedBy(v)
	})
}

func OrNot[T any](s1, s2 Spec[T]) Spec[T] {
	return Or(s1, Not(s2))
}
