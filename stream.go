package gstream

type Stream[T any] struct {
	*sCtx[T]
}

func NewStream[T any](list []T) *Stream[T] {
	return newStreamWithCtx(&sCtx[T]{list, new(Seqential[T])})
}

func newStreamWithCtx[T any](ctx *sCtx[T]) *Stream[T] {
	return &Stream[T]{ctx}
}

func (s *Stream[T]) ForEach(f func(T)) {
	s.defaultLoop.forEach(s.data, func(val T) {
		f(val)
	})
}

func (s *Stream[T]) forEachSeqly(f func(T)) {
	new(Seqential[T]).forEach(s.data, func(val T) {
		f(val)
	})
}

func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	nl := make([]T, 0)
	s.forEachSeqly(func(t T) {
		if f(t) {
			nl = append(nl, t)
		}
	})
	return &Stream[T]{newSCtxWithoutData(s.sCtx, nl)}
}

func (s *Stream[T]) MapInt(f func(T) int) *IntStream {
	nlist := make([]int, 0)
	s.ForEach(func(t T) {
		nlist = append(nlist, f(t))
	})
	return newIntStreamWithCtx(newSCtxWithoutData(s.sCtx, nlist))
}

func (s *Stream[T]) MapStr(f func(T) string) *StringStream {
	nlist := make([]string, 0)
	s.ForEach(func(t T) {
		nlist = append(nlist, f(t))
	})
	return newStringStreamWithCtx(newSCtxWithoutData(s.sCtx, nlist))
}

func (s *Stream[T]) Count() int {
	return len(s.data)
}

func (s *Stream[T]) AnyMatch(func(T) bool) {

}
