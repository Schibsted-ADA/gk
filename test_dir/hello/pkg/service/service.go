package service

import "context"

type HelloService interface {
	// Write your interface methods
	World(ctx context.Context, s string) (rs string, err error)
}

type stubHelloService struct{}

func New() (s stubHelloService) {
	s = stubHelloService{}
	return s
}
func (he *stubHelloService) World(ctx context.Context, s string) (rs string, err error) {
	// Implement your business logic here
	return rs, err
}
