package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kujtimiihoxha/gk/test_dir/hello/pkg/endpoints"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/world", httptransport.NewServer(
		endpoints.WorldEndpoint,
		DecodeWorldRequest,
		EncodeWorldResponse,
	))
	return m
}

func DecodeWorldRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.WorldRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err

}

func EncodeWorldResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	json.NewEncoder(w).Encode(response)
	return err

}
