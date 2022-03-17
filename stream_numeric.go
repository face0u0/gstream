package gstream

type Numeric interface{ ~int | ~float64 }

type NumericStream[T Numeric] struct {
	*OrderStream[T]
}

func NewNumericStream[T Numeric](list []T) *NumericStream[T] {
	return newNumericStreamWithCtx(&sCtx[T]{list, ST_SEQUENTIAL})
}

func newNumericStreamWithCtx[T Numeric](ctx *sCtx[T]) *NumericStream[T] {
	return &NumericStream[T]{newOrderStreamWithCtx(ctx)}
}

func (s *NumericStream[T]) Parallel() *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Parallel()}
}

func (s *NumericStream[T]) Sequential() *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Sequential()}
}

func (s *NumericStream[T]) Skip(n int) *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Skip(n)}
}

func (s *NumericStream[T]) Limit(max int) *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Limit(max)}
}

func (s *NumericStream[T]) Filter(f func(T) bool) *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Filter(f)}
}

func (s *NumericStream[T]) Reverse() *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Reverse()}
}

func (s *NumericStream[T]) Sort() *NumericStream[T] {
	return &NumericStream[T]{s.OrderStream.Sort()}
}

func (s *NumericStream[T]) Sum() T {
	var _sum T
	s.forEachSequential(func(val T) (brk bool) {
		_sum = _sum + val
		return
	})
	return _sum
}
