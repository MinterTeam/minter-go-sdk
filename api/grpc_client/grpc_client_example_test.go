package grpc_client_test

import (
	"io"
	"log"
	"math/big"
	"time"

	"github.com/MinterTeam/minter-go-sdk/v2/api"
	"github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ExampleClient_Subscribe_newBlock() {
	client, _ := grpc_client.New("localhost:8842")

	subscribeClient, err := client.Subscribe(api.QueryEvent(api.EventNewBlock))
	if err != nil {
		log.Fatal(err)
	}

	defer subscribeClient.CloseSend()

	for {
		recv, err := subscribeClient.Recv()
		if err == io.EOF {
			log.Fatal(err)
		}
		if code := status.Code(err); code != codes.OK {
			if code == codes.DeadlineExceeded || code == codes.Canceled {
				log.Fatal(errors.Wrap(err, "event subscription error in node configuration"))
			}
			log.Fatal(err)
		}

		marshal, err := client.Marshal(recv)
		if err != nil {
			log.Fatal(err)
		}

		parse, err := api.SubscribeNewBlockParse(marshal)

		block, err := client.WithCallOption(
			grpc_retry.WithCodes(codes.NotFound),
			grpc_retry.WithBackoff(grpc_retry.BackoffExponential(time.Second/2)),
			grpc_retry.WithMax(5),
		).BlockExtended(parse.Data.Block.Header.Height, true, false)
		if err != nil {
			log.Fatal(err)
		}
		for _, tx := range block.Transactions {
			txDecode, err := transaction.Decode(tx.RawTx)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%#v", txDecode.Data())
		}
	}
}

func ExampleClient_SendTransaction() {
	client, _ := grpc_client.New("localhost:8842")
	coinID, _ := client.CoinID("SYMBOL")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	nonce, _ := client.Nonce(w.Address)
	data := transaction.NewSendData().SetCoin(coinID).SetValue(transaction.BipToPip(big.NewInt(1))).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(nonce).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()

	subscribeClient, _ := client.Subscribe(api.QueryHash(hash))
	defer subscribeClient.CloseSend()

	res, err := client.WithCallOption(
		grpc_retry.WithCodes(codes.FailedPrecondition),
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(time.Second/2)),
		grpc_retry.WithMax(5),
	).SendTransaction(encode)
	if err != nil {
		_, _, _ = client.ErrorBody(err)
	}
	if res.Code != 0 {
		panic(res.Log)
	}

	{
		recv, err := subscribeClient.Recv()
		if err == io.EOF {
			return
		}
		if code := status.Code(err); code != codes.OK {
			if code == codes.DeadlineExceeded || code == codes.Canceled {
				return
			}
			panic(err)
		}

		marshal, _ := client.Marshal(recv)
		findedTx, _ := api.SubscribeNewTxToTx(marshal)
		_, _ = findedTx.GetTransaction(), findedTx.Data().(*transaction.SendData)
	}
	// or
	{
		response, _ := client.WithCallOption(
			grpc_retry.WithCodes(codes.NotFound),
			grpc_retry.WithBackoff(grpc_retry.BackoffExponential(time.Second/2)),
			grpc_retry.WithMax(5),
		).Transaction(hash)
		_, _ = client.Marshal(response)
		sendData := new(api_pb.SendData)
		_ = response.Data.UnmarshalTo(sendData)
		_, _ = client.Marshal(sendData)
	}
}
