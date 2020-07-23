package grpc_client

import (
	"context"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"strconv"
)

type Client struct {
	grpcClient api_pb.ApiServiceClient
	ctxFunc    func() context.Context
}

// New API v2 client
func New(address string) (*Client, error) {
	clientConn, err := grpc.Dial(address,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor()),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{grpcClient: api_pb.NewApiServiceClient(clientConn), ctxFunc: context.Background}, nil
}

func (c *Client) WithContextFunc(contextFunc func(context.Context) func() context.Context) *Client {
	return &Client{grpcClient: c.grpcClient, ctxFunc: contextFunc(c.ctxFunc())}
}

func (c *Client) GRPCClient() api_pb.ApiServiceClient {
	return c.grpcClient
}

//Halts
func (c *Client) Halts(height int) (*api_pb.HaltsResponse, error) {
	return c.grpcClient.Halts(c.ctxFunc(), &api_pb.HaltsRequest{Height: uint64(height)})
}

//Genesis returns genesis file.
func (c *Client) Genesis() (*api_pb.GenesisResponse, error) {
	return c.grpcClient.Genesis(c.ctxFunc(), &empty.Empty{})
}

//Nonce returns next transaction number (nonce) of an address.
func (c *Client) Nonce(address string) (uint64, error) {
	status, err := c.Address(address)
	if err != nil {
		return 0, err
	}

	transactionsCount, err := strconv.Atoi(status.TransactionsCount)
	if err != nil {
		return 0, err
	}

	return uint64(transactionsCount) + 1, err
}

//Status returns node status including pubkey, latest block.
func (c *Client) Status() (*api_pb.StatusResponse, error) {
	return c.grpcClient.Status(c.ctxFunc(), &empty.Empty{})
}

//Address returns coins list, balance and transaction count of an address.
func (c *Client) Address(address string, optionalHeight ...int) (*api_pb.AddressResponse, error) {
	return c.grpcClient.Address(c.ctxFunc(), &api_pb.AddressRequest{Height: optionalInt(optionalHeight), Address: address})
}

//Addresses returns list of addresses.
func (c *Client) Addresses(addresses []string, optionalHeight ...int) (*api_pb.AddressesResponse, error) {
	return c.grpcClient.Addresses(c.ctxFunc(), &api_pb.AddressesRequest{Addresses: addresses, Height: optionalInt(optionalHeight)})
}

//Block returns block data at given height.
func (c *Client) Block(height int) (*api_pb.BlockResponse, error) {
	return c.grpcClient.Block(c.ctxFunc(), &api_pb.BlockRequest{Height: uint64(height)})
}

//Candidate returns candidate’s info by provided public_key. It will respond with 404 code if candidate is not found.
func (c *Client) Candidate(publicKey string, optionalHeight ...int) (*api_pb.CandidateResponse, error) {
	return c.grpcClient.Candidate(c.ctxFunc(), &api_pb.CandidateRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

//Candidates returns list of candidates.
func (c *Client) Candidates(includeStakes bool, optionalHeight ...int) (*api_pb.CandidatesResponse, error) {
	return c.grpcClient.Candidates(c.ctxFunc(), &api_pb.CandidatesRequest{Height: optionalInt(optionalHeight), IncludeStakes: includeStakes})
}

//CoinId returns information about coin ID. Note: this method does not return information about base coins (MNT and BIP).
func (c *Client) CoinId(id uint32, optionalHeight ...int) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinId(c.ctxFunc(), &api_pb.CoinIdRequest{Height: optionalInt(optionalHeight), Id: id})
}

//CoinInfo returns information about coin symbol. Note: this method does not return information about base coins (MNT and BIP).
func (c *Client) CoinInfo(symbol string, optionalHeight ...int) (*api_pb.CoinInfoResponse, error) {
	return c.grpcClient.CoinInfo(c.ctxFunc(), &api_pb.CoinInfoRequest{Height: optionalInt(optionalHeight), Symbol: symbol})
}

//EstimateCoinBuy return estimate of buy coin transaction.
func (c *Client) EstimateCoinBuy(coinToSell, coinToBuy uint32, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinBuyResponse, error) {
	return c.grpcClient.EstimateCoinBuy(c.ctxFunc(), &api_pb.EstimateCoinBuyRequest{Height: optionalInt(optionalHeight), CoinIdToSell: coinToSell, CoinIdToBuy: coinToBuy, ValueToBuy: valueToBuy})
}

//EstimateCoinSell return estimate of sell coin transaction.
func (c *Client) EstimateCoinSell(coinToBuy, coinToSell uint32, valueToBuy string, optionalHeight ...int) (*api_pb.EstimateCoinSellResponse, error) {
	return c.grpcClient.EstimateCoinSell(c.ctxFunc(), &api_pb.EstimateCoinSellRequest{Height: optionalInt(optionalHeight), CoinIdToBuy: coinToBuy, CoinIdToSell: coinToSell, ValueToSell: valueToBuy})
}

//EstimateCoinSellAll return estimate of sell all coin transaction.
func (c *Client) EstimateCoinSellAll(coinToBuy, coinToSell uint32, valueToBuy string, gasPrice int, optionalHeight ...int) (*api_pb.EstimateCoinSellAllResponse, error) {
	return c.grpcClient.EstimateCoinSellAll(c.ctxFunc(), &api_pb.EstimateCoinSellAllRequest{Height: optionalInt(optionalHeight), CoinIdToBuy: coinToBuy, CoinIdToSell: coinToSell, ValueToSell: valueToBuy, GasPrice: uint64(gasPrice)})
}

//EstimateTxCommissionTx return estimate of encoding transaction.
func (c *Client) EstimateTxCommissionTx(tx string, optionalHeight ...int) (*api_pb.EstimateTxCommissionResponse, error) {
	return c.grpcClient.EstimateTxCommission(c.ctxFunc(), &api_pb.EstimateTxCommissionRequest{Height: optionalInt(optionalHeight), Tx: tx})
}

//EstimateTxCommissionData return estimate of transaction data.
func (c *Client) EstimateTxCommissionData(gasCoin uint32, t api_pb.EstimateTxCommissionRequest_TransactionData_Type, payload []byte, mtxsCount int64, optionalHeight ...int) (*api_pb.EstimateTxCommissionResponse, error) {
	return c.grpcClient.EstimateTxCommission(c.ctxFunc(), &api_pb.EstimateTxCommissionRequest{
		Height: optionalInt(optionalHeight),
		Data: &api_pb.EstimateTxCommissionRequest_TransactionData{
			GasCoinId: gasCoin,
			Type:      t,
			Payload:   payload,
			Mtxs:      mtxsCount,
		},
	})
}

//Events returns events at given height.
func (c *Client) Events(optionalHeight ...int) (*api_pb.EventsResponse, error) {
	return c.grpcClient.Events(c.ctxFunc(), &api_pb.EventsRequest{Height: optionalInt(optionalHeight)})
}

//MaxGas returns current max gas.
func (c *Client) MaxGas(optionalHeight ...int) (*api_pb.MaxGasResponse, error) {
	return c.grpcClient.MaxGas(c.ctxFunc(), &api_pb.MaxGasRequest{Height: optionalInt(optionalHeight)})
}

//MinGasPrice returns current min gas price.
func (c *Client) MinGasPrice() (*api_pb.MinGasPriceResponse, error) {
	return c.grpcClient.MinGasPrice(c.ctxFunc(), &empty.Empty{})
}

//MissedBlocks returns missed blocks by validator public key.
func (c *Client) MissedBlocks(publicKey string, optionalHeight ...int) (*api_pb.MissedBlocksResponse, error) {
	return c.grpcClient.MissedBlocks(c.ctxFunc(), &api_pb.MissedBlocksRequest{Height: optionalInt(optionalHeight), PublicKey: publicKey})
}

//NetInfo returns network info
func (c *Client) NetInfo() (*api_pb.NetInfoResponse, error) {
	return c.grpcClient.NetInfo(c.ctxFunc(), &empty.Empty{})
}

//SendTransaction returns the result of sending signed tx.
//To ensure that transaction was successfully committed to the blockchain,
//you need to find the transaction by the hash and ensure that the status code equals to 0.
func (c *Client) SendTransaction(tx string) (*api_pb.SendTransactionResponse, error) {
	return c.grpcClient.SendTransaction(c.ctxFunc(), &api_pb.SendTransactionRequest{Tx: tx})
}

//Transaction returns transaction info.
func (c *Client) Transaction(hash string) (*api_pb.TransactionResponse, error) {
	return c.grpcClient.Transaction(c.ctxFunc(), &api_pb.TransactionRequest{Hash: hash})
}

//Transactions return transactions by query.
func (c *Client) Transactions(query string, page, perPage int) (*api_pb.TransactionsResponse, error) {
	return c.grpcClient.Transactions(c.ctxFunc(), &api_pb.TransactionsRequest{Query: query, Page: int32(page), PerPage: int32(perPage)})
}

//UnconfirmedTxs returns unconfirmed transactions.
func (c *Client) UnconfirmedTxs(limit ...int) (*api_pb.UnconfirmedTxsResponse, error) {
	return c.grpcClient.UnconfirmedTxs(c.ctxFunc(), &api_pb.UnconfirmedTxsRequest{Limit: int32(optionalInt(limit))})
}

//Validators returns list of active validators.
func (c *Client) Validators(page, perPage int, limit ...int) (*api_pb.ValidatorsResponse, error) {
	return c.grpcClient.Validators(c.ctxFunc(), &api_pb.ValidatorsRequest{Height: optionalInt(limit), Page: int32(page), PerPage: int32(perPage)})
}

//Subscribe returns a subscription for events by query.
func (c *Client) Subscribe(query string) (api_pb.ApiService_SubscribeClient, error) {
	return c.grpcClient.Subscribe(c.ctxFunc(), &api_pb.SubscribeRequest{Query: query})
}

func optionalInt(height []int) uint64 {
	if len(height) == 1 {
		return uint64(height[0])
	}
	return 0
}
