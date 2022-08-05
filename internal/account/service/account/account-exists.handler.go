package account_service

import (
	"context"
	"encoding/json"
	"net/http"
	transportUtils "refound/internal/account/utils/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func decodeAccountExistsRequest(_ context.Context, req *http.Request) (interface{}, error) {
	var request AccountExistsRequestDTO
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeAccountExistsHandler(svc AccountService) *httptransport.Server {
	return httptransport.NewServer(makeAccountExistsEndpoint(svc), decodeAccountExistsRequest, transportUtils.EncodeResponse)
}
