package meld

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type createAWSEC2Request struct {
}

type createAWSEC2Response struct {
	ID  string `json:"cluster_id,omitempty"`
	Err error  `json:"error,omitempty"`
}

func (r createAWSEC2Response) error() error { return r.Err }

func MakeCreateAWSEC2Endpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(createAWSEC2Request)
		id, err := s.CreateAWSEC2()
		return createAWSEC2Response{ID: id, Err: err}, nil
	}
}

type readAWSEC2Request struct {
}

type readAWSEC2Response struct {
	ID  string `json:"cluster_id,omitempty"`
	Err error  `json:"error,omitempty"`
}

func (r readAWSEC2Response) error() error { return r.Err }

func MakeReadAWSEC2Endpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(readAWSEC2Request)
		id, err := s.ReadAWSEC2()
		return readAWSEC2Response{ID: id, Err: err}, nil
	}
}

type updateAWSEC2Request struct {
}

type updateAWSEC2Response struct {
	ID  string `json:"cluster_id,omitempty"`
	Err error  `json:"error,omitempty"`
}

func (r updateAWSEC2Response) error() error { return r.Err }

func MakeUpdateAWSEC2Endpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(updateAWSEC2Request)
		id, err := s.UpdateAWSEC2()
		return updateAWSEC2Response{ID: id, Err: err}, nil
	}
}

type deleteAWSEC2Request struct {
}

type deleteAWSEC2Response struct {
	ID  string `json:"cluster_id,omitempty"`
	Err error  `json:"error,omitempty"`
}

func (r deleteAWSEC2Response) error() error { return r.Err }

func MakeDeleteAWSEC2Endpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(deleteAWSEC2Request)
		id, err := s.DeleteAWSEC2()
		return deleteAWSEC2Response{ID: id, Err: err}, nil
	}
}
