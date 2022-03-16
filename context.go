package gstream

type sCtx[T any] struct {
	data        []T
	defaultLoop LoopStrat[T]
}

func newSCtxWithoutData[K, T any](src *sCtx[K], data []T) *sCtx[T] {
	return &sCtx[T]{data: data, defaultLoop: newLoopStratSameAs[K, T](src.defaultLoop)}
}

func (s *sCtx[T]) setData(data []T) *sCtx[T] {
	s.data = data
	return s
}

// func (s *sCtx[T]) setLoop(loop LoopStrat[T]) *sCtx[T] {
// 	s.defaultLoop = loop
// 	return s
// }
