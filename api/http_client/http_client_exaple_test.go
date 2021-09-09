package http_client_test

import (
	"context"
	"github.com/MinterTeam/minter-go-sdk/v2/api"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"io"
	"math/big"
	"time"
)

func ExampleClient_SendTransaction() {
	client, _ := http_client.New("http://localhost:8843/v2")
	coinID, _ := client.CoinID("SYMBOL")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	data := transaction.NewSendData().SetCoin(coinID).SetValue(transaction.BipToPip(big.NewInt(1))).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(1).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()

	subscribeClient, _ := client.Subscribe(api.QueryHash(hash))
	defer subscribeClient.CloseSend()

	res, err := client.SendTransaction(encode)
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
		if err == context.Canceled || err == context.DeadlineExceeded {
			return
		}
		if err != nil {
			panic(err)
		}

		marshal, _ := client.Marshal(recv)
		findedTx, _ := api.SubscribeNewTxToTx(marshal)
		_, _ = findedTx.GetTransaction(), findedTx.Data().(*transaction.SendData)
	}
	// or
	{
		time.Sleep(5 * time.Second)
		response, _ := client.Transaction(hash)
		_, _ = client.Marshal(response)
		sendData := new(models.SendData)
		_ = response.Data.UnmarshalTo(sendData)
		_, _ = client.Marshal(sendData)
	}
}
