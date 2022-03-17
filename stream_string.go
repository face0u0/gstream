package gstream

import "strings"

type StringStream struct {
	*Stream[string]
}

func NewStringStream(list []string) *StringStream {
	return newStringStreamWithCtx(&sCtx[string]{list, ST_SEQUENTIAL})
}

func newStringStreamWithCtx(ctx *sCtx[string]) *StringStream {
	return &StringStream{newStreamWithCtx(ctx)}
}

func (s *StringStream) Parallel() *StringStream {
	return &StringStream{s.Stream.Parallel()}
}

func (s *StringStream) Sequential() *StringStream {
	return &StringStream{s.Stream.Sequential()}
}

func (s *StringStream) Skip(n int) *StringStream {
	return &StringStream{s.Stream.Skip(n)}
}

func (s *StringStream) Limit(max int) *StringStream {
	return &StringStream{s.Stream.Limit(max)}
}

func (s *StringStream) Filter(f func(string) bool) *StringStream {
	return &StringStream{s.Stream.Filter(f)}
}

func (s *StringStream) Join(str string) string {
	return strings.Join(s.values, str)
}
