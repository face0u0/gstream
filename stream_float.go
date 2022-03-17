package gstream

type FloatStream struct {
	*NumericStream[float64]
}

func NewFloatStream(list []float64) *FloatStream {
	return newFloatStreamWithCtx(&sCtx[float64]{list, ST_SEQUENTIAL})
}

func newFloatStreamWithCtx(ctx *sCtx[float64]) *FloatStream {
	return &FloatStream{newNumericStreamWithCtx(ctx)}
}

func (s *FloatStream) Parallel() *FloatStream {
	return &FloatStream{s.NumericStream.Parallel()}
}

func (s *FloatStream) Sequential() *FloatStream {
	return &FloatStream{s.NumericStream.Sequential()}
}

func (s *FloatStream) Skip(n int) *FloatStream {
	return &FloatStream{s.NumericStream.Skip(n)}
}

func (s *FloatStream) Limit(max int) *FloatStream {
	return &FloatStream{s.NumericStream.Limit(max)}
}

func (s *FloatStream) Filter(f func(float64) bool) *FloatStream {
	return &FloatStream{s.NumericStream.Filter(f)}
}

func (s *FloatStream) Reverse() *FloatStream {
	return &FloatStream{s.NumericStream.Reverse()}
}

func (s *FloatStream) Sort() *FloatStream {
	return &FloatStream{s.NumericStream.Sort()}
}
