package service

import "context"

type TestHttpService interface {
	// Write your interface methods
	Foo(ctx context.Context, s string) (string, error)
	Bar(ctx context.Context, i int) (int, error)
}

type stubTestHttpService struct{}

func New() (s *stubTestHttpService) {
	s = &stubTestHttpService{}
	return s
}
func (te *stubTestHttpService) Foo(ctx context.Context, s string) (s0 string, e1 error) {
	// Implement your business logic here
	return s0, e1
}
func (te *stubTestHttpService) Bar(ctx context.Context, i int) (i0 int, e1 error) {
	// Implement your business logic here
	return i0, e1
}
