package main

type Stream[T any] struct {
	data []T
}

func NewStream[T any](list []T) *Stream[T] {
	return &Stream[T]{list}
}

func (s *Stream[T]) mapInt(f func(T) int) IntStream {
	return IntStream{}
}

func (s *Stream[T]) mapStr(f func(T) string) StringStream {
	return StringStream{}
}

type IntStream struct {
	*NumericStream[int]
}

type StringStream struct {
	Stream[string]
}

func (s *StringStream) Join(str string) string {
	base := ""
	for _, v := range s.data {
		base = base + v
	}
	return base
}

type NumericStream[T ~int | ~float64] struct {
	*Stream[T]
}

func (s *NumericStream[T]) sum() T {
	var _sum T
	for _, v := range s.data {
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
	ss := s.mapInt(Cast[int]).sum()
}
