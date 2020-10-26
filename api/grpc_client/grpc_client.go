package grpc_client

import (
	"context"
	"errors"
	"fmt"
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
	"net/http"
	"strconv"
)

// Client gRPC
type Client struct {
	grpcClient api_pb.ApiServiceClient
	ctxFunc    func() context.Context
	marshaler  runtime.Marshaler
	opts       []grpc.CallOption
}

// New returns gRPC Client
func New(address string, _ ...string) (*Client, error) {
	clientConn, err := grpc.Dial(address,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(200000000)))
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

// GRPCClient return gRPC client api_pb.ApiServiceClient
func (c *Client) GRPCClient() api_pb.ApiServiceClient {
	return c.grpcClient
}

// CoinID returns ID of coin symbol.
func (c *Client) CoinID(symbol string) (uint64, error) {
	info, err := c.CoinInfo(symbol)
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
	return c.AddressWithMoreInfo(address, true, false, optionalHeight...)
}

// Addresses returns list of addresses.
func (c *Client) Addresses(addresses []string, optionalHeight ...uint64) (*api_pb.AddressesResponse, error) {
	return c.AddressesWithMoreInfo(addresses, true, false, optionalHeight...)
}

// AddressWithMoreInfo returns coins list with bipValue, balance, delegated and transaction count of an address.
func (c *Client) AddressWithMoreInfo(address string, bipValue, delegated bool, optionalHeight ...uint64) (*api_pb.AddressResponse, error) {
	return c.grpcClient.Address(c.ctxFunc(), &api_pb.AddressRequest{Height: optionalInt(optionalHeight), Address: address, BipValue: bipValue, Delegated: delegated}, c.opts...)
}

// AddressesWithMoreInfo returns list of addresses with bipValue and delegated.
func (c *Client) AddressesWithMoreInfo(addresses []string, bipValue, delegated bool, optionalHeight ...uint64) (*api_pb.AddressesResponse, error) {
	return c.grpcClient.Addresses(c.ctxFunc(), &api_pb.AddressesRequest{Addresses: addresses, Height: optionalInt(optionalHeight), BipValue: bipValue, Delegated: delegated}, c.opts...)
}

// Block returns block data at given height.
func (c *Client) Block(height uint64) (*api_pb.BlockResponse, error) {
	return c.grpcClient.Block(c.ctxFunc(), &api_pb.BlockRequest{Height: height}, c.opts...)
}

// Candidate returns candidate’s info by provided public_key.
func (c *Client) Candidate(publicKey string, optionalHeight ...uint64) (*api_pb.CandidateResponse, error) {
	return c.grpcClient.Candidate(c.ctxFunc(), &api_pb.CandidateRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey}, c.opts...)
}

// Candidates returns list of candidates.
func (c *Client) Candidates(includeStakes bool, status api_pb.CandidatesRequest_CandidateStatus, optionalHeight ...uint64) (*api_pb.CandidatesResponse, error) {
	return c.grpcClient.Candidates(c.ctxFunc(), &api_pb.CandidatesRequest{Height: optionalInt(optionalHeight), IncludeStakes: includeStakes, Status: status}, c.opts...)
}

// CoinInfoByID returns information about coin ID.
func (c *Client) CoinInfoByID(id uint64, optionalHeight ...uint64) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfoById(c.ctxFunc(), &api_pb.CoinIdRequest{Height: optionalInt(optionalHeight), Id: id}, c.opts...)
}

// CoinInfo returns information about coin symbol.
func (c *Client) CoinInfo(symbol string, optionalHeight ...uint64) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfo(c.ctxFunc(), &api_pb.CoinInfoRequest{Height: optionalInt(optionalHeight), Symbol: symbol}, c.opts...)
}

// EstimateCoinSymbolBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinSymbolBuy(coinToSell, coinToBuy string, valueToBuy string, optionalHeight ...uint64) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToBuy: valueToBuy}, c.opts...)
}

// EstimateCoinSymbolSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinSymbolSell(coinToBuy, coinToSell string, valueToSell string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToSell}, c.opts...)
}

// EstimateCoinSymbolSellAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinSymbolSellAll(coinToBuy, coinToSell string, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToSell, GasPrice: gasPrice}, c.opts...)
}

// EstimateCoinIDBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinIDBuy(coinToSell, coinToBuy uint64, valueToBuy string, optionalHeight ...uint64) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToBuy: valueToBuy}, c.opts...)
}

// EstimateCoinIDSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinIDSell(coinToBuy, coinToSell uint64, valueToSell string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToSell}, c.opts...)
}

// EstimateCoinIDSellAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinIDSellAll(coinToBuy, coinToSell uint64, gasPrice uint64, valueToSell string, optionalHeight ...uint64) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToSell, GasPrice: gasPrice}, c.opts...)
}

// EstimateTxCommission return estimate of encoding transaction.
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
func (c *Client) WaitList(publicKey, address string, height ...uint64) (*api_pb.WaitListResponse, error) {
	return c.grpcClient.WaitList(c.ctxFunc(), &api_pb.WaitListRequest{Height: optionalInt(height), PublicKey: publicKey, Address: address}, c.opts...)
}

// Subscribe returns a subscription for events by query.
func (c *Client) Subscribe(query string) (api_pb.ApiService_SubscribeClient, error) {
	return c.grpcClient.Subscribe(c.ctxFunc(), &api_pb.SubscribeRequest{Query: query}, c.opts...)
}

// Frozen returns frozen balance.
func (c *Client) Frozen(address string, optionalCoinID ...uint64) (*api_pb.FrozenResponse, error) {
	if len(optionalCoinID) > 1 {
		return nil, errors.New("CoinID needed single value") // todo: change message
	}

	var coin *wrapperspb.UInt64Value
	if len(optionalCoinID) == 1 {
		coin = wrapperspb.UInt64(optionalInt(optionalCoinID))
	}
	return c.grpcClient.Frozen(c.ctxFunc(), &api_pb.FrozenRequest{Address: address, CoinId: coin}, c.opts...)
}

func optionalInt(height []uint64) uint64 {
	if len(height) == 1 {
		return height[0]
	}
	return 0
}
