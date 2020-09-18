package grpc_client

import (
	"context"
	"fmt"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_struct "google.golang.org/protobuf/types/known/structpb"
	"net/http"
	"strconv"
)

type Client struct {
	grpcClient api_pb.ApiServiceClient
	ctxFunc    func() context.Context
	marshaler  runtime.Marshaler
}

// New gRPC Client
func New(address string) (*Client, error) {
	clientConn, err := grpc.Dial(address,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{grpcClient: api_pb.NewApiServiceClient(clientConn), ctxFunc: context.Background, marshaler: &runtime.JSONPb{OrigName: true, EmitDefaults: true}}, nil
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
	return &Client{grpcClient: c.grpcClient, ctxFunc: contextFunc(c.ctxFunc())}
}

// GRPCClient return gRPC client ApiServiceClient
func (c *Client) GRPCClient() api_pb.ApiServiceClient {
	return c.grpcClient
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

// HttpError returns error as JSON API
func (c *Client) HttpError(statusError error) (statusCode int, json string, err error) {
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

// Halts
func (c *Client) Halts(height int) (*api_pb.HaltsResponse, error) {
	return c.grpcClient.Halts(c.ctxFunc(), &api_pb.HaltsRequest{Height: uint64(height)})
}

// Genesis returns genesis file.
func (c *Client) Genesis() (*api_pb.GenesisResponse, error) {
	return c.grpcClient.Genesis(c.ctxFunc(), &empty.Empty{})
}

// Nonce returns next transaction number (nonce) of an address.
func (c *Client) Nonce(address string) (uint64, error) {
	res, err := c.Address(address)
	if err != nil {
		return 0, err
	}

	transactionsCount, err := strconv.Atoi(res.TransactionCount)
	if err != nil {
		return 0, err
	}

	return uint64(transactionsCount) + 1, nil
}

// Status returns node status including pubkey, latest block.
func (c *Client) Status() (*api_pb.StatusResponse, error) {
	return c.grpcClient.Status(c.ctxFunc(), &empty.Empty{})
}

// Address returns coins list, balance and transaction count of an address.
func (c *Client) Address(address string, optionalHeight ...int) (*api_pb.AddressResponse, error) {
	return c.grpcClient.Address(c.ctxFunc(), &api_pb.AddressRequest{Height: optionalInt(optionalHeight), Address: address})
}

// Addresses returns list of addresses.
func (c *Client) Addresses(addresses []string, optionalHeight ...int) (*api_pb.AddressesResponse, error) {
	return c.grpcClient.Addresses(c.ctxFunc(), &api_pb.AddressesRequest{Addresses: addresses, Height: optionalInt(optionalHeight)})
}

// Block returns block data at given height.
func (c *Client) Block(height int) (*api_pb.BlockResponse, error) {
	return c.grpcClient.Block(c.ctxFunc(), &api_pb.BlockRequest{Height: uint64(height)})
}

// Candidate returns candidate’s info by provided public_key. It will respond with 404 code if candidate is not found.
func (c *Client) Candidate(publicKey string, optionalHeight ...int) (*api_pb.CandidateResponse, error) {
	return c.grpcClient.Candidate(c.ctxFunc(), &api_pb.CandidateRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

// Candidates returns list of candidates.
func (c *Client) Candidates(includeStakes bool, optionalHeight ...int) (*api_pb.CandidatesResponse, error) {
	return c.grpcClient.Candidates(c.ctxFunc(), &api_pb.CandidatesRequest{Height: optionalInt(optionalHeight), IncludeStakes: includeStakes})
}

// CoinInfoByID returns information about coin ID. Note: this method does not return information about base coins (MNT and BIP).
func (c *Client) CoinInfoByID(id uint32, optionalHeight ...int) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfoById(c.ctxFunc(), &api_pb.CoinIdRequest{Height: optionalInt(optionalHeight), Id: id})
}

// CoinInfo returns information about coin symbol. Note: this method does not return information about base coins (MNT and BIP).
func (c *Client) CoinInfo(symbol string, optionalHeight ...int) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfo(c.ctxFunc(), &api_pb.CoinInfoRequest{Height: optionalInt(optionalHeight), Symbol: symbol})
}

// EstimateCoinSymbolBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinSymbolBuy(coinToSell, coinToBuy string, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToBuy: valueToBuy})
}

// EstimateCoinSymbolSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinSymbolSell(coinToBuy, coinToSell string, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToBuy})
}

// EstimateCoinSymbolSellAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinSymbolSellAll(coinToBuy, coinToSell string, valueToBuy string, gasPrice int, optionalHeight ...int) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinToSell{CoinToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinToBuy{CoinToBuy: coinToBuy}, ValueToSell: valueToBuy, GasPrice: uint64(gasPrice)})
}

// EstimateCoinIDBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinIDBuy(coinToSell, coinToBuy uint32, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinBuyRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinBuyRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToBuy: valueToBuy})
}

// EstimateCoinIDSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinIDSell(coinToBuy, coinToSell uint32, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToBuy})
}

// EstimateCoinSellIDAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinIDSellAll(coinToBuy, coinToSell uint32, valueToBuy string, gasPrice int, optionalHeight ...int) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), Sell: &api_pb.EstimateCoinSellAllRequest_CoinIdToSell{CoinIdToSell: coinToSell}, Buy: &api_pb.EstimateCoinSellAllRequest_CoinIdToBuy{CoinIdToBuy: coinToBuy}, ValueToSell: valueToBuy, GasPrice: uint64(gasPrice)})
}

// EstimateTxCommission return estimate of encoding transaction.
func (c *Client) EstimateTxCommission(tx string, optionalHeight ...int) (*api_pb.EstimateTxCommissionResponse, error) {
	return c.grpcClient.EstimateTxCommission(c.ctxFunc(), &api_pb.EstimateTxCommissionRequest{Height: optionalInt(optionalHeight), Tx: tx})
}

// Events returns events at given height.
func (c *Client) Events(optionalHeight ...int) (*api_pb.EventsResponse, error) {
	return c.grpcClient.Events(c.ctxFunc(), &api_pb.EventsRequest{Height: optionalInt(optionalHeight)})
}

// MaxGas returns current max gas.
func (c *Client) MaxGas(optionalHeight ...int) (*api_pb.MaxGasResponse, error) {
	return c.grpcClient.MaxGas(c.ctxFunc(), &api_pb.MaxGasRequest{Height: optionalInt(optionalHeight)})
}

// MinGasPrice returns current min gas price.
func (c *Client) MinGasPrice() (*api_pb.MinGasPriceResponse, error) {
	return c.grpcClient.MinGasPrice(c.ctxFunc(), &empty.Empty{})
}

// MissedBlocks returns missed blocks by validator public key.
func (c *Client) MissedBlocks(publicKey string, optionalHeight ...int) (*api_pb.MissedBlocksResponse, error) {
	return c.grpcClient.MissedBlocks(c.ctxFunc(), &api_pb.MissedBlocksRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

// NetInfo returns network info
func (c *Client) NetInfo() (*api_pb.NetInfoResponse, error) {
	return c.grpcClient.NetInfo(c.ctxFunc(), &empty.Empty{})
}

// SendTransaction returns the result of sending signed tx.
// To ensure that transaction was successfully committed to the blockchain,
// you need to find the transaction by the hash and ensure that the status code equals to 0.
func (c *Client) SendTransaction(tx string) (*api_pb.SendTransactionResponse, error) {
	return c.grpcClient.SendTransaction(c.ctxFunc(), &api_pb.SendTransactionRequest{Tx: tx})
}

// Transaction returns transaction info.
func (c *Client) Transaction(hash string) (*api_pb.TransactionResponse, error) {
	return c.grpcClient.Transaction(c.ctxFunc(), &api_pb.TransactionRequest{Hash: hash})
}

// Transactions return transactions by query.
func (c *Client) Transactions(query string, page, perPage int) (*api_pb.TransactionsResponse, error) {
	return c.grpcClient.Transactions(c.ctxFunc(), &api_pb.TransactionsRequest{Query: query, Page: int32(page), PerPage: int32(perPage)})
}

// UnconfirmedTxs returns unconfirmed transactions.
func (c *Client) UnconfirmedTxs(limit ...int) (*api_pb.UnconfirmedTxsResponse, error) {
	return c.grpcClient.UnconfirmedTxs(c.ctxFunc(), &api_pb.UnconfirmedTxsRequest{Limit: int32(optionalInt(limit))})
}

// Validators returns list of active validators.
func (c *Client) Validators(page, perPage int, limit ...int) (*api_pb.ValidatorsResponse, error) {
	return c.grpcClient.Validators(c.ctxFunc(), &api_pb.ValidatorsRequest{Height: optionalInt(limit), Page: int32(page), PerPage: int32(perPage)})
}

// WaitList returns the list of address stakes in waitlist.
func (c *Client) WaitList(publicKey, address string, limit ...int) (*api_pb.WaitListResponse, error) {
	return c.grpcClient.WaitList(c.ctxFunc(), &api_pb.WaitListRequest{Height: optionalInt(limit), PublicKey: publicKey, Address: address})
}

// Subscribe returns a subscription for events by query.
func (c *Client) Subscribe(query string) (api_pb.ApiService_SubscribeClient, error) {
	return c.grpcClient.Subscribe(c.ctxFunc(), &api_pb.SubscribeRequest{Query: query})
}

// Frozen returns frozen balance.
func (c *Client) Frozen(address, coin string) (*api_pb.FrozenResponse, error) {
	return c.grpcClient.Frozen(c.ctxFunc(), &api_pb.FrozenRequest{Address: address, Coin: coin})
}

func optionalInt(height []int) uint64 {
	if len(height) == 1 {
		return uint64(height[0])
	}
	return 0
}
