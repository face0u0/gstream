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

type streamData[T any] struct {
	data        []T
	defaultLoop iterLoop[T]
}

func (s *streamData[T]) setData(data []T) *streamData[T] {
	s.data = data
	return s
}

func (s *streamData[T]) setLoop(loop iterLoop[T]) *streamData[T] {
	s.defaultLoop = loop
	return s
}

type Stream[T any] struct {
	sd streamData[T]
}

func (s *Stream[T]) setStreamData(data *streamData[T]) {
	s.sd = *data
}

func NewStream[T any](list []T) *Stream[T] {
	sd := streamData[T]{list, &sequentialLoop[T]{}}
	return &Stream[T]{sd}
}

func (s *Stream[T]) mapInt(f func(T) int) *IntStream {
	nlist := make([]int, 0)
	s.sd.defaultLoop.forEach(s.sd.data, func(val T) {
		nlist = append(nlist, f(val))
	})
	ns := &IntStream{&NumericStream[int]{&Stream[int]{sd: streamData[int]{nlist, &sequentialLoop[int]{}}}}}
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
	for _, v := range s.sd.data {
		base = base + v
	}
	return base
}

type NumericStream[T ~int | ~float64] struct {
	*Stream[T]
}

func (s *NumericStream[T]) sum() T {
	var _sum T
	for _, v := range s.sd.data {
		_sum = _sum + v
	}
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
