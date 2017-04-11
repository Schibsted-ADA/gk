package service

import "context"

type MyService interface {
	// Write your interface methods
	Foo(ctx context.Context, a []string) (map[string]string, error)
}

type stubMyService struct{}

func New() (s stubMyService) {
	s = stubMyService{}
	return s
}
func (my *stubMyService) Foo(ctx context.Context, a []string) (m0 map[string]string, e1 error) {
	// Implement your business logic here
	return m0, e1
}
