package gstream

const (
	ST_SEQUENTIAL = iota
	ST_PARALLEL   = iota
)

type sCtx[T any] struct {
	values []T
	loop   int
}

func newSCtxFrom[T, K any](src *sCtx[T], values []K) *sCtx[K] {
	return &sCtx[K]{values: values, loop: ST_SEQUENTIAL}
}

func (c *sCtx[T]) forEachSequential(f func(val T) (brk bool)) {
	for _, v := range c.values {
		if f(v) {
			break
		}
	}
}

func (c *sCtx[T]) forEachParallel(f func(val T) (brk bool)) {
	for _, v := range c.values {
		if f(v) {
			break
		}
	}
}

func (c *sCtx[T]) forEachDefault(f func(val T) (brk bool)) {
	switch c.loop {
	case ST_SEQUENTIAL:
		c.forEachSequential(f)
	case ST_PARALLEL:
		c.forEachParallel(f)
	default:
		c.forEachSequential(f)
	}
}
