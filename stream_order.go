package gstream

import "sort"

type Order interface{ Numeric | ~string }

type OrderStream[T Order] struct {
	*Stream[T]
}

func NewOrderStream[T Order](list []T) *OrderStream[T] {
	return newOrderStreamWithCtx(&sCtx[T]{list, ST_SEQUENTIAL})
}

func newOrderStreamWithCtx[T Order](ctx *sCtx[T]) *OrderStream[T] {
	return &OrderStream[T]{newStreamWithCtx(ctx)}
}

func (s *OrderStream[T]) Parallel() *OrderStream[T] {
	return &OrderStream[T]{s.Stream.Parallel()}
}

func (s *OrderStream[T]) Sequential() *OrderStream[T] {
	return &OrderStream[T]{s.Stream.Sequential()}
}

func (s *OrderStream[T]) Skip(n int) *OrderStream[T] {
	return &OrderStream[T]{s.Stream.Skip(n)}
}

func (s *OrderStream[T]) Limit(max int) *OrderStream[T] {
	return &OrderStream[T]{s.Stream.Limit(max)}
}

func (s *OrderStream[T]) Filter(f func(T) bool) *OrderStream[T] {
	return &OrderStream[T]{s.Stream.Filter(f)}
}

func (s *OrderStream[T]) Reverse() *OrderStream[T] {
	return &OrderStream[T]{s.Stream.Reverse()}
}

func (s *OrderStream[T]) Sort() *OrderStream[T] {
	nl := make([]T, len(s.values))
	copy(nl, s.values)
	sort.Slice(nl, func(i, j int) bool { return nl[i] < nl[j] })
	return &OrderStream[T]{&Stream[T]{newSCtxFrom(s.sCtx, nl)}}
}

func (s *OrderStream[T]) Max() T {
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

func (s *OrderStream[T]) Min() T {
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
