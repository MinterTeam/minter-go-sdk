package v2

import (
	"github.com/MinterTeam/minter-go-sdk/api/v2/client"
	"github.com/MinterTeam/minter-go-sdk/api/v2/client/api_service"
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
