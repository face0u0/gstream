package gstream

type Numeric interface{ ~int | ~float64 }

type NumericStream[T Numeric] struct {
	*Stream[T]
}

func NewNumericStream[T Numeric](list []T) *NumericStream[T] {
	return newNumericStreamWithCtx(&sCtx[T]{list, ST_SEQUENTIAL})
}

func newNumericStreamWithCtx[T Numeric](ctx *sCtx[T]) *NumericStream[T] {
	return &NumericStream[T]{newStreamWithCtx(ctx)}
}

func (s *NumericStream[T]) Parallel() *NumericStream[T] {
	return &NumericStream[T]{s.Stream.Parallel()}
}

func (s *NumericStream[T]) Sequential() *NumericStream[T] {
	return &NumericStream[T]{s.Stream.Sequential()}
}

func (s *NumericStream[T]) Skip(n int) *NumericStream[T] {
	return &NumericStream[T]{s.Stream.Skip(n)}
}

func (s *NumericStream[T]) Limit(max int) *NumericStream[T] {
	return &NumericStream[T]{s.Stream.Limit(max)}
}

func (s *NumericStream[T]) Filter(f func(T) bool) *NumericStream[T] {
	return &NumericStream[T]{s.Stream.Filter(f)}
}

func (s *NumericStream[T]) Sum() T {
	var _sum T
	s.forEachSequential(func(val T) (brk bool) {
		_sum = _sum + val
		return
	})
	return _sum
}

func (s *NumericStream[T]) Max() T {
	updated := false
	var max T
	s.forEachSequential(func(val T) (brk bool) {
		if !updated {
			max = val
			updated = true
			return
		}
		if max < val {
			max = val
		}
		return
	})
	if !updated {
		panic("0 len stream not define max")
	}
	return max
}

func (s *NumericStream[T]) Min() T {
	updated := false
	var min T
	s.forEachSequential(func(val T) (brk bool) {
		if !updated {
			min = val
			updated = true
			return
		}
		if min > val {
			min = val
		}
		return
	})
	if !updated {
		panic("0 len stream not define max")
	}
	return min
}
