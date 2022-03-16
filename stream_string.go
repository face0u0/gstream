package gstream

type StringStream struct {
	*Stream[string]
}

func NewStringStreamNumeric(list []string) *StringStream {
	return newStringStreamWithCtx(&sCtx[string]{list, new(Seqential[string])})
}

func newStringStreamWithCtx(ctx *sCtx[string]) *StringStream {
	return &StringStream{newStreamWithCtx(ctx)}
}

func (s *StringStream) Filter(f func(string) bool) *StringStream {
	base := s.Stream.Filter(f)
	return &StringStream{base}
}

func (s *StringStream) Join(str string) string {
	base := ""
	s.forEachSeqly(func(str string) {
		base = base + str
	})
	return base
}
