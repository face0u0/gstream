package gstream

const (
	ST_SEQUENTIAL = iota
	ST_PARALLEL   = iota
)

type LoopStrat[T any] interface {
	forEach([]T, func(T))

	Name() int
}

func newLoopStratSameAs[T, K any](strat LoopStrat[T]) LoopStrat[K] {
	switch strat.Name() {
	case ST_SEQUENTIAL:
		return &Seqential[K]{}
	case ST_PARALLEL:
		return &Parallel[K]{}
	default:
		return &Seqential[K]{}
	}
}

type Seqential[T any] struct{}

func (s *Seqential[T]) forEach(values []T, f func(T)) {
	for _, v := range values {
		f(v)
	}
}

func (s *Seqential[T]) Name() int {
	return ST_SEQUENTIAL
}

type Parallel[T any] struct{}

func (s *Parallel[T]) forEach(values []T, f func(T)) {
	for _, v := range values {
		f(v)
	}
}

func (s *Parallel[T]) Name() int {
	return ST_PARALLEL
}
