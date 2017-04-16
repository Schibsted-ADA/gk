package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kujtimiihoxha/gk/test_dir/test_grpc/pkg/service"
)

type Endpoints struct {
	FooEndpoint endpoint.Endpoint
	BarEndpoint endpoint.Endpoint
}
type FooRequest struct {
	S string
}
type FooResponse struct {
	S0 string
	E1 error
}
type BarRequest struct {
	I int
}
type BarResponse struct {
	I0 int
	E1 error
}

func New(svc service.TestGrpcService) (ep Endpoints) {
	ep.FooEndpoint = MakeFooEndpoint(svc)
	ep.BarEndpoint = MakeBarEndpoint(svc)
	return ep
}

func MakeFooEndpoint(svc service.TestGrpcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		s0, e1 := svc.Foo(ctx, req.S)
		return FooResponse{S0: s0, E1: e1}, nil
	}
}

func MakeBarEndpoint(svc service.TestGrpcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BarRequest)
		i0, e1 := svc.Bar(ctx, req.I)
		return BarResponse{I0: i0, E1: e1}, nil
	}
}
