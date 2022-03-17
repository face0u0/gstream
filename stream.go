package gstream

type Stream[T any] struct {
	*sCtx[T]
}

func NewStream[T any](list []T) *Stream[T] {
	return newStreamWithCtx(&sCtx[T]{list, ST_SEQUENTIAL})
}

func newStreamWithCtx[T any](ctx *sCtx[T]) *Stream[T] {
	return &Stream[T]{sCtx: ctx}
}

func Pass[T any](value T) T {
	return value
}

func Map[T, K any](src *Stream[T], f func(T) K) *Stream[K] {
	nl := make([]K, 0)
	src.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newStreamWithCtx(newSCtxFrom(src.sCtx, nl))
}

func (s *Stream[T]) Parallel() *Stream[T] {
	s.loop = ST_PARALLEL
	return s
}

func (s *Stream[T]) Sequential() *Stream[T] {
	s.loop = ST_SEQUENTIAL
	return s
}

func (s *Stream[T]) Skip(n int) *Stream[T] {
	return &Stream[T]{sCtx: newSCtxFrom(s.sCtx, s.values[n:])}
}

func (s *Stream[T]) Limit(max int) *Stream[T] {
	return &Stream[T]{sCtx: newSCtxFrom(s.sCtx, s.values[:max])}
}

func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	nl := make([]T, 0)
	s.forEachSequential(func(val T) (brk bool) {
		if f(val) {
			nl = append(nl, val)
		}
		return
	})
	return &Stream[T]{sCtx: newSCtxFrom(s.sCtx, nl)}
}

func (s *Stream[T]) Reverse() *Stream[T] {
	length := len(s.values)
	nl := make([]T, length)
	copy(nl, s.values)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		nl[i], nl[j] = nl[j], nl[i]
	}
	return &Stream[T]{newSCtxFrom(s.sCtx, nl)}
}

func (s *Stream[T]) MapToInt(f func(T) int) *IntStream {
	nl := make([]int, 0)
	s.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newIntStreamWithCtx(newSCtxFrom(s.sCtx, nl))
}

func (s *Stream[T]) MapToFloat(f func(T) float64) *FloatStream {
	nl := make([]float64, 0)
	s.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newFloatStreamWithCtx(newSCtxFrom(s.sCtx, nl))
}

func (s *Stream[T]) MapToStr(f func(T) string) *StringStream {
	nl := make([]string, 0)
	s.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newStringStreamWithCtx(newSCtxFrom(s.sCtx, nl))
}

func (s *Stream[T]) Map(f func(T) any) *Stream[any] {
	nl := make([]any, 0)
	s.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newStreamWithCtx(newSCtxFrom(s.sCtx, nl))
}

func (s *Stream[T]) ForEach(f func(T)) {
	s.forEachDefault(func(val T) (brk bool) {
		f(val)
		return
	})
}

func (s *Stream[T]) Count() int {
	return len(s.values)
}

func (s *Stream[T]) ToSlice() []T {
	return s.values
}
