package grpc_client

import (
	"context"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type client struct {
	GrpcClient api_pb.ApiServiceClient
	ctx        func() context.Context
}

func (c *client) Halts(height int) (*api_pb.HaltsResponse, error) {
	return c.GrpcClient.Halts(c.ctx(), &api_pb.HaltsRequest{Height: uint64(height)})
}

func (c *client) Genesis() (*api_pb.GenesisResponse, error) {
	return c.GrpcClient.Genesis(c.ctx(), &empty.Empty{})
}

func (c *client) Status() (*api_pb.StatusResponse, error) {
	return c.GrpcClient.Status(c.ctx(), &empty.Empty{})
}

func (c *client) Address(address string, optionalHeight ...int) (*api_pb.AddressResponse, error) {
	return c.GrpcClient.Address(c.ctx(), &api_pb.AddressRequest{Height: optionalInt(optionalHeight), Address: address})
}

func (c *client) Addresses(addresses []string, optionalHeight ...int) (*api_pb.AddressesResponse, error) {
	return c.GrpcClient.Addresses(c.ctx(), &api_pb.AddressesRequest{Addresses: addresses, Height: optionalInt(optionalHeight)})
}

func (c *client) Block(height int) (*api_pb.BlockResponse, error) {
	return c.GrpcClient.Block(c.ctx(), &api_pb.BlockRequest{Height: uint64(height)})
}

func (c *client) CoinInfo(symbol string, optionalHeight ...int) (*api_pb.CoinInfoResponse, error) {
	return c.GrpcClient.CoinInfo(c.ctx(), &api_pb.CoinInfoRequest{Height: optionalInt(optionalHeight), Symbol: symbol})
}

func (c *client) Candidate(publicKey string, optionalHeight ...int) (*api_pb.CandidateResponse, error) {
	return c.GrpcClient.Candidate(c.ctx(), &api_pb.CandidateRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

func (c *client) Candidates(includeStakes bool, optionalHeight ...int) (*api_pb.CandidatesResponse, error) {
	return c.GrpcClient.Candidates(c.ctx(), &api_pb.CandidatesRequest{Height: optionalInt(optionalHeight), IncludeStakes: includeStakes})
}

func (c *client) Events(optionalHeight ...int) (*api_pb.EventsResponse, error) {
	return c.GrpcClient.Events(c.ctx(), &api_pb.EventsRequest{Height: optionalInt(optionalHeight)})
}

func (c *client) MaxGas(optionalHeight ...int) (*api_pb.MaxGasResponse, error) {
	return c.GrpcClient.MaxGas(c.ctx(), &api_pb.MaxGasRequest{Height: optionalInt(optionalHeight)})
}

func (c *client) MinGasPrice() (*api_pb.MinGasPriceResponse, error) {
	return c.GrpcClient.MinGasPrice(c.ctx(), &empty.Empty{})
}

func (c *client) MissedBlocks(publicKey string, optionalHeight ...int) (*api_pb.MissedBlocksResponse, error) {
	return c.GrpcClient.MissedBlocks(c.ctx(), &api_pb.MissedBlocksRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

func (c *client) NetInfo() (*api_pb.NetInfoResponse, error) {
	return c.GrpcClient.NetInfo(c.ctx(), &empty.Empty{})
}

func (c *client) SendTransaction(tx string) (*api_pb.SendTransactionResponse, error) {
	return c.GrpcClient.SendGetTransaction(c.ctx(), &api_pb.SendTransactionRequest{Tx: tx})
}

func (c *client) Transaction(hash string) (*api_pb.TransactionResponse, error) {
	return c.GrpcClient.Transaction(c.ctx(), &api_pb.TransactionRequest{Hash: hash})
}

func (c *client) Transactions(query string, page, perPage int) (*api_pb.TransactionsResponse, error) {
	return c.GrpcClient.Transactions(c.ctx(), &api_pb.TransactionsRequest{Query: query, Page: int32(page), PerPage: int32(perPage)})
}

func (c *client) UnconfirmedTxs(limit ...int) (*api_pb.UnconfirmedTxsResponse, error) {
	return c.GrpcClient.UnconfirmedTxs(c.ctx(), &api_pb.UnconfirmedTxsRequest{Limit: int32(optionalInt(limit))})
}

func (c *client) Validators(page, perPage int, limit ...int) (*api_pb.ValidatorsResponse, error) {
	return c.GrpcClient.Validators(c.ctx(), &api_pb.ValidatorsRequest{Height: optionalInt(limit), Page: int32(page), PerPage: int32(perPage)})
}

func (c *client) Subscribe(query string) (api_pb.ApiService_SubscribeClient, error) {
	return c.GrpcClient.Subscribe(c.ctx(), &api_pb.SubscribeRequest{Query: query})
}

func (c *client) EstimateCoinBuy(coinToSell, coinToBuy, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.GrpcClient.EstimateCoinBuy(c.ctx(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), CoinToSell: coinToSell, CoinToBuy: coinToBuy, ValueToBuy: valueToBuy})
}

func (c *client) EstimateCoinSell(coinToBuy, coinToSell, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinSellResponse, error) {
	return c.GrpcClient.EstimateCoinSell(c.ctx(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), CoinToBuy: coinToBuy, CoinToSell: coinToSell, ValueToSell: valueToBuy})
}

func (c *client) EstimateCoinSellAll(coinToBuy, coinToSell, valueToBuy string, gasPrice int, optionalHeight ...int) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.GrpcClient.EstimateCoinSellAll(c.ctx(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), CoinToBuy: coinToBuy, CoinToSell: coinToSell, ValueToSell: valueToBuy, GasPrice: uint64(gasPrice)})
}

func (c *client) EstimateTxCommission(tx string, optionalHeight ...int) (*api_pb.EstimateTxCommissionResponse, error) {
	return c.GrpcClient.EstimateTxCommission(c.ctx(), &api_pb.EstimateTxCommissionRequest{Height: optionalInt(optionalHeight), Tx: tx})
}

func New(address string) *client {
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &client{GrpcClient: api_pb.NewApiServiceClient(clientConn), ctx: context.Background}
}

func optionalInt(height []int) uint64 {
	if len(height) == 1 {
		return uint64(height[0])
	}
	return 0
}
