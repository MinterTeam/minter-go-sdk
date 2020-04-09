package grpc_client

import (
	"context"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type Client struct {
	grpcClient api_pb.ApiServiceClient
	ctxFunc    func() context.Context
}

func (c *Client) GRPCClient() api_pb.ApiServiceClient {
	return c.grpcClient
}

func (c *Client) Halts(height int) (*api_pb.HaltsResponse, error) {
	return c.grpcClient.Halts(c.ctxFunc(), &api_pb.HaltsRequest{Height: uint64(height)})
}

func (c *Client) Genesis() (*api_pb.GenesisResponse, error) {
	return c.grpcClient.Genesis(c.ctxFunc(), &empty.Empty{})
}

func (c *Client) Status() (*api_pb.StatusResponse, error) {
	return c.grpcClient.Status(c.ctxFunc(), &empty.Empty{})
}

func (c *Client) Address(address string, optionalHeight ...int) (*api_pb.AddressResponse, error) {
	return c.grpcClient.Address(c.ctxFunc(), &api_pb.AddressRequest{Height: optionalInt(optionalHeight), Address: address})
}

func (c *Client) Addresses(addresses []string, optionalHeight ...int) (*api_pb.AddressesResponse, error) {
	return c.grpcClient.Addresses(c.ctxFunc(), &api_pb.AddressesRequest{Addresses: addresses, Height: optionalInt(optionalHeight)})
}

func (c *Client) Block(height int) (*api_pb.BlockResponse, error) {
	return c.grpcClient.Block(c.ctxFunc(), &api_pb.BlockRequest{Height: uint64(height)})
}

func (c *Client) CoinInfo(symbol string, optionalHeight ...int) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfo(c.ctxFunc(), &api_pb.CoinInfoRequest{Height: optionalInt(optionalHeight), Symbol: symbol})
}

func (c *Client) Candidate(publicKey string, optionalHeight ...int) (*api_pb.CandidateResponse, error) {
	return c.grpcClient.Candidate(c.ctxFunc(), &api_pb.CandidateRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

func (c *Client) Candidates(includeStakes bool, optionalHeight ...int) (*api_pb.CandidatesResponse, error) {
	return c.grpcClient.Candidates(c.ctxFunc(), &api_pb.CandidatesRequest{Height: optionalInt(optionalHeight), IncludeStakes: includeStakes})
}

func (c *Client) Events(optionalHeight ...int) (*api_pb.EventsResponse, error) {
	return c.grpcClient.Events(c.ctxFunc(), &api_pb.EventsRequest{Height: optionalInt(optionalHeight)})
}

func (c *Client) MaxGas(optionalHeight ...int) (*api_pb.MaxGasResponse, error) {
	return c.grpcClient.MaxGas(c.ctxFunc(), &api_pb.MaxGasRequest{Height: optionalInt(optionalHeight)})
}

func (c *Client) MinGasPrice() (*api_pb.MinGasPriceResponse, error) {
	return c.grpcClient.MinGasPrice(c.ctxFunc(), &empty.Empty{})
}

func (c *Client) MissedBlocks(publicKey string, optionalHeight ...int) (*api_pb.MissedBlocksResponse, error) {
	return c.grpcClient.MissedBlocks(c.ctxFunc(), &api_pb.MissedBlocksRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

func (c *Client) NetInfo() (*api_pb.NetInfoResponse, error) {
	return c.grpcClient.NetInfo(c.ctxFunc(), &empty.Empty{})
}

func (c *Client) SendTransaction(tx string) (*api_pb.SendTransactionResponse, error) {
	return c.grpcClient.SendGetTransaction(c.ctxFunc(), &api_pb.SendTransactionRequest{Tx: tx})
}

func (c *Client) Transaction(hash string) (*api_pb.TransactionResponse, error) {
	return c.grpcClient.Transaction(c.ctxFunc(), &api_pb.TransactionRequest{Hash: hash})
}

func (c *Client) Transactions(query string, page, perPage int) (*api_pb.TransactionsResponse, error) {
	return c.grpcClient.Transactions(c.ctxFunc(), &api_pb.TransactionsRequest{Query: query, Page: int32(page), PerPage: int32(perPage)})
}

func (c *Client) UnconfirmedTxs(limit ...int) (*api_pb.UnconfirmedTxsResponse, error) {
	return c.grpcClient.UnconfirmedTxs(c.ctxFunc(), &api_pb.UnconfirmedTxsRequest{Limit: int32(optionalInt(limit))})
}

func (c *Client) Validators(page, perPage int, limit ...int) (*api_pb.ValidatorsResponse, error) {
	return c.grpcClient.Validators(c.ctxFunc(), &api_pb.ValidatorsRequest{Height: optionalInt(limit), Page: int32(page), PerPage: int32(perPage)})
}

func (c *Client) Subscribe(query string) (api_pb.ApiService_SubscribeClient, error) {
	return c.grpcClient.Subscribe(c.ctxFunc(), &api_pb.SubscribeRequest{Query: query})
}

func (c *Client) EstimateCoinBuy(coinToSell, coinToBuy, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), CoinToSell: coinToSell, CoinToBuy: coinToBuy, ValueToBuy: valueToBuy})
}

func (c *Client) EstimateCoinSell(coinToBuy, coinToSell, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), CoinToBuy: coinToBuy, CoinToSell: coinToSell, ValueToSell: valueToBuy})
}

func (c *Client) EstimateCoinSellAll(coinToBuy, coinToSell, valueToBuy string, gasPrice int, optionalHeight ...int) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), CoinToBuy: coinToBuy, CoinToSell: coinToSell, ValueToSell: valueToBuy, GasPrice: uint64(gasPrice)})
}

func (c *Client) EstimateTxCommission(tx string, optionalHeight ...int) (*api_pb.EstimateTxCommissionResponse, error) {
	return c.grpcClient.EstimateTxCommission(c.ctxFunc(), &api_pb.EstimateTxCommissionRequest{Height: optionalInt(optionalHeight), Tx: tx})
}

func New(address string) *Client {
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &Client{grpcClient: api_pb.NewApiServiceClient(clientConn), ctxFunc: context.Background}
}

func (c *Client) WithContextFunc(contextFunc func() context.Context) *Client {
	c.ctxFunc = contextFunc
	return c
}

func optionalInt(height []int) uint64 {
	if len(height) == 1 {
		return uint64(height[0])
	}
	return 0
}
