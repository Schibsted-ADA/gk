package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kujtimiihoxha/gk/test_dir/hello/pkg/service"
)

type Endpoints struct {
	WorldEndpoint endpoint.Endpoint
}
type WorldRequest struct {
	S string
}
type WorldResponse struct {
	Rs  string
	Err error
}

func New(svc service.HelloService) (ep Endpoints) {
	ep.WorldEndpoint = MakeWorldEndpoint(svc)
	return ep
}

func MakeWorldEndpoint(svc service.HelloService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WorldRequest)
		rs, err := svc.World(ctx, req.S)
		return WorldResponse{Rs: rs, Err: err}, nil
	}
}
