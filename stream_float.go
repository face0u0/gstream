package gstream

type FloatStream struct {
	*NumericStream[float64]
}

func NewFloatStreamNumeric(list []float64) *FloatStream {
	return newFloatStreamWithCtx(&sCtx[float64]{list, ST_SEQUENTIAL})
}

func newFloatStreamWithCtx(ctx *sCtx[float64]) *FloatStream {
	return &FloatStream{newNumericStreamWithCtx(ctx)}
}

func (s *FloatStream) Filter(f func(float64) bool) *FloatStream {
	base := s.NumericStream.Filter(f)
	return &FloatStream{base}
}
