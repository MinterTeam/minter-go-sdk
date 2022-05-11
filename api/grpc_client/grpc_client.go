package grpc_client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	_struct "google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Client gRPC
type Client struct {
	grpcClient api_pb.ApiServiceClient
	ctxFunc    func() context.Context
	marshaler  runtime.Marshaler
	opts       []grpc.CallOption
}

// New returns gRPC Client
func New(address string) (*Client, error) {
	clientConn, err := grpc.Dial(address,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1000000000)))
	if err != nil {
		return nil, err
	}

	return &Client{grpcClient: api_pb.NewApiServiceClient(clientConn), ctxFunc: context.Background,
		marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}, nil
}

// WithContextFunc returns new Client with new context
// Example:
// 		timeout := func(c context.Context) func() context.Context {
//			return func() context.Context {
//				ctx, _ := context.WithTimeout(c, 10*time.Second)
//				return ctx
//			}
//		}
func (c *Client) WithContextFunc(contextFunc func(context.Context) func() context.Context) *Client {
	return &Client{grpcClient: c.grpcClient, ctxFunc: contextFunc(c.ctxFunc()), opts: c.opts}
}

// WithCallOption returns new Client with additional grpc.CallOption
func (c *Client) WithCallOption(opts ...grpc.CallOption) *Client {
	return &Client{grpcClient: c.grpcClient, ctxFunc: c.ctxFunc, opts: append(c.opts, opts...)}
}

// GRPCClient returns gRPC client api_pb.ApiServiceClient
func (c *Client) GRPCClient() api_pb.ApiServiceClient {
	return c.grpcClient
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

// CoinID returns ID of coin symbol.
func (c *Client) CoinID(symbol string, optionalHeight ...uint64) (uint64, error) {
	info, err := c.CoinInfo(symbol, optionalHeight...)
	if err != nil {
		return 0, err
	}
	return info.Id, err
}

// Nonce returns next transaction number (nonce) of an address.
func (c *Client) Nonce(address string) (uint64, error) {
	res, err := c.Address(address)
	if err != nil {
		return 0, err
	}

	return res.TransactionCount + 1, nil
}

// ErrorBody returns error as API model
func (c *Client) ErrorBody(err error) (int, *api_pb.ErrorBody, error) {
	if err == nil {
		return http.StatusOK, nil, nil
	}

	s, ok := status.FromError(err)
	if !ok {
		return 0, nil, err
	}

	if s.Code() == codes.OK {
		return http.StatusOK, nil, nil
	}

	statusCode := runtime.HTTPStatusFromCode(s.Code())
	errorBody := &api_pb.ErrorBody{
		Error: &api_pb.ErrorBody_Error{
			Code:    strconv.Itoa(statusCode),
			Message: s.Message(),
			Data:    map[string]string{},
		},
	}

	details := s.Details()
	if len(details) == 0 {
		return statusCode, errorBody, nil
	}

	detail, ok := details[0].(*_struct.Struct)
	if !ok {
		return statusCode, errorBody, nil
	}

	data := detail.AsMap()
	code, ok := data["code"]
	if ok {
		errorBody.Error.Code = fmt.Sprintf("%v", code)
		delete(data, "code")
	}
	for k, v := range data {
		errorBody.Error.Data[k] = fmt.Sprintf("%s", v)
	}

	return statusCode, errorBody, nil
}

// HTTPError returns error as JSON API
func (c *Client) HTTPError(statusError error) (statusCode int, json string, err error) {
	statusCode, errorBody, err := c.ErrorBody(statusError)
	if err != nil {
		return 0, "", err
	}

	if errorBody == nil {
		return statusCode, "", nil
	}

	jErr, err := c.Marshal(errorBody)
	if err != nil {
		return statusCode, "", err
	}

	return statusCode, jErr, nil
}

// Marshal returns model in JSON format
func (c *Client) Marshal(m proto.Message) (json string, err error) {
	marshal, err := c.marshaler.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

// Halts returns the candidate votes for stopping the network at block.
func (c *Client) Halts(height uint64) (*api_pb.HaltsResponse, error) {
	return c.grpcClient.Halts(c.ctxFunc(), &api_pb.HaltsRequest{Height: height}, c.opts...)
}

// Genesis returns genesis file.
func (c *Client) Genesis() (*api_pb.GenesisResponse, error) {
	return c.grpcClient.Genesis(c.ctxFunc(), &empty.Empty{}, c.opts...)
}

// Status returns node status including pubkey, latest block.
func (c *Client) Status() (*api_pb.StatusResponse, error) {
	return c.grpcClient.Status(c.ctxFunc(), &empty.Empty{}, c.opts...)
}

// Address returns coins list, balance and transaction count of an address.
func (c *Client) Address(address string, optionalHeight ...uint64) (*api_pb.AddressResponse, error) {
	return c.AddressExtended(address, false, optionalHeight...)
}

// Addresses returns list of addresses.
func (c *Client) Addresses(addresses []string, optionalHeight ...uint64) (*api_pb.AddressesResponse, error) {
	return c.AddressesExtended(addresses, false, optionalHeight...)
}

// AddressExtended returns coins list with balance, delegated and transaction count of an address.
func (c *Client) AddressExtended(address string, delegated bool, optionalHeight ...uint64) (*api_pb.AddressResponse, error) {
	return c.grpcClient.Address(c.ctxFunc(), &api_pb.AddressRequest{Height: optionalInt(optionalHeight), Address: address, Delegated: delegated}, c.opts...)
}

// AddressesExtended returns list of addresses with bipValue and delegated.
func (c *Client) AddressesExtended(addresses []string, delegated bool, optionalHeight ...uint64) (*api_pb.AddressesResponse, error) {
	return c.grpcClient.Addresses(c.ctxFunc(), &api_pb.AddressesRequest{Addresses: addresses, Height: optionalInt(optionalHeight), Delegated: delegated}, c.opts...)
}

// Block returns block data at given height.
func (c *Client) Block(height uint64) (*api_pb.BlockResponse, error) {
	return c.grpcClient.Block(c.ctxFunc(), &api_pb.BlockRequest{Height: height}, c.opts...)
}

// Blocks returns blocks at given interval.
func (c *Client) Blocks(from, to uint64, failedTxs, events bool, fieldsBlock ...string) (*api_pb.BlocksResponse, error) {
	var fields []api_pb.BlockField
	for _, field := range fieldsBlock {
		fields = append(fields, api_pb.BlockField(api_pb.BlockField_value[field]))
	}
	return c.grpcClient.Blocks(c.ctxFunc(), &api_pb.BlocksRequest{FromHeight: from, ToHeight: to, Fields: fields, FailedTxs: failedTxs, Events: events}, c.opts...)
}

// BlockExtended returns block data at given height by filtered data.
func (c *Client) BlockExtended(height uint64, failedTxs, events bool, fieldsBlock ...string) (*api_pb.BlockResponse, error) {
	var fields []api_pb.BlockField
	for _, field := range fieldsBlock {
		fields = append(fields, api_pb.BlockField(api_pb.BlockField_value[field]))
	}
	return c.grpcClient.Block(c.ctxFunc(), &api_pb.BlockRequest{Height: height, FailedTxs: failedTxs, Fields: fields, Events: events}, c.opts...)
}

// Candidate returns candidate’s info by provided public_key.
func (c *Client) Candidate(publicKey string, optionalHeight ...uint64) (*api_pb.CandidateResponse, error) {
	return c.CandidateExtended(publicKey, false, optionalHeight...)
}

// Candidates returns list of candidates.
func (c *Client) Candidates(includeStakes bool, optionalHeight ...uint64) (*api_pb.CandidatesResponse, error) {
	return c.CandidatesExtended(includeStakes, true, "", optionalHeight...)
}

// CandidateExtended returns candidate’s info by provided public_key.
func (c *Client) CandidateExtended(publicKey string, notShowStakes bool, optionalHeight ...uint64) (*api_pb.CandidateResponse, error) {
	return c.grpcClient.Candidate(c.ctxFunc(), &api_pb.CandidateRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey, NotShowStakes: notShowStakes}, c.opts...)
}

// CandidatesExtended returns list of candidates.
func (c *Client) CandidatesExtended(includeStakes, notShowStakes bool, status string, optionalHeight ...uint64) (*api_pb.CandidatesResponse, error) {
	return c.grpcClient.Candidates(c.ctxFunc(), &api_pb.CandidatesRequest{Height: optionalInt(optionalHeight), IncludeStakes: includeStakes, NotShowStakes: notShowStakes, Status: api_pb.CandidatesRequest_CandidateStatus(api_pb.CandidatesRequest_CandidateStatus_value[status])}, c.opts...)
}

// CoinInfoByID returns information about coin ID.
func (c *Client) CoinInfoByID(id uint64, optionalHeight ...uint64) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfoById(c.ctxFunc(), &api_pb.CoinIdRequest{Height: optionalInt(optionalHeight), Id: id}, c.opts...)
}

// CoinInfo returns information about coin symbol.
func (c *Client) CoinInfo(symbol string, optionalHeight ...uint64) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfo(c.ctxFunc(), &api_pb.CoinInfoRequest{Height: optionalInt(optionalHeight), Symbol: symbol}, c.opts...)
}

// EstimateCoinSymbolBuy returns estimate of buy coin transaction.
func (c *Client) EstimateCoinSymbolBuy(coinToSell, coinToBuy string, valueToBuy string, coinCommission string, optionalHeight ...uint64) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Commission: &api_pb.EstimateCoinBuyRequest_CoinCommission{CoinCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToBuy: valueToBuy}, c.opts...)
}

// EstimateCoinSymbolSell returns estimate of sell coin transaction.
func (c *Client) EstimateCoinSymbolSell(coinToBuy, coinToSell string, valueToSell string, coinCommission string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Commission: &api_pb.EstimateCoinSellRequest_CoinCommission{CoinCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToSell}, c.opts...)
}

// EstimateCoinSymbolSellAll returns estimate of sell all coin transaction.
func (c *Client) EstimateCoinSymbolSellAll(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToSell, GasPrice: gasPrice}, c.opts...)
}

// EstimateCoinIDBuy returns estimate of buy coin transaction.
func (c *Client) EstimateCoinIDBuy(coinToSell, coinToBuy uint64, valueToBuy string, coinCommission uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Commission: &api_pb.EstimateCoinBuyRequest_CoinIdCommission{CoinIdCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToBuy: valueToBuy}, c.opts...)
}

// EstimateCoinIDSell returns estimate of sell coin transaction.
func (c *Client) EstimateCoinIDSell(coinToBuy, coinToSell uint64, valueToSell string, coinCommission uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Commission: &api_pb.EstimateCoinSellRequest_CoinIdCommission{CoinIdCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToSell}, c.opts...)
}

// EstimateCoinIDSellAll returns estimate of sell all coin transaction.
func (c *Client) EstimateCoinIDSellAll(coinToBuy, coinToSell uint64, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToSell, GasPrice: gasPrice}, c.opts...)
}

// EstimateCoinSymbolBuyExtended returns estimate of buy coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinSymbolBuyExtended(coinToSell, coinToBuy string, valueToBuy string, coinCommission string, swapFrom string, route []uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Commission: &api_pb.EstimateCoinBuyRequest_CoinCommission{CoinCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToBuy: valueToBuy, SwapFrom: api_pb.SwapFrom(api_pb.SwapFrom_value[swapFrom]), Route: route}, c.opts...)
}

// EstimateCoinSymbolSellExtended returns estimate of sell coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinSymbolSellExtended(coinToBuy, coinToSell string, valueToSell string, coinCommission string, swapFrom string, route []uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Commission: &api_pb.EstimateCoinSellRequest_CoinCommission{CoinCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToSell, SwapFrom: api_pb.SwapFrom(api_pb.SwapFrom_value[swapFrom]), Route: route}, c.opts...)
}

// EstimateCoinSymbolSellAllExtended returns estimate of sell all coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinSymbolSellAllExtended(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, swapFrom string, route []uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToSell, GasPrice: gasPrice, SwapFrom: api_pb.SwapFrom(api_pb.SwapFrom_value[swapFrom]), Route: route}, c.opts...)
}

// EstimateCoinIDBuyExtended returns estimate of buy coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinIDBuyExtended(coinToSell, coinToBuy uint64, valueToBuy string, coinCommission uint64, swapFrom string, route []uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Commission: &api_pb.EstimateCoinBuyRequest_CoinIdCommission{CoinIdCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToBuy: valueToBuy, SwapFrom: api_pb.SwapFrom(api_pb.SwapFrom_value[swapFrom]), Route: route}, c.opts...)
}

// EstimateCoinIDSellExtended returns estimate of sell coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinIDSellExtended(coinToBuy, coinToSell uint64, valueToSell string, coinCommission uint64, swapFrom string, route []uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Commission: &api_pb.EstimateCoinSellRequest_CoinIdCommission{CoinIdCommission: coinCommission}, Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToSell, SwapFrom: api_pb.SwapFrom(api_pb.SwapFrom_value[swapFrom]), Route: route}, c.opts...)
}

// EstimateCoinIDSellAllExtended returns estimate of sell all coin transaction with choice of the exchange source.
func (c *Client) EstimateCoinIDSellAllExtended(coinToBuy, coinToSell uint64, gasPrice uint64, valueToSell string, swapFrom string, route []uint64, optionalHeight ...uint64) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToSell, GasPrice: gasPrice, SwapFrom: api_pb.SwapFrom(api_pb.SwapFrom_value[swapFrom]), Route: route}, c.opts...)
}

// EstimateTxCommission returns estimate of encoding transaction with choice of the exchange source.
func (c *Client) EstimateTxCommission(tx string, optionalHeight ...uint64) (*api_pb.EstimateTxCommissionResponse, error) {
	return c.grpcClient.EstimateTxCommission(c.ctxFunc(), &api_pb.EstimateTxCommissionRequest{Height: optionalInt(optionalHeight), Tx: tx}, c.opts...)
}

// Events returns events at given height.
func (c *Client) Events(height uint64, search ...string) (*api_pb.EventsResponse, error) {
	return c.grpcClient.Events(c.ctxFunc(), &api_pb.EventsRequest{Height: height, Search: search}, c.opts...)
}

// MaxGasPrice returns current max gas.
func (c *Client) MaxGasPrice(optionalHeight ...uint64) (*api_pb.MaxGasPriceResponse, error) {
	return c.grpcClient.MaxGasPrice(c.ctxFunc(), &api_pb.MaxGasPriceRequest{Height: optionalInt(optionalHeight)}, c.opts...)
}

// MinGasPrice returns current min gas price.
func (c *Client) MinGasPrice() (*api_pb.MinGasPriceResponse, error) {
	return c.grpcClient.MinGasPrice(c.ctxFunc(), &empty.Empty{}, c.opts...)
}

// MissedBlocks returns missed blocks by validator public key.
func (c *Client) MissedBlocks(publicKey string, optionalHeight ...uint64) (*api_pb.MissedBlocksResponse, error) {
	return c.grpcClient.MissedBlocks(c.ctxFunc(), &api_pb.MissedBlocksRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey}, c.opts...)
}

// NetInfo returns network info
func (c *Client) NetInfo() (*api_pb.NetInfoResponse, error) {
	return c.grpcClient.NetInfo(c.ctxFunc(), &empty.Empty{}, c.opts...)
}

// SendTransaction returns the result of sending signed tx.
// To ensure that transaction was successfully committed to the blockchain,
// you need to find the transaction by the hash and ensure that the status code equals to 0.
func (c *Client) SendTransaction(tx string) (*api_pb.SendTransactionResponse, error) {
	return c.grpcClient.SendTransaction(c.ctxFunc(), &api_pb.SendTransactionRequest{Tx: tx}, c.opts...)
}

// Transaction returns transaction info.
func (c *Client) Transaction(hash string) (*api_pb.TransactionResponse, error) {
	return c.grpcClient.Transaction(c.ctxFunc(), &api_pb.TransactionRequest{Hash: hash}, c.opts...)
}

// Transactions returns transactions by query.
func (c *Client) Transactions(query string, page, perPage int) (*api_pb.TransactionsResponse, error) {
	return c.grpcClient.Transactions(c.ctxFunc(), &api_pb.TransactionsRequest{Query: query, Page: int32(page), PerPage: int32(perPage)}, c.opts...)
}

// UnconfirmedTxs returns unconfirmed transactions.
func (c *Client) UnconfirmedTxs(limit ...uint64) (*api_pb.UnconfirmedTxsResponse, error) {
	return c.grpcClient.UnconfirmedTxs(c.ctxFunc(), &api_pb.UnconfirmedTxsRequest{Limit: int32(optionalInt(limit))}, c.opts...)
}

// Validators returns list of active validators.
func (c *Client) Validators(optionalHeight ...uint64) (*api_pb.ValidatorsResponse, error) {
	return c.grpcClient.Validators(c.ctxFunc(), &api_pb.ValidatorsRequest{Height: optionalInt(optionalHeight)}, c.opts...)
}

// WaitList returns the list of address stakes in waitlist.
func (c *Client) WaitList(publicKey, address string, optionalHeight ...uint64) (*api_pb.WaitListResponse, error) {
	return c.grpcClient.WaitList(c.ctxFunc(), &api_pb.WaitListRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey, Address: address}, c.opts...)
}

// WaitLists returns the list addresses and stakes in waitlist.
func (c *Client) WaitLists(optionalHeight ...uint64) (*api_pb.WaitListsResponse, error) {
	return c.grpcClient.WaitLists(c.ctxFunc(), &api_pb.WaitListsRequest{Height: optionalInt(optionalHeight)}, c.opts...)
}

// SwapPool returns total supply and reserves.
func (c *Client) SwapPool(coin0, coin1 uint64, optionalHeight ...uint64) (*api_pb.SwapPoolResponse, error) {
	return c.grpcClient.SwapPool(c.ctxFunc(), &api_pb.SwapPoolRequest{Height: optionalInt(optionalHeight), Coin0: coin0, Coin1: coin1}, c.opts...)
}

// SwapPoolProvider returns reserves and liquidity balance of provider.
func (c *Client) SwapPoolProvider(coin0, coin1 uint64, provider string, optionalHeight ...uint64) (*api_pb.SwapPoolResponse, error) {
	return c.grpcClient.SwapPoolProvider(c.ctxFunc(), &api_pb.SwapPoolProviderRequest{Height: optionalInt(optionalHeight), Coin0: coin0, Coin1: coin1, Provider: provider}, c.opts...)
}

// VersionNetwork returns versions network.
func (c *Client) VersionNetwork() (*api_pb.VersionNetworkResponse, error) {
	return c.grpcClient.VersionNetwork(c.ctxFunc(), &api_pb.VersionNetworkRequest{}, c.opts...)
}

// PriceCommission returns commissions.
func (c *Client) PriceCommission(optionalHeight ...uint64) (*api_pb.PriceCommissionResponse, error) {
	return c.grpcClient.PriceCommission(c.ctxFunc(), &api_pb.PriceCommissionRequest{Height: optionalInt(optionalHeight)}, c.opts...)
}

// CommissionVotes returns votes for update commissions.
func (c *Client) CommissionVotes(target uint64, optionalHeight ...uint64) (*api_pb.CommissionVotesResponse, error) {
	return c.grpcClient.CommissionVotes(c.ctxFunc(), &api_pb.CommissionVotesRequest{Height: optionalInt(optionalHeight), TargetVersion: target}, c.opts...)
}

// UpdateVotes returns votes for update network.
func (c *Client) UpdateVotes(target uint64, optionalHeight ...uint64) (*api_pb.UpdateVotesResponse, error) {
	return c.grpcClient.UpdateVotes(c.ctxFunc(), &api_pb.UpdateVotesRequest{Height: optionalInt(optionalHeight), TargetVersion: target}, c.opts...)
}

// LimitOrder returns order by ID.
func (c *Client) LimitOrder(orderID uint64, optionalHeight ...uint64) (*api_pb.LimitOrderResponse, error) {
	return c.grpcClient.LimitOrder(c.ctxFunc(), &api_pb.LimitOrderRequest{Height: optionalInt(optionalHeight), OrderId: orderID}, c.opts...)
}

// LimitOrders returns orders by IDs.
func (c *Client) LimitOrders(orderIDs []uint64, optionalHeight ...uint64) (*api_pb.LimitOrdersResponse, error) {
	return c.grpcClient.LimitOrders(c.ctxFunc(), &api_pb.LimitOrdersRequest{Height: optionalInt(optionalHeight), Ids: orderIDs}, c.opts...)
}

// LimitOrdersOfPool returns sell orders for a pair of coins.
func (c *Client) LimitOrdersOfPool(sellCoin, buyCoin uint64, optionalHeight ...uint64) (*api_pb.LimitOrdersOfPoolResponse, error) {
	return c.grpcClient.LimitOrdersOfPool(c.ctxFunc(), &api_pb.LimitOrdersOfPoolRequest{Height: optionalInt(optionalHeight), SellCoin: sellCoin, BuyCoin: buyCoin}, c.opts...)
}

// Subscribe returns a subscription for events by query.
func (c *Client) Subscribe(query string) (api_pb.ApiService_SubscribeClient, error) {
	return c.grpcClient.Subscribe(c.ctxFunc(), &api_pb.SubscribeRequest{Query: query}, c.opts...)
}

// Frozen returns frozen balance.
// Deprecated: Use FrozenAll instead.
func (c *Client) Frozen(address string, coinID *uint64, optionalHeight ...uint64) (*api_pb.FrozenResponse, error) {
	var coin *wrapperspb.UInt64Value
	if coinID != nil {
		coin = wrapperspb.UInt64(*coinID)
	}
	return c.grpcClient.Frozen(c.ctxFunc(), &api_pb.FrozenRequest{Address: address, CoinId: coin, Height: optionalInt(optionalHeight)}, c.opts...)
}

// FrozenAll returns frozen balance.
func (c *Client) FrozenAll(addresses []string, coinIDs []uint64, startHeight, endHeight uint64, optionalHeight ...uint64) (*api_pb.FrozenResponse, error) {
	return c.grpcClient.FrozenAll(c.ctxFunc(), &api_pb.FrozenAllRequest{Addresses: addresses, CoinIds: coinIDs, StartHeight: startHeight, EndHeight: endHeight, Height: optionalInt(optionalHeight)}, c.opts...)
}

// SwapPools returns list of all pools.
func (c *Client) SwapPools(orders bool, optionalHeight ...uint64) (*api_pb.SwapPoolsResponse, error) {
	return c.grpcClient.SwapPools(c.ctxFunc(), &api_pb.SwapPoolsRequest{
		Height: optionalInt(optionalHeight),
		Orders: orders,
	}, c.opts...)
}

// BestTrade returns optimal exchange route.
func (c *Client) BestTrade(sellCoinID, buyCoinID uint64, t api_pb.BestTradeRequest_Type, amount string, maxDepth int32, optionalHeight ...uint64) (*api_pb.BestTradeResponse, error) {
	return c.grpcClient.BestTrade(c.ctxFunc(), &api_pb.BestTradeRequest{
		SellCoin: sellCoinID,
		BuyCoin:  buyCoinID,
		Amount:   amount,
		Type:     t,
		Height:   optionalInt(optionalHeight),
		MaxDepth: maxDepth,
	}, c.opts...)
}

func optionalInt(height []uint64) uint64 {
	if len(height) == 1 {
		return height[0]
	}
	return 0
}
