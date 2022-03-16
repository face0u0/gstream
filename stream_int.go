package gstream

type IntStream struct {
	*NumericStream[int]
}

func NewIntStreamNumeric(list []int) *IntStream {
	return newIntStreamWithCtx(&sCtx[int]{list, new(Seqential[int])})
}

func newIntStreamWithCtx(ctx *sCtx[int]) *IntStream {
	return &IntStream{newNumericStreamWithCtx(ctx)}
}

func (s *IntStream) Filter(f func(int) bool) *IntStream {
	base := s.NumericStream.Filter(f)
	return &IntStream{base}
}
