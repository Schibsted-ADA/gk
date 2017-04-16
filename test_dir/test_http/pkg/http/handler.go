package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kujtimiihoxha/gk/test_dir/test_http/pkg/endpoints"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/foo", httptransport.NewServer(
		endpoints.FooEndpoint,
		DecodeFooRequest,
		EncodeFooResponse,
	))
	m.Handle("/bar", httptransport.NewServer(
		endpoints.BarEndpoint,
		DecodeBarRequest,
		EncodeBarResponse,
	))
	return m
}

func DecodeFooRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.FooRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err

}

func EncodeFooResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(response)
	return err

}

func DecodeBarRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.BarRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err

}

func EncodeBarResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(response)
	return err

}
