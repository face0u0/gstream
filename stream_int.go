package gstream

type IntStream struct {
	*NumericStream[int]
}

func NewIntStreamNumeric(list []int) *IntStream {
	return newIntStreamWithCtx(&sCtx[int]{list, ST_SEQUENTIAL})
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
	base := s.NumericStream.Filter(f)
	return &IntStream{base}
}
