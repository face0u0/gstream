package gstream

import "strings"

type StringStream struct {
	*OrderStream[string]
}

func NewStringStream(list []string) *StringStream {
	return newStringStreamWithCtx(&sCtx[string]{values: list, loop: ST_SEQUENTIAL})
}

func newStringStreamWithCtx(ctx *sCtx[string]) *StringStream {
	return &StringStream{newOrderStreamWithCtx(ctx)}
}

func (s *StringStream) Parallel() *StringStream {
	return &StringStream{s.OrderStream.Parallel()}
}

func (s *StringStream) Sequential() *StringStream {
	return &StringStream{s.OrderStream.Sequential()}
}

func (s *StringStream) Skip(n int) *StringStream {
	return &StringStream{s.OrderStream.Skip(n)}
}

func (s *StringStream) Limit(max int) *StringStream {
	return &StringStream{s.OrderStream.Limit(max)}
}

func (s *StringStream) Filter(f func(string) bool) *StringStream {
	return &StringStream{s.OrderStream.Filter(f)}
}

func (s *StringStream) ErrorFilter(f func(string) error) *StringStream {
	return &StringStream{s.OrderStream.ErrorFilter(f)}
}

func (s *StringStream) Reverse() *StringStream {
	return &StringStream{s.OrderStream.Reverse()}
}

func (s *StringStream) Sort() *StringStream {
	return &StringStream{s.OrderStream.Sort()}
}

func (s *StringStream) Join(str string) string {
	return strings.Join(s.values, str)
}
