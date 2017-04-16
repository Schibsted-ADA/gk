package service

import "context"

type TestGrpcService interface {
	// Write your interface methods
	Foo(ctx context.Context, s string) (string, error)
	Bar(ctx context.Context, i int) (int, error)
}

type stubTestGrpcService struct{}

func New() (s *stubTestGrpcService) {
	s = &stubTestGrpcService{}
	return s
}
func (te *stubTestGrpcService) Foo(ctx context.Context, s string) (s0 string, e1 error) {
	// Implement your business logic here
	return s0, e1
}
func (te *stubTestGrpcService) Bar(ctx context.Context, i int) (i0 int, e1 error) {
	// Implement your business logic here
	return i0, e1
}
