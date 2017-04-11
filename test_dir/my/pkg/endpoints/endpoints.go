package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kujtimiihoxha/gk/test_dir/my/pkg/service"
)

type Endpoints struct {
	FooEndpoint endpoint.Endpoint
}
type FooRequest struct {
	A []string
}
type FooResponse struct {
	M0 map[string]string
	E1 error
}

func New(svc service.MyService) (ep Endpoints) {
	ep.FooEndpoint = MakeFooEndpoint(svc)
	return ep
}

func MakeFooEndpoint(svc service.MyService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		m0, e1 := svc.Foo(ctx, req.A)
		return FooResponse{M0: m0, E1: e1}, nil
	}
}
