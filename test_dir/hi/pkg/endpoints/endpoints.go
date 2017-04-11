package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kujtimiihoxha/gk/test_dir/blla"
	"github.com/kujtimiihoxha/gk/test_dir/hi/pkg/service"
)

type Endpoints struct {
	FooEndpoint endpoint.Endpoint
}
type FooRequest struct {
	Mp []map[[]string][]*blla.Bar
}
type FooResponse struct {
	B0 []blla.Bar
}

func New(svc service.HiService) (ep Endpoints) {
	ep.FooEndpoint = MakeFooEndpoint(svc)
	return ep
}

func MakeFooEndpoint(svc service.HiService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		b0 := svc.Foo(ctx, req.Mp)
		return FooResponse{B0: b0}, nil
	}
}
