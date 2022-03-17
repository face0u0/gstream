package gstream

type IntStream struct {
	*NumericStream[int]
}

func NewIntStream(list []int) *IntStream {
	return newIntStreamWithCtx(&sCtx[int]{values: list, loop: ST_SEQUENTIAL})
}

func newIntStreamWithCtx(ctx *sCtx[int]) *IntStream {
	return &IntStream{newNumericStreamWithCtx(ctx)}
}

func (s *IntStream) Parallel() *IntStream {
	return &IntStream{s.NumericStream.Parallel()}
}

func (s *IntStream) Sequential() *IntStream {
	return &IntStream{s.NumericStream.Sequential()}
}

func (s *IntStream) Skip(n int) *IntStream {
	return &IntStream{s.NumericStream.Skip(n)}
}

func (s *IntStream) Limit(max int) *IntStream {
	return &IntStream{s.NumericStream.Limit(max)}
}

func (s *IntStream) Filter(f func(int) bool) *IntStream {
	return &IntStream{s.NumericStream.Filter(f)}
}

func (s *IntStream) ErrorFilter(f func(int) error) *IntStream {
	return &IntStream{s.NumericStream.ErrorFilter(f)}
}

func (s *IntStream) Reverse() *IntStream {
	return &IntStream{s.NumericStream.Reverse()}
}

func (s *IntStream) Sort() *IntStream {
	return &IntStream{s.NumericStream.Sort()}
}
