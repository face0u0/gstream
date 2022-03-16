package gstream

type Numeric interface{ ~int | ~float64 }

type NumericStream[T Numeric] struct {
	*Stream[T]
}

func NewNumericStream[T Numeric](list []T) *NumericStream[T] {
	return newNumericStreamWithCtx(&sCtx[T]{list, new(Seqential[T])})
}

func newNumericStreamWithCtx[T Numeric](ctx *sCtx[T]) *NumericStream[T] {
	return &NumericStream[T]{newStreamWithCtx(ctx)}
}

func (s *NumericStream[T]) Filter(f func(T) bool) *NumericStream[T] {
	base := s.Stream.Filter(f)
	return &NumericStream[T]{base}
}

func (s *NumericStream[T]) Sum() T {
	var _sum T
	s.ForEach(func(t T) {
		_sum = _sum + t
	})
	return _sum
}
