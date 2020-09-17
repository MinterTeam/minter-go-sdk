package http_client

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client/api_service"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"net/url"
)

func New(address string) (api_service.ClientService, error) {
	parseAddress, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return client.NewHTTPClientWithConfig(nil,
		client.DefaultTransportConfig().
			WithHost(parseAddress.Host).
			WithBasePath(parseAddress.Path).
			WithSchemes([]string{parseAddress.Scheme}),
	).APIService, nil
}

type DefaultError interface {
	Code() int
	GetPayload() *models.APIPbErrorBody
}

func ErrorBody(err error) (int, *models.APIPbErrorBody, error) {
	if errorBody, ok := err.(DefaultError); ok {
		return errorBody.Code(), errorBody.GetPayload(), nil
	}

	return 0, nil, err
}
