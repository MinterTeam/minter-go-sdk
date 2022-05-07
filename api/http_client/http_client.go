package http_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/swag"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/client/api_service"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"github.com/go-openapi/strfmt"
	"github.com/gorilla/websocket"
)

// Client http apiClient HTTP client, with wrapper over go-swagger SwagClient methods
type Client struct {
	api_service.ClientService
	host    string
	path    string
	ssl     bool
	timeout time.Duration
	ctxFunc func() context.Context
	debug   bool
	logger  logger.Logger
	headers map[string][]string
	opts    []api_service.ClientOption
}

// New returns apiClient HTTP client, with wrapper over go-swagger methods
func New(address string) (*Client, error) {
	parseAddress, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	c := &Client{
		host:    parseAddress.Host,
		path:    parseAddress.Path,
		ssl:     parseAddress.Scheme == "https",
		timeout: 10 * time.Second,
		ctxFunc: context.Background,
		logger:  logger.StandardLogger{},
		debug:   logger.DebugEnabled(),
		headers: map[string][]string{},
	}

	return c.setClientService(nil), nil
}

func (c *Client) setClientService(clientService api_service.ClientService) *Client {
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

// WithCallOption returns new Client with additional api_service.ClientOption
func (c *Client) WithCallOption(opts ...api_service.ClientOption) *Client {
	return &Client{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  c.logger,
		headers: c.headers,
		opts:    append(c.opts, opts...),
	}
}

// WithTimeout returns copy of Client with timeout.
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	apiClient := &Client{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  c.logger,
		headers: c.headers,
		opts:    c.opts,
	}
	return apiClient.setClientService(nil)
}

// WithHeaders returns copy of Client with custom headers.
func (c *Client) WithHeaders(headers map[string][]string) *Client {
	apiClient := &Client{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  c.logger,
		headers: headers,
		opts:    c.opts,
	}
	return apiClient.setClientService(nil)
}

// WithDebug returns copy of Client with debug.
func (c *Client) WithDebug(debug bool) *Client {
	apiClient := &Client{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   debug,
		logger:  c.logger,
		headers: c.headers,
		opts:    c.opts,
	}
	return apiClient.setClientService(nil)
}

// WithLogger returns copy of Client with custom logger.
func (c *Client) WithLogger(logger logger.Logger) *Client {
	apiClient := &Client{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: c.ctxFunc,
		debug:   c.debug,
		logger:  logger,
		headers: c.headers,
		opts:    c.opts,
	}
	return apiClient.setClientService(nil)
}

// WithContextFunc returns new Client client with new context
// Example:
// 		timeout := func(c context.Context) func() context.Context {
//			return func() context.Context {
//				ctx, _ := context.WithTimeout(c, 10*time.Second)
//				return ctx
//			}
//		}
func (c *Client) WithContextFunc(contextFunc func(context.Context) func() context.Context) *Client {
	apiClient := &Client{
		host:    c.host,
		path:    c.path,
		ssl:     c.ssl,
		timeout: c.timeout,
		ctxFunc: contextFunc(c.ctxFunc()),
		debug:   c.debug,
		logger:  c.logger,
		headers: c.headers,
		opts:    c.opts,
	}
	return apiClient.setClientService(nil)
}

// CoinID returns ID of coin symbol.
func (c *Client) CoinID(symbol string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.CoinInfo(symbol, optionalHeight...)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

// Halts returns the candidate votes for stopping the network at block.
func (c *Client) Halts(height uint64) (*models.HaltsResponse, error) {
	res, err := c.ClientService.Halts(api_service.NewHaltsParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Genesis returns genesis file.
func (c *Client) Genesis() (*models.GenesisResponse, error) {
	res, err := c.ClientService.Genesis(api_service.NewGenesisParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Status returns node status including pubkey, latest block.
func (c *Client) Status() (*models.StatusResponse, error) {
	res, err := c.ClientService.Status(api_service.NewStatusParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Nonce returns next transaction number (nonce) of an address.
func (c *Client) Nonce(address string, optionalHeight ...uint64) (uint64, error) {
	res, err := c.Address(address, optionalHeight...)
	if err != nil {
		return 0, err
	}

	return res.TransactionCount + 1, nil
}

// Address returns coins list, balance and transaction count of an address.
func (c *Client) Address(address string, optionalHeight ...uint64) (*models.AddressResponse, error) {
	res, err := c.AddressExtended(address, false, optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Addresses returns list of addresses.
func (c *Client) Addresses(addresses []string, optionalHeight ...uint64) (*models.AddressesResponse, error) {
	res, err := c.AddressesExtended(addresses, false, optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AddressExtended returns coins list with bipValue, balance, delegated and transaction count of an address.
func (c *Client) AddressExtended(address string, delegated bool, optionalHeight ...uint64) (*models.AddressResponse, error) {
	res, err := c.ClientService.Address(api_service.NewAddressParamsWithTimeout(c.timeout).WithAddress(address).WithHeight(optionalInt(optionalHeight)).WithDelegated(&delegated).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// AddressesExtended returns list of addresses with bipValue and delegated.
func (c *Client) AddressesExtended(addresses []string, delegated bool, optionalHeight ...uint64) (*models.AddressesResponse, error) {
	res, err := c.ClientService.Addresses(api_service.NewAddressesParamsWithTimeout(c.timeout).WithAddresses(addresses).WithHeight(optionalInt(optionalHeight)).WithDelegated(&delegated).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Block returns block data at given height.
func (c *Client) Block(height uint64) (*models.BlockResponse, error) {
	res, err := c.ClientService.Block(api_service.NewBlockParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Blocks returns blocks at given interval.
func (c *Client) Blocks(from, to uint64, failedTxs, events bool, fieldsBlock ...string) (*models.BlocksResponse, error) {
	res, err := c.ClientService.Blocks(api_service.NewBlocksParamsWithTimeout(c.timeout).WithFailedTxs(&failedTxs).WithFields(fieldsBlock).WithFromHeight(from).WithToHeight(to).WithEvents(&events).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// BlockExtended returns block data at given height by filtered data.
func (c *Client) BlockExtended(height uint64, failedTxs, events bool, fieldsBlock ...string) (*models.BlockResponse, error) {
	res, err := c.ClientService.Block(api_service.NewBlockParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithFailedTxs(&failedTxs).WithFields(fieldsBlock).WithEvents(&events).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Candidate returns candidate’s res by provided public_key.
func (c *Client) Candidate(publicKey string, optionalHeight ...uint64) (*models.CandidateResponse, error) {
	res, err := c.CandidateExtended(publicKey, false, optionalHeight...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Candidates returns list of candidates.
func (c *Client) Candidates(includeStakes bool, status string, optionalHeight ...uint64) (*models.CandidatesResponse, error) {
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
func (c *Client) CandidateExtended(publicKey string, notShowStakes bool, optionalHeight ...uint64) (*models.CandidateResponse, error) {
	res, err := c.ClientService.Candidate(api_service.NewCandidateParamsWithTimeout(c.timeout).WithNotShowStakes(&notShowStakes).WithPublicKey(publicKey).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CandidatesExtended returns list of candidates.
func (c *Client) CandidatesExtended(includeStakes, notShowStakes bool, status string, optionalHeight ...uint64) (*models.CandidatesResponse, error) {
	if status == "" {
		status = "all"
	}
	res, err := c.ClientService.Candidates(api_service.NewCandidatesParamsWithTimeout(c.timeout).WithIncludeStakes(&includeStakes).WithNotShowStakes(&notShowStakes).WithStatus(&status).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CoinInfoByID returns information about coin ID.
func (c *Client) CoinInfoByID(id uint64, optionalHeight ...uint64) (*models.CoinInfoResponse, error) {
	res, err := c.ClientService.CoinInfoByID(api_service.NewCoinInfoByIDParamsWithTimeout(c.timeout).WithID(strconv.Itoa(int(id))).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CoinInfo returns information about coin symbol.
func (c *Client) CoinInfo(symbol string, optionalHeight ...uint64) (*models.CoinInfoResponse, error) {
	res, err := c.ClientService.CoinInfo(api_service.NewCoinInfoParamsWithTimeout(c.timeout).WithSymbol(symbol).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinSymbolBuy(coinToSell, coinToBuy, valueToBuy string, coinCommission string, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithTimeout(c.timeout).WithCoinCommission(&coinCommission).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinSymbolSell(coinToBuy, coinToSell, valueToSell string, coinCommission string, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParamsWithTimeout(c.timeout).WithCoinCommission(&coinCommission).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSellAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinSymbolSellAll(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParamsWithTimeout(c.timeout).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinIDBuy(coinToSell, coinToBuy uint64, valueToBuy string, coinCommission uint64, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithTimeout(c.timeout).WithCoinIDCommission(&coinCommission).WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinIDSell(coinToBuy, coinToSell uint64, valueToSell string, coinCommission uint64, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParamsWithTimeout(c.timeout).WithCoinIDCommission(&coinCommission).WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSellAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinIDSellAll(coinToBuy, coinToSell, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParamsWithTimeout(c.timeout).WithCoinIDToSell(&coinToSell).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolBuyExtended return estimate of buy coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinSymbolBuyExtended(coinToSell, coinToBuy, valueToBuy string, coinCommission string, swapFrom string, route []uint64, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	var coinsRoute []string
	for _, coin := range route {
		coinsRoute = append(coinsRoute, strconv.Itoa(int(coin)))
	}
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithTimeout(c.timeout).WithCoinCommission(&coinCommission).WithCoinToSell(&coinToSell).WithCoinToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithSwapFrom(&swapFrom).WithRoute(coinsRoute).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSellExtended return estimate of sell coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinSymbolSellExtended(coinToBuy, coinToSell, valueToSell string, coinCommission string, swapFrom string, route []uint64, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	var coinsRoute []string
	for _, coin := range route {
		coinsRoute = append(coinsRoute, strconv.Itoa(int(coin)))
	}
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParamsWithTimeout(c.timeout).WithCoinCommission(&coinCommission).WithCoinToSell(&coinToSell).WithSwapFrom(&swapFrom).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithRoute(coinsRoute).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinSymbolSellAllExtended return estimate of sell all coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinSymbolSellAllExtended(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, swapFrom string, route []uint64, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	var coinsRoute []string
	for _, coin := range route {
		coinsRoute = append(coinsRoute, strconv.Itoa(int(coin)))
	}
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParamsWithTimeout(c.timeout).WithCoinToSell(&coinToSell).WithSwapFrom(&swapFrom).WithCoinToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithRoute(coinsRoute).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDBuyExtended return estimate of buy coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinIDBuyExtended(coinToSell, coinToBuy uint64, valueToBuy string, coinCommission uint64, swapFrom string, route []uint64, optionalHeight ...uint64) (*models.EstimateCoinBuyResponse, error) {
	var coinsRoute []string
	for _, coin := range route {
		coinsRoute = append(coinsRoute, strconv.Itoa(int(coin)))
	}
	res, err := c.ClientService.EstimateCoinBuy(api_service.NewEstimateCoinBuyParamsWithTimeout(c.timeout).WithCoinIDCommission(&coinCommission).WithCoinIDToSell(&coinToSell).WithSwapFrom(&swapFrom).WithCoinIDToBuy(&coinToBuy).WithValueToBuy(valueToBuy).WithHeight(optionalInt(optionalHeight)).WithRoute(coinsRoute).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSellExtended return estimate of sell coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinIDSellExtended(coinToBuy, coinToSell uint64, valueToSell string, coinCommission uint64, swapFrom string, route []uint64, optionalHeight ...uint64) (*models.EstimateCoinSellResponse, error) {
	var coinsRoute []string
	for _, coin := range route {
		coinsRoute = append(coinsRoute, strconv.Itoa(int(coin)))
	}
	res, err := c.ClientService.EstimateCoinSell(api_service.NewEstimateCoinSellParamsWithTimeout(c.timeout).WithCoinIDCommission(&coinCommission).WithCoinIDToSell(&coinToSell).WithSwapFrom(&swapFrom).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithHeight(optionalInt(optionalHeight)).WithRoute(coinsRoute).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateCoinIDSellAllExtended return estimate of sell all coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinIDSellAllExtended(coinToBuy, coinToSell, gasPrice uint64, valueToSell string, swapFrom string, route []uint64, optionalHeight ...uint64) (*models.EstimateCoinSellAllResponse, error) {
	var coinsRoute []string
	for _, coin := range route {
		coinsRoute = append(coinsRoute, strconv.Itoa(int(coin)))
	}
	res, err := c.ClientService.EstimateCoinSellAll(api_service.NewEstimateCoinSellAllParamsWithTimeout(c.timeout).WithCoinIDToSell(&coinToSell).WithSwapFrom(&swapFrom).WithCoinIDToBuy(&coinToBuy).WithValueToSell(valueToSell).WithGasPrice(&gasPrice).WithHeight(optionalInt(optionalHeight)).WithRoute(coinsRoute).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// EstimateTxCommission return estimate of encoding transaction.
func (c *Client) EstimateTxCommission(tx string, optionalHeight ...uint64) (*models.EstimateTxCommissionResponse, error) {
	res, err := c.ClientService.EstimateTxCommission(api_service.NewEstimateTxCommissionParamsWithTimeout(c.timeout).WithTx(tx).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Events returns events at given height.
func (c *Client) Events(height uint64, search ...string) (*models.EventsResponse, error) {
	res, err := c.ClientService.Events(api_service.NewEventsParamsWithTimeout(c.timeout).WithHeight(strconv.Itoa(int(height))).WithSearch(search).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MaxGasPrice returns current max gas.
func (c *Client) MaxGasPrice(optionalHeight ...uint64) (*models.MaxGasPriceResponse, error) {
	res, err := c.ClientService.MaxGasPrice(api_service.NewMaxGasPriceParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MinGasPrice returns current min gas price.
func (c *Client) MinGasPrice() (*models.MinGasPriceResponse, error) {
	res, err := c.ClientService.MinGasPrice(api_service.NewMinGasPriceParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// MissedBlocks returns missed blocks by validator public key.
func (c *Client) MissedBlocks(publicKey string, optionalHeight ...uint64) (*models.MissedBlocksResponse, error) {
	res, err := c.ClientService.MissedBlocks(api_service.NewMissedBlocksParamsWithTimeout(c.timeout).WithPublicKey(publicKey).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// NetInfo returns network res
func (c *Client) NetInfo() (*models.NetInfoResponse, error) {
	res, err := c.ClientService.NetInfo(api_service.NewNetInfoParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SendTransaction returns the result of sending signed tx.
// To ensure that transaction was successfully committed to the blockchain,
// you need to find the transaction by the hash and ensure that the status code equals to 0.
func (c *Client) SendTransaction(tx string) (*models.SendTransactionResponse, error) {
	res, err := c.ClientService.SendTransaction(api_service.NewSendTransactionParamsWithTimeout(c.timeout).WithTx(tx).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Transaction returns transaction res.
func (c *Client) Transaction(hash string) (*models.TransactionResponse, error) {
	res, err := c.ClientService.Transaction(api_service.NewTransactionParamsWithTimeout(c.timeout).WithHash(hash).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Transactions returns transactions by query.
func (c *Client) Transactions(query string, page, perPage int32) (*models.TransactionsResponse, error) {
	res, err := c.ClientService.Transactions(api_service.NewTransactionsParamsWithTimeout(c.timeout).WithQuery(query).WithPage(&page).WithPerPage(&perPage).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// UnconfirmedTxs returns unconfirmed transactions.
func (c *Client) UnconfirmedTxs(limit int32) (*models.UnconfirmedTxsResponse, error) {
	res, err := c.ClientService.UnconfirmedTxs(api_service.NewUnconfirmedTxsParamsWithTimeout(c.timeout).WithLimit(&limit).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Validators returns list of active validators.
func (c *Client) Validators(optionalHeight ...uint64) (*models.ValidatorsResponse, error) {
	res, err := c.ClientService.Validators(api_service.NewValidatorsParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// WaitList returns the list of address stakes in waitlist.
func (c *Client) WaitList(publicKey, address string, optionalHeight ...uint64) (*models.WaitListResponse, error) {
	res, err := c.ClientService.WaitList(api_service.NewWaitListParamsWithTimeout(c.timeout).WithAddress(address).WithPublicKey(&publicKey).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SwapPool returns total supply and reserves.
func (c *Client) SwapPool(coin0, coin1 uint64, optionalHeight ...uint64) (*models.SwapPoolResponse, error) {
	res, err := c.ClientService.SwapPool(api_service.NewSwapPoolParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithCoin0(strconv.Itoa(int(coin0))).WithCoin1(strconv.Itoa(int(coin1))).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SwapPoolProvider returns reserves and liquidity balance of provider.
func (c *Client) SwapPoolProvider(coin0, coin1 uint64, provider string, optionalHeight ...uint64) (*models.SwapPoolResponse, error) {
	res, err := c.ClientService.SwapPoolProvider(api_service.NewSwapPoolProviderParamsWithTimeout(c.timeout).WithProvider(provider).WithCoin0(strconv.Itoa(int(coin0))).WithCoin1(strconv.Itoa(int(coin1))).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// PriceCommission returns commissions.
func (c *Client) PriceCommission(optionalHeight ...uint64) (*models.PriceCommissionResponse, error) {
	res, err := c.ClientService.PriceCommission(api_service.NewPriceCommissionParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// VersionNetwork returns versions network.
func (c *Client) VersionNetwork() (*models.VersionNetworkResponse, error) {
	res, err := c.ClientService.VersionNetwork(api_service.NewVersionNetworkParamsWithTimeout(c.timeout).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// CommissionVotes returns votes for update commissions.
func (c *Client) CommissionVotes(target uint64, optionalHeight ...uint64) (*models.CommissionVotesResponse, error) {
	res, err := c.ClientService.CommissionVotes(api_service.NewCommissionVotesParamsWithTimeout(c.timeout).WithTargetVersion(strconv.Itoa(int(target))).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// UpdateVotes returns votes for update network.
func (c *Client) UpdateVotes(target uint64, optionalHeight ...uint64) (*models.UpdateVotesResponse, error) {
	res, err := c.ClientService.UpdateVotes(api_service.NewUpdateVotesParamsWithTimeout(c.timeout).WithTargetVersion(strconv.Itoa(int(target))).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// LimitOrder returns order by ID.
func (c *Client) LimitOrder(orderID uint64, optionalHeight ...uint64) (*models.LimitOrderResponse, error) {
	res, err := c.ClientService.LimitOrder(api_service.NewLimitOrderParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithOrderID(strconv.Itoa(int(orderID))).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// LimitOrders returns orders by IDs.
func (c *Client) LimitOrders(orderIDs []uint64, optionalHeight ...uint64) (*models.LimitOrdersResponse, error) {
	var ids []string
	for _, id := range orderIDs {
		ids = append(ids, strconv.Itoa(int(id)))
	}
	res, err := c.ClientService.LimitOrders(api_service.NewLimitOrdersParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithIds(ids).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// LimitOrdersOfPool returns sell orders for a pair of coins.
func (c *Client) LimitOrdersOfPool(sellCoin, buyCoin uint64, optionalHeight ...uint64) (*models.LimitOrdersOfPoolResponse, error) {
	res, err := c.ClientService.LimitOrdersOfPool(api_service.NewLimitOrdersOfPoolParamsWithTimeout(c.timeout).WithHeight(optionalInt(optionalHeight)).WithSellCoin(strconv.Itoa(int(sellCoin))).WithBuyCoin(strconv.Itoa(int(buyCoin))).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// Frozen returns frozen balance.
// Deprecated: Use FrozenAll instead.
func (c *Client) Frozen(address string, coinID *uint64, optionalHeight ...uint64) (*models.FrozenResponse, error) {
	res, err := c.ClientService.Frozen(api_service.NewFrozenParamsWithTimeout(c.timeout).WithAddress(address).WithCoinID(coinID).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// FrozenAll returns frozen balance.
func (c *Client) FrozenAll(addresses []string, coinIDs []uint64, startHeight, endHeight uint64, optionalHeight ...uint64) (*models.FrozenResponse, error) {
	var coinIds []string
	for _, id := range coinIDs {
		coinIds = append(coinIds, strconv.Itoa(int(id)))
	}
	res, err := c.ClientService.FrozenAll(api_service.NewFrozenAllParamsWithTimeout(c.timeout).WithAddresses(addresses).WithCoinIds(coinIds).WithStartHeight(&startHeight).WithEndHeight(&endHeight).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

// SwapPools returns list of all pools.
func (c *Client) SwapPools(orders bool, optionalHeight ...uint64) (*models.SwapPoolsResponse, error) {
	res, err := c.ClientService.SwapPools(api_service.NewSwapPoolsParamsWithTimeout(c.timeout).WithOrders(&orders).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
	if err != nil {
		return nil, err
	}

	return res.GetPayload(), nil
}

type BestTradeRequest_Type string

const (
	BestTradeRequest_input  BestTradeRequest_Type = "input"
	BestTradeRequest_output BestTradeRequest_Type = "output"
)

// BestTrade returns optimal exchange route.
func (c *Client) BestTrade(sellCoinID, buyCoinID uint64, t BestTradeRequest_Type, amount string, maxDepth int32, optionalHeight ...uint64) (*models.BestTradeResponse, error) {
	res, err := c.ClientService.BestTrade(api_service.NewBestTradeParamsWithTimeout(c.timeout).WithSellCoin(strconv.Itoa(int(sellCoinID))).WithBuyCoin(strconv.Itoa(int(buyCoinID))).WithAmount(amount).WithMaxDepth(&maxDepth).WithType(string(t)).WithHeight(optionalInt(optionalHeight)).WithContext(c.ctxFunc()), c.opts...)
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
	return s.conn.CloseHandler()(websocket.CloseNormalClosure, "")
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
		//if recv.Error != nil {
		//	return nil, errors.New(fmt.Sprintf("%v",recv.Error))
		//} todo
		//return recv.Result, err

		return &recv, err
	}
}

// Subscribe returns a subscription for events by query.
func (c *Client) Subscribe(query string) (*SubscriberClient, error) {

	subClient := &SubscriberClient{}

	subClient.ctx, subClient.cancel = context.WithCancel(c.ctxFunc())
	scheme := c.getWSProtocol()
	u := url.URL{Scheme: scheme, Host: c.host, Path: c.path + "/subscribe", RawQuery: "query=" + url.QueryEscape(query)}

	var err error
	subClient.conn, _, err = websocket.DefaultDialer.DialContext(subClient.ctx, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return subClient, nil
}

func (c *Client) getWSProtocol() string {
	if c.ssl {
		return "wss"
	}
	return "ws"
}

func (c *Client) getHTTPProtocol() string {
	if c.ssl {
		return "https"
	}
	return "http"
}

// CheckVersion compares the prefix in the version name and checks the testnet mode
func (c *Client) CheckVersion(version string, isTestnet bool) error {
	response, err := c.Status()
	if err != nil {
		return err
	}
	if len(response.Version) < len(version) {
		return fmt.Errorf("node version is %s", response.Version)
	}
	if !strings.HasPrefix(response.Version, version) {
		return fmt.Errorf("node version is %s", response.Version[:len(version)])
	}
	if isTestnet != strings.HasSuffix(response.Version, "testnet") {
		return errors.New("node version is not testnet")
	}
	return nil
}

// Marshal returns model in JSON format
func (c *Client) Marshal(i interface{}) (json string, err error) {
	return Marshal(i)
}

// Marshal returns model in JSON format
func Marshal(i interface{}) (json string, err error) {
	if i == nil {
		return "", nil
	}

	m, ok := i.(interface{ MarshalBinary() ([]byte, error) })
	if !ok {
		marshal, err := swag.WriteJSON(m)
		if err != nil {
			return "", err
		}

		return string(marshal), nil
	}
	marshal, err := m.MarshalBinary()
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

type defaultError interface {
	Code() int
	GetPayload() *models.ErrorBody
}

// ErrorBody returns error as API model
func (c *Client) ErrorBody(err error) (int, *models.ErrorBody, error) {
	return ErrorBody(err)
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
