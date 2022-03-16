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

func (s *Stream[T]) Parallel() *Stream[T] {
	s.loop = ST_PARALLEL
	return s
}

func (s *Stream[T]) Sequential() *Stream[T] {
	s.loop = ST_SEQUENTIAL
	return s
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

func (s *Stream[T]) MapToInt(f func(T) int) *IntStream {
	nl := make([]int, 0)
	s.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newIntStreamWithCtx(newSCtxFrom(s.sCtx, nl))
}

func (s *Stream[T]) MapToStr(f func(T) string) *StringStream {
	nl := make([]string, 0)
	s.forEachSequential(func(val T) (brk bool) {
		nl = append(nl, f(val))
		return
	})
	return newStringStreamWithCtx(newSCtxFrom(s.sCtx, nl))
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

func (s *Stream[T]) AnyMatch(func(T) bool) {

}
