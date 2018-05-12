package meld

import (
	"context"
	"encoding/json"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createAWSEC2Handler := kithttp.NewServer(
		MakeCreateAWSEC2Endpoint(s),
		decodeCreateAWSEC2Request,
		encodeResponse,
		opts...,
	)

	readAWSEC2Handler := kithttp.NewServer(
		MakeReadAWSEC2Endpoint(s),
		decodeReadAWSEC2Request,
		encodeResponse,
		opts...,
	)

	updateAWSEC2Handler := kithttp.NewServer(
		MakeUpdateAWSEC2Endpoint(s),
		decodeUpdateAWSEC2Request,
		encodeResponse,
		opts...,
	)

	deleteAWSEC2Handler := kithttp.NewServer(
		MakeDeleteAWSEC2Endpoint(s),
		decodeDeleteAWSEC2Request,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/meld/v1/aws/ec2", createAWSEC2Handler).Methods("POST")
	r.Handle("/meld/v1/aws/ec2/{id}", readAWSEC2Handler).Methods("GET")
	r.Handle("/meld/v1/aws/ec2/{id}", updateAWSEC2Handler).Methods("PUT")
	r.Handle("/meld/v1/aws/ec2/{id}", deleteAWSEC2Handler).Methods("DELETE")

	return r
}

func decodeCreateAWSEC2Request(_ context.Context, r *http.Request) (interface{}, error) {
	return createAWSEC2Request{}, nil
}

func decodeReadAWSEC2Request(_ context.Context, r *http.Request) (interface{}, error) {
	return readAWSEC2Request{}, nil
}

func decodeUpdateAWSEC2Request(_ context.Context, r *http.Request) (interface{}, error) {
	return updateAWSEC2Request{}, nil
}

func decodeDeleteAWSEC2Request(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteAWSEC2Request{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
