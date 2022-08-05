package account_service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeAccountExistsEndpoint(svc AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(AccountExistsRequestDTO)
		if !ok {
			return AccountExistsResponseDTO{Exists: false, Error: errors.New("invalid request").Error()}, nil
		}

		exists, err := svc.Exists(req.Id)
		if err != nil {
			return AccountExistsResponseDTO{Exists: false, Error: err.Error()}, nil
		}

		return AccountExistsResponseDTO{Exists: exists, Error: ""}, nil
	}
}
