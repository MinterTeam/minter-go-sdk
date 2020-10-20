package http_client

import (
	"context"
	"errors"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client/api_service"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Concise http concise HTTP client, with wrapper over go-swagger Client methods
type Concise struct {
	api_service.ClientService
	host    string
	path    string
	ssl     bool
	timeout time.Duration
}

// NewConcise returns concise HTTP client, with wrapper over go-swagger methods
func NewConcise(address string) (*Concise, error) {
	parseAddress, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return &Concise{
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

// WithTimeout returns copy of Concise with timeout.
func (c *Concise) WithTimeout(timeout time.Duration) *Concise {
	return &Concise{
		ClientService: c.ClientService,
		host:          c.host,
		path:          c.path,
		ssl:           c.ssl,
		timeout:       timeout,
	}
}

// CoinID returns ID of coin symbol.
func (c *Concise) CoinID(symbol string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.ClientService.CoinInfo(api_service.NewCoinInfoParams().WithSymbol(symbol).WithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)))
	if err != nil {
		return 0, err
	}

	return res.GetPayload().ID, nil
}

// Halts returns the candidate votes for stopping the network at block.
func (c *Concise) Halts(height uint64) (*models.HaltsResponse, error) {
	res, err := c.ClientService.Halts(api_service.NewHaltsParams().WithHeight(strconv.Itoa(int(height))).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Genesis returns genesis file.
func (c *Concise) Genesis() (*models.GenesisResponse, error) {
	res, err := c.ClientService.Genesis(api_service.NewGenesisParams().WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Status returns node status including pubkey, latest block.
func (c *Concise) Status() (*models.StatusResponse, error) {
	res, err := c.ClientService.Status(api_service.NewStatusParams().WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Nonce returns next transaction number (nonce) of an address.
func (c *Concise) Nonce(address string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.Address(address, optionalHeight...)
	if err != nil {
		return 0, err
	}

	return res.TransactionCount + 1, nil
}

// Address returns coins list, balance and transaction count of an address.
func (c *Concise) Address(address string, optionalHeight ...uint64) (*models.AddressResponse, error) {
	res, err := c.ClientService.Address(api_service.NewAddressParams().WithAddress(address).WithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Addresses returns list of addresses.
func (c *Concise) Addresses(addresses []string, optionalHeight ...uint64) (*models.AddressesResponse, error) {
	res, err := c.ClientService.Addresses(api_service.NewAddressesParams().WithAddresses(addresses).WithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Block returns block data at given height.
func (c *Concise) Block(height uint64) (*models.BlockResponse, error) {
	res, err := c.ClientService.Block(api_service.NewBlockParams().WithHeight(strconv.Itoa(int(height))).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Candidate returns candidate’s res by provided public_key. It will respond with 404 code if candidate is not found.
func (c *Concise) Candidate(publicKey string, optionalHeight ...uint64) (*models.CandidateResponse, error) {
	res, err := c.ClientService.Candidate(api_service.NewCandidateParams().WithPublicKey(publicKey).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Candidates returns list of candidates.
func (c *Concise) Candidates(includeStakes bool, status string, optionalHeight ...uint64) (*models.CandidatesResponse, error) {
	if status == "" {
		status = "all"
	}
	res, err := c.ClientService.Candidates(api_service.NewCandidatesParams().WithIncludeStakes(&includeStakes).WithStatus(&status).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CoinInfoByID returns information about coin ID.
func (c *Concise) CoinInfoByID(id uint64, optionalHeight ...uint64) (*models.CoinInfoResponse, error) {
	res, err := c.ClientService.CoinInfoByID(api_service.NewCoinInfoByIDParams().WithID(strconv.Itoa(int(id))).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CoinInfo returns information about coin symbol.
func (c *Concise) CoinInfo(symbol string, optionalHeight ...uint64) (*models.CoinInfoResponse, error) {
	res, err := c.ClientService.CoinInfo(api_service.NewCoinInfoParams().WithSymbol(symbol).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolBuy return estimate of buy coin transaction.
func (c *Concise) EstimateCoinSymbolBuy(coinToSell, coinToBuy, valueToBuy string, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParams().WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSell return estimate of sell coin transaction.
func (c *Concise) EstimateCoinSymbolSell(coinToBuy, coinToSell, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParams().WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSellAll return estimate of sell all coin transaction.
func (c *Concise) EstimateCoinSymbolSellAll(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParams().WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDBuy return estimate of buy coin transaction.
func (c *Concise) EstimateCoinIDBuy(coinToSell, coinToBuy uint64, valueToBuy string, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParams().WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSell return estimate of sell coin transaction.
func (c *Concise) EstimateCoinIDSell(coinToBuy, coinToSell uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParams().WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSellAll return estimate of sell all coin transaction.
func (c *Concise) EstimateCoinIDSellAll(coinToBuy, coinToSell, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParams().WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateTxCommission return estimate of encoding transaction.
func (c *Concise) EstimateTxCommission(tx string, optionalHeight ...uint64) (*models.EstimateTxCommissionResponse, error) {
	res, err := c.ClientService.EstimateTxCommission(api_service.NewEstimateTxCommissionParams().WithTx(tx).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Events returns events at given height.
func (c *Concise) Events(height uint64, search ...string) (*models.EventsResponse, error) {
	res, err := c.ClientService.Events(api_service.NewEventsParams().WithHeight(strconv.Itoa(int(height))).WithTimeout(c.timeout).WithSearch(search))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MaxGasPrice returns current max gas.
func (c *Concise) MaxGasPrice(optionalHeight ...uint64) (*models.MaxGasPriceResponse, error) {
	res, err := c.ClientService.MaxGasPrice(api_service.NewMaxGasPriceParams().WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MinGasPrice returns current min gas price.
func (c *Concise) MinGasPrice() (*models.MinGasPriceResponse, error) {
	res, err := c.ClientService.MinGasPrice(api_service.NewMinGasPriceParams().WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MissedBlocks returns missed blocks by validator public key.
func (c *Concise) MissedBlocks(publicKey string, optionalHeight ...uint64) (*models.MissedBlocksResponse, error) {
	res, err := c.ClientService.MissedBlocks(api_service.NewMissedBlocksParams().WithPublicKey(publicKey).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// NetInfo returns network res
func (c *Concise) NetInfo() (*models.NetInfoResponse, error) {
	res, err := c.ClientService.NetInfo(api_service.NewNetInfoParams().WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SendTransaction returns the result of sending signed tx.
// To ensure that transaction was successfully committed to the blockchain,
// you need to find the transaction by the hash and ensure that the status code equals to 0.
func (c *Concise) SendTransaction(tx string) (*models.SendTransactionResponse, error) {
	res, err := c.ClientService.SendTransaction(api_service.NewSendTransactionParams().WithTx(tx).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Transaction returns transaction res.
func (c *Concise) Transaction(hash string) (*models.TransactionResponse, error) {
	res, err := c.ClientService.Transaction(api_service.NewTransactionParams().WithHash(hash).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Transactions returns transactions by query.
func (c *Concise) Transactions(query string, page, perPage int32) (*models.TransactionsResponse, error) {
	res, err := c.ClientService.Transactions(api_service.NewTransactionsParams().WithQuery(query).WithPage(&page).WithPerPage(&perPage).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// UnconfirmedTxs returns unconfirmed transactions.
func (c *Concise) UnconfirmedTxs(limit int32) (*models.UnconfirmedTxsResponse, error) {
	res, err := c.ClientService.UnconfirmedTxs(api_service.NewUnconfirmedTxsParams().WithLimit(&limit).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Validators returns list of active validators.
func (c *Concise) Validators(optionalHeight ...uint64) (*models.ValidatorsResponse, error) {
	res, err := c.ClientService.Validators(api_service.NewValidatorsParams().WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// WaitList returns the list of address stakes in waitlist.
func (c *Concise) WaitList(publicKey, address string, optionalHeight ...uint64) (*models.WaitListResponse, error) {
	res, err := c.ClientService.WaitList(api_service.NewWaitListParams().WithAddress(address).WithPublicKey(&publicKey).WithHeight(optionalInt(optionalHeight)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Frozen returns frozen balance.
func (c *Concise) Frozen(address string, optionalCoinID ...uint64) (*models.FrozenResponse, error) {
	if len(optionalCoinID) > 1 {
		return nil, errors.New("CoinID needed single value") // todo: change message
	}

	res, err := c.ClientService.Frozen(api_service.NewFrozenParams().WithAddress(address).WithCoinID(optionalInt(optionalCoinID)).WithTimeout(c.timeout))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SubscriberClient is subscriber
type SubscriberClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	conn   *websocket.Conn
	closed bool
}

// CloseSend closes the send direction of the stream.
func (s *SubscriberClient) CloseSend() error {
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
		_ = s.CloseSend()
		return nil, s.ctx.Err()
	default:
		var recv api_service.SubscribeOKBody
		err := s.conn.ReadJSON(&recv)
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			_ = s.CloseSend()
			err = io.EOF
		}
		if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
			_ = s.CloseSend()
			err = context.Canceled
		}

		return &recv, err
	}
}

// Subscribe returns a subscription for events by query.
func (c *Concise) Subscribe(ctx context.Context, query string) (*SubscriberClient, error) {

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

func optionalInt(height []uint64) *uint64 {
	if len(height) == 1 {
		return nil
	}
	return &height[0]
}

// Client is go-swagger HTTP client
type Client struct {
	api_service.ClientService
}

// New returns HTTP client from go-swagger generator
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
	}, nil
}

// CoinID returns ID of coin symbol.
func (c *Client) CoinID(symbol string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.ClientService.CoinInfo(api_service.NewCoinInfoParams().WithSymbol(symbol).WithHeight(optionalInt(optionalHeight)))
	if err != nil {
		return 0, err
	}

	return res.GetPayload().ID, nil
}

// Nonce returns next transaction number (nonce) of an address.
func (c *Client) Nonce(address string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.ClientService.Address(api_service.NewAddressParams().WithAddress(address).WithHeight(optionalInt(optionalHeight)))
	if err != nil {
		return 0, err
	}

	return res.Payload.TransactionCount + 1, nil
}
