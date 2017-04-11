package service

import (
	"context"

	"github.com/kujtimiihoxha/gk/test_dir/blla"
)

type HiService interface {
	// Write your interface methods
	Foo(ctx context.Context, mp []map[[]string][]*blla.Bar) []blla.Bar
}

type stubHiService struct{}

func New() (s stubHiService) {
	s = stubHiService{}
	return s
}
func (hi *stubHiService) Foo(ctx context.Context, mp []map[[]string][]*blla.Bar) (b0 []blla.Bar) {
	// Implement your business logic here
	return b0
}
