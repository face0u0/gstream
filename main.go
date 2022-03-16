package main

import "fmt"

type iterLoop[T any] interface {
	forEach([]T, func(T))
}

type sequentialLoop[T any] struct{}

func (s *sequentialLoop[T]) forEach(values []T, f func(T)) {
	for _, v := range values {
		f(v)
	}
}

type streamCtx[T any] struct {
	data        []T
	defaultLoop iterLoop[T]
}

func (s *streamCtx[T]) setData(data []T) *streamCtx[T] {
	s.data = data
	return s
}

func (s *streamCtx[T]) setLoop(loop iterLoop[T]) *streamCtx[T] {
	s.defaultLoop = loop
	return s
}

type Stream[T any] struct {
	*streamCtx[T]
}

func (s *Stream[T]) ForEach(f func(T)) {
	s.defaultLoop.forEach(s.data, func(val T) {
		f(val)
	})
}

func (s *Stream[T]) forEachSeqly(f func(T)) {
	new(sequentialLoop[T]).forEach(s.data, func(val T) {
		f(val)
	})
}

func (s *Stream[T]) filter(f func(T) bool) *Stream[T] {
	nl := make([]T, 0)
	s.forEachSeqly(func(t T) {
		if f(t) {
			nl = append(nl, t)
		}
	})
	return &Stream[T]{s.setData(nl)}
}

func NewStream[T any](list []T) *Stream[T] {
	sd := streamCtx[T]{list, &sequentialLoop[T]{}}
	return &Stream[T]{&sd}
}

func (s *Stream[T]) mapInt(f func(T) int) *IntStream {
	nlist := make([]int, 0)
	s.ForEach(func(t T) {
		nlist = append(nlist, f(t))
	})
	ns := &IntStream{&NumericStream[int]{&Stream[int]{&streamCtx[int]{nlist, &sequentialLoop[int]{}}}}}
	return ns
}

func (s *Stream[T]) mapStr(f func(T) string) *StringStream {
	return &StringStream{}
}

type IntStream struct {
	*NumericStream[int]
}

type StringStream struct {
	Stream[string]
}

func (s *StringStream) Join(str string) string {
	base := ""
	s.forEachSeqly(func(str string) {
		base = base + str
	})
	return base
}

type NumericStream[T ~int | ~float64] struct {
	*Stream[T]
}

func (s *NumericStream[T]) sum() T {
	var _sum T
	s.ForEach(func(t T) {
		_sum = _sum + t
	})
	return _sum
}

func Cast[T any](value T) T {
	return value
}

func main() {
	s := NewStream([]int{1, 2, 3, 4}).mapInt(Cast[int]).sum()
	// s.Join(",")
	fmt.Println(s)
}
