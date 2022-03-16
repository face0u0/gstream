package gstream

import "fmt"

// type streamCtx[T any] struct {
// 	data        []T
// 	defaultLoop iterLoop[T]
// }

// func (s *streamCtx[T]) setData(data []T) *streamCtx[T] {
// 	s.data = data
// 	return s
// }

// func (s *streamCtx[T]) setLoop(loop iterLoop[T]) *streamCtx[T] {
// 	s.defaultLoop = loop
// 	return s
// }

// type Stream[K any, T BaseStream[K]] struct {
// 	*BaseStream[K]
// }

// func (s *Stream[K, T]) filter(f func(K) bool) *T {
// 	nl := make([]K, 0)
// 	s.forEachSeqly(func(t K) {
// 		if f(t) {
// 			nl = append(nl, t)
// 		}
// 	})
// 	return &T{s.setData(nl)}
// }

// type BaseStream[T any] struct {
// 	*streamCtx[T]
// }

// func (s *BaseStream[T]) ForEach(f func(T)) {
// 	s.defaultLoop.forEach(s.data, func(val T) {
// 		f(val)
// 	})
// }

// func (s *BaseStream[T]) forEachSeqly(f func(T)) {
// 	new(sequentialLoop[T]).forEach(s.data, func(val T) {
// 		f(val)
// 	})
// }

// func (s *BaseStream[T]) filter(f func(T) bool) *BaseStream[T] {
// 	nl := make([]T, 0)
// 	s.forEachSeqly(func(t T) {
// 		if f(t) {
// 			nl = append(nl, t)
// 		}
// 	})
// 	return &BaseStream[T]{s.setData(nl)}
// }

// func NewStream[T any](list []T) *BaseStream[T] {
// 	sd := streamCtx[T]{list, &sequentialLoop[T]{}}
// 	return &BaseStream[T]{&sd}
// }

// func (s *BaseStream[T]) mapInt(f func(T) int) *IntStream {
// 	nlist := make([]int, 0)
// 	s.ForEach(func(t T) {
// 		nlist = append(nlist, f(t))
// 	})
// 	ns := &IntStream{&NumericStream[int]{&BaseStream[int]{&streamCtx[int]{nlist, &Seqential[int]{}}}}}
// 	return ns
// }

func Cast[T any](value T) T {
	return value
}

func main() {
	s := NewStream([]int{1, 2, 3, 4}).MapInt(Cast[int]).Sum()
	// s.Join(",")
	fmt.Println(s)
}
