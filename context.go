package gstream

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

const (
	ST_SEQUENTIAL = iota
	ST_PARALLEL   = iota
)

type sCtx[T any] struct {
	values []T
	errors []error
	loop   int
}

func newSCtxFrom[T, K any](src *sCtx[T], values []K) *sCtx[K] {
	return &sCtx[K]{values: values, loop: ST_SEQUENTIAL, errors: src.errors}
}

func (c *sCtx[T]) forEachSequential(f func(val T) (brk bool)) {
	for _, v := range c.values {
		if f(v) {
			break
		}
	}
}

func (c *sCtx[T]) forEachParallel(f func(val T) (brk bool)) {
	length := len(c.values)
	threads := min(runtime.NumCPU(), length)
	var breaked bool
	var wg sync.WaitGroup
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		step := int(math.Ceil(float64(length) / float64(threads)))
		start := i * step
		end := min((i+1)*step, length)
		fmt.Println(start, end)
		go func(src []T) {
			defer wg.Done()
			for _, v := range src {
				if f(v) {
					breaked = true
					return
				}
				if breaked {
					return
				}
			}
		}(c.values[start:end])
	}
	wg.Wait()
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

func min[T ~int](a T, b T) T {
	if a > b {
		return b
	}
	return a
}
