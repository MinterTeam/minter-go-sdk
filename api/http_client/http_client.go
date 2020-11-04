package http_client

import (
	"context"
	"errors"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client/api_service"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"github.com/go-openapi/strfmt"
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
	ctxFunc func() context.Context
	debug   bool
	logger  logger.Logger
	headers map[string][]string
}

// NewConcise returns concise HTTP client, with wrapper over go-swagger methods
func NewConcise(address string) (*Concise, error) {
	parseAddress, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	concise := &Concise{
		host:    parseAddress.Host,
		path:    parseAddress.Path,
		ssl:     parseAddress.Scheme == "https",
		timeout: 10 * time.Second,
		ctxFunc: context.Background,
		logger:  logger.StandardLogger{},
		debug:   logger.DebugEnabled(),
		headers: map[string][]string{},
	}

	return concise.setClientService(nil), nil
}

func (c *Concise) setClientService(clientService api_service.ClientService) *Concise {
	if clientService == nil {
		r := httptransport.New(c.host, c.path, []string{c.getHTTPProtocol()})
		r.SetDebug(c.debug)
		r.SetLogger(c.logger)

		r.DefaultAuthentication = clientAuthInfoWriterFunc(c.headers)
		clientService = client.New(r, nil).APIService
	}
	c.ClientService = clientService
	return c
}

func clientAuthInfoWriterFunc(headers map[string][]string) runtime.ClientAuthInfoWriterFunc {
	return func(req runtime.ClientRequest, reg strfmt.Registry) error {
		for key, value := range headers {
			err := req.SetHeaderParam(key, value...)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// WithTimeout returns copy of Concise with timeout.
func (c *Concise) WithTimeout(timeout time.Duration) *Concise {
	concise := &Concise{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  c.logger,
		headers: c.headers,
	}
	return concise.setClientService(nil)
}

// WithHeaders returns copy of Concise with custom headers.
func (c *Concise) WithHeaders(headers map[string][]string) *Concise {
	concise := &Concise{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  c.logger,
		headers: headers,
	}
	return concise.setClientService(nil)
}

// WithDebug returns copy of Concise with debug.
func (c *Concise) WithDebug(debug bool) *Concise {
	concise := &Concise{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   debug,
		logger:  c.logger,
		headers: c.headers,
	}
	return concise.setClientService(nil)
}

// WithLogger returns copy of Concise with custom logger.
func (c *Concise) WithLogger(logger logger.Logger) *Concise {
	concise := &Concise{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  logger,
		headers: c.headers,
	}
	return concise.setClientService(nil)
}

// WithContextFunc returns new Concise client with new context
// Example:
// 		timeout := func(c context.Context) func() context.Context {
//			return func() context.Context {
//				ctx, _ := context.WithTimeout(c, 10*time.Second)
//				return ctx
//			}
//		}
func (c *Concise) WithContextFunc(contextFunc func(context.Context) func() context.Context) *Concise {
	concise := &Concise{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: contextFunc(c.ctxFunc()),
		debug:   c.debug,
		logger:  c.logger,
		headers: c.headers,
	}
	return concise.setClientService(nil)
}

// CoinID returns ID of coin symbol.
func (c *Concise) CoinID(symbol string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.CoinInfo(symbol, optionalHeight...)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

// Halts returns the candidate votes for stopping the network at block.
func (c *Concise) Halts(height uint64) (*models.HaltsResponse, error) {
	res, err := c.ClientService.Halts(api_service.NewHaltsParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Genesis returns genesis file.
func (c *Concise) Genesis() (*models.GenesisResponse, error) {
	res, err := c.ClientService.Genesis(api_service.NewGenesisParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Status returns node status including pubkey, latest block.
func (c *Concise) Status() (*models.StatusResponse, error) {
	res, err := c.ClientService.Status(api_service.NewStatusParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()))
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
	res, err := c.AddressExtended(address, true, optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Addresses returns list of addresses.
func (c *Concise) Addresses(addresses []string, optionalHeight ...uint64) (*models.AddressesResponse, error) {
	res, err := c.AddressesExtended(addresses, true, optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AddressExtended returns coins list with bipValue, balance, delegated and transaction count of an address.
func (c *Concise) AddressExtended(address string, delegated bool, optionalHeight ...uint64) (*models.AddressResponse, error) {
	res, err := c.ClientService.Address(api_service.NewAddressParamsWithTimeout(c.timeout).WithAddress(address).WithHeight(optionalInt(optionalHeight)).WithDelegated(&delegated).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// AddressesExtended returns list of addresses with bipValue and delegated.
func (c *Concise) AddressesExtended(addresses []string, delegated bool, optionalHeight ...uint64) (*models.AddressesResponse, error) {
	res, err := c.ClientService.Addresses(api_service.NewAddressesParamsWithTimeout(c.timeout).WithAddresses(addresses).WithHeight(optionalInt(optionalHeight)).WithDelegated(&delegated).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Block returns block data at given height.
func (c *Concise) Block(height uint64) (*models.BlockResponse, error) {
	res, err := c.ClientService.Block(api_service.NewBlockParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Candidate returns candidate’s res by provided public_key.
func (c *Concise) Candidate(publicKey string, optionalHeight ...uint64) (*models.CandidateResponse, error) {
	res, err := c.CandidateExtended(publicKey, false, optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Candidates returns list of candidates.
func (c *Concise) Candidates(includeStakes bool, status string, optionalHeight ...uint64) (*models.CandidatesResponse, error) {
	if status == "" {
		status = "all"
	}
	res, err := c.CandidatesExtended(includeStakes, true, "", optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CandidateExtended returns candidate’s res by provided public_key.
func (c *Concise) CandidateExtended(publicKey string, notShowStakes bool, optionalHeight ...uint64) (*models.CandidateResponse, error) {
	res, err := c.ClientService.Candidate(api_service.NewCandidateParamsWithTimeout(c.timeout).WithNotShowStakes(&notShowStakes).WithPublicKey(publicKey).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CandidatesExtended returns list of candidates.
func (c *Concise) CandidatesExtended(includeStakes, notShowStakes bool, status string, optionalHeight ...uint64) (*models.CandidatesResponse, error) {
	if status == "" {
		status = "all"
	}
	res, err := c.ClientService.Candidates(api_service.NewCandidatesParamsWithTimeout(c.timeout).WithIncludeStakes(&includeStakes).WithNotShowStakes(&notShowStakes).WithStatus(&status).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CoinInfoByID returns information about coin ID.
func (c *Concise) CoinInfoByID(id uint64, optionalHeight ...uint64) (*models.CoinInfoResponse, error) {
	res, err := c.ClientService.CoinInfoByID(api_service.NewCoinInfoByIDParamsWithTimeout(c.timeout).WithID(strconv.Itoa(int(id))).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CoinInfo returns information about coin symbol.
func (c *Concise) CoinInfo(symbol string, optionalHeight ...uint64) (*models.CoinInfoResponse, error) {
	res, err := c.ClientService.CoinInfo(api_service.NewCoinInfoParamsWithTimeout(c.timeout).WithSymbol(symbol).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolBuy return estimate of buy coin transaction.
func (c *Concise) EstimateCoinSymbolBuy(coinToSell, coinToBuy, valueToBuy string, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithTimeout(c.timeout).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSell return estimate of sell coin transaction.
func (c *Concise) EstimateCoinSymbolSell(coinToBuy, coinToSell, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParamsWithTimeout(c.timeout).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSellAll return estimate of sell all coin transaction.
func (c *Concise) EstimateCoinSymbolSellAll(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParamsWithTimeout(c.timeout).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDBuy return estimate of buy coin transaction.
func (c *Concise) EstimateCoinIDBuy(coinToSell, coinToBuy uint64, valueToBuy string, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithTimeout(c.timeout).WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSell return estimate of sell coin transaction.
func (c *Concise) EstimateCoinIDSell(coinToBuy, coinToSell uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParamsWithTimeout(c.timeout).WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSellAll return estimate of sell all coin transaction.
func (c *Concise) EstimateCoinIDSellAll(coinToBuy, coinToSell, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParamsWithTimeout(c.timeout).WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateTxCommission return estimate of encoding transaction.
func (c *Concise) EstimateTxCommission(tx string, optionalHeight ...uint64) (*models.EstimateTxCommissionResponse, error) {
	res, err := c.ClientService.EstimateTxCommission(api_service.NewEstimateTxCommissionParamsWithTimeout(c.timeout).WithTx(tx).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Events returns events at given height.
func (c *Concise) Events(height uint64, search ...string) (*models.EventsResponse, error) {
	res, err := c.ClientService.Events(api_service.NewEventsParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithSearch(search).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MaxGasPrice returns current max gas.
func (c *Concise) MaxGasPrice(optionalHeight ...uint64) (*models.MaxGasPriceResponse, error) {
	res, err := c.ClientService.MaxGasPrice(api_service.NewMaxGasPriceParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MinGasPrice returns current min gas price.
func (c *Concise) MinGasPrice() (*models.MinGasPriceResponse, error) {
	res, err := c.ClientService.MinGasPrice(api_service.NewMinGasPriceParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MissedBlocks returns missed blocks by validator public key.
func (c *Concise) MissedBlocks(publicKey string, optionalHeight ...uint64) (*models.MissedBlocksResponse, error) {
	res, err := c.ClientService.MissedBlocks(api_service.NewMissedBlocksParamsWithTimeout(c.timeout).WithPublicKey(publicKey).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// NetInfo returns network res
func (c *Concise) NetInfo() (*models.NetInfoResponse, error) {
	res, err := c.ClientService.NetInfo(api_service.NewNetInfoParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SendTransaction returns the result of sending signed tx.
// To ensure that transaction was successfully committed to the blockchain,
// you need to find the transaction by the hash and ensure that the status code equals to 0.
func (c *Concise) SendTransaction(tx string) (*models.SendTransactionResponse, error) {
	res, err := c.ClientService.SendTransaction(api_service.NewSendTransactionParamsWithTimeout(c.timeout).WithTx(tx).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Transaction returns transaction res.
func (c *Concise) Transaction(hash string) (*models.TransactionResponse, error) {
	res, err := c.ClientService.Transaction(api_service.NewTransactionParamsWithTimeout(c.timeout).WithHash(hash).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Transactions returns transactions by query.
func (c *Concise) Transactions(query string, page, perPage int32) (*models.TransactionsResponse, error) {
	res, err := c.ClientService.Transactions(api_service.NewTransactionsParamsWithTimeout(c.timeout).WithQuery(query).WithPage(&page).WithPerPage(&perPage).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// UnconfirmedTxs returns unconfirmed transactions.
func (c *Concise) UnconfirmedTxs(limit int32) (*models.UnconfirmedTxsResponse, error) {
	res, err := c.ClientService.UnconfirmedTxs(api_service.NewUnconfirmedTxsParamsWithTimeout(c.timeout).WithLimit(&limit).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Validators returns list of active validators.
func (c *Concise) Validators(optionalHeight ...uint64) (*models.ValidatorsResponse, error) {
	res, err := c.ClientService.Validators(api_service.NewValidatorsParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// WaitList returns the list of address stakes in waitlist.
func (c *Concise) WaitList(publicKey, address string, optionalHeight ...uint64) (*models.WaitListResponse, error) {
	res, err := c.ClientService.WaitList(api_service.NewWaitListParamsWithTimeout(c.timeout).WithAddress(address).WithPublicKey(&publicKey).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()))
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

	res, err := c.ClientService.Frozen(api_service.NewFrozenParamsWithTimeout(c.timeout).WithAddress(address).WithCoinID(optionalInt(optionalCoinID)).WithContext(c.ctxFunc()))
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
	scheme := c.getWSProtocol()
	u := url.URL{Scheme: scheme, Host: c.host, Path: c.path + "/subscribe", RawQuery: "query=" + url.QueryEscape(query)}

	var err error
	subClient.conn, _, err = websocket.DefaultDialer.DialContext(subClient.ctx, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return subClient, nil
}

func (c *Concise) getWSProtocol() string {
	if c.ssl {
		return "wss"
	}
	return "ws"
}

func (c *Concise) getHTTPProtocol() string {
	if c.ssl {
		return "https"
	}
	return "http"
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
	if len(height) == 0 {
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
