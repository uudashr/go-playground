package spec

type Composite[T any] struct {
	wrapped Spec[T]
}

func Wrap[T any](s Spec[T]) *Composite[T] {
	return &Composite[T]{wrapped: s}
}

func (s *Composite[T]) SatisfiedBy(v T) bool {
	return s.wrapped.SatisfiedBy(v)
}

func (s *Composite[T]) And(s2 Spec[T]) *Composite[T] {
	return Wrap(And(s.wrapped, s2))
}

func (s *Composite[T]) AndNot(s2 Spec[T]) *Composite[T] {
	return s.And(Not(s2))
}

func (s *Composite[T]) Or(s2 Spec[T]) *Composite[T] {
	return Wrap(Or(s.wrapped, s2))
}

func (s *Composite[T]) OrNot(s2 Spec[T]) *Composite[T] {
	return s.Or(Not(s2))
}

func (s *Composite[T]) Not() *Composite[T] {
	return Wrap(Not(s.wrapped))
}
