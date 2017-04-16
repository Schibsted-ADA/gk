package grpc

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/kujtimiihoxha/gk/test_dir/test_grpc/pkg/endpoints"
	"github.com/kujtimiihoxha/gk/test_dir/test_grpc/pkg/grpc/pb"
	oldcontext "golang.org/x/net/context"
)

type grpcServer struct {
	foo grpctransport.Handler
	bar grpctransport.Handler
}

func MakeGRPCServer(endpoints endpoints.Endpoints) (req pb.TestGrpcServer) {
	req = &grpcServer{
		foo: grpctransport.NewServer(
			endpoints.FooEndpoint,
			DecodeGRPCFooRequest,
			EncodeGRPCFooResponse,
		),

		bar: grpctransport.NewServer(
			endpoints.BarEndpoint,
			DecodeGRPCBarRequest,
			EncodeGRPCBarResponse,
		),
	}
	return req
}

func DecodeGRPCFooRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Foo' Decoder is not impelement")
	return req, err
}

func EncodeGRPCFooResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Foo' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Foo(ctx oldcontext.Context, req *pb.FooRequest) (rep *pb.FooReply, err error) {
	_, rp, err := s.foo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.FooReply)
	return rep, err
}

func DecodeGRPCBarRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Bar' Decoder is not impelement")
	return req, err
}

func EncodeGRPCBarResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Bar' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Bar(ctx oldcontext.Context, req *pb.BarRequest) (rep *pb.BarReply, err error) {
	_, rp, err := s.bar.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.BarReply)
	return rep, err
}
