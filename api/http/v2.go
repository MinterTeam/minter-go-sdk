package http

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api/http/client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http/client/api_service"
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
