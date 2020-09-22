package http_client

import (
	"context"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client/api_service"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// Client HTTP
type Client struct {
	api_service.ClientService
	host string
	path string
	ssl  bool
}

// New returns HTTP client api_service.ClientService
func New(address string) (*Client, error) {
	parseAddress, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return &Client{
		ClientService: client.NewHTTPClientWithConfig(nil,
			client.DefaultTransportConfig().
				WithHost(parseAddress.Host).
				WithBasePath(parseAddress.Path).
				WithSchemes([]string{parseAddress.Scheme}),
		).APIService,
		host: parseAddress.Host,
		path: parseAddress.Path,
		ssl:  parseAddress.Scheme == "https",
	}, nil
}

type defaultError interface {
	Code() int
	GetPayload() *models.ErrorBody
}

// ErrorBody returns error as API model
func ErrorBody(err error) (int, *models.ErrorBody, error) {
	if err == nil {
		return http.StatusOK, nil, nil
	}

	if errorBody, ok := err.(defaultError); ok {
		return errorBody.Code(), errorBody.GetPayload(), nil
	}

	return 0, nil, err
}

// CoinID returns ID of coin symbol.
func (c *Client) CoinID(symbol string) (uint32, error) {
	info, err := c.CoinInfo(api_service.NewCoinInfoParams().WithSymbol(symbol))
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(info.GetPayload().ID)
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}

// SubscriberClient is subscriber
type SubscriberClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	conn   *websocket.Conn
	closed bool
}

// Close closes connection.
func (s *SubscriberClient) Close() error {
	if s.closed {
		return nil
	}
	s.cancel()
	s.closed = true
	return s.conn.Close()
}

// Recv returns message SubscribeOKBody.
func (s *SubscriberClient) Recv() (*api_service.SubscribeOKBody, error) {
	if s.closed {
		return nil, io.EOF
	}

	select {
	case <-s.ctx.Done():
		s.Close()
		return nil, io.EOF
	default:
		var recv api_service.SubscribeOKBody
		err := s.conn.ReadJSON(&recv)
		if websocket.IsCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
			s.Close()
			err = io.EOF
		}

		return &recv, err
	}
}

// Subscribe returns a subscription for events by query.
func (c *Client) Subscribe(ctx context.Context, query string) (*SubscriberClient, error) {

	subClient := &SubscriberClient{}

	subClient.ctx, subClient.cancel = context.WithCancel(ctx)
	scheme := "ws"
	if c.ssl {
		scheme += "s"
	}
	u := url.URL{Scheme: scheme, Host: c.host, Path: c.path + "/subscribe", RawQuery: "query=" + url.QueryEscape(query)}

	var err error
	subClient.conn, _, err = websocket.DefaultDialer.DialContext(subClient.ctx, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return subClient, nil
}
