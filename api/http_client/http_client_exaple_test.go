package http_client_test

import (
	"context"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"io"
	"math/big"
	"strings"
)

func Example() {
	client, _ := http_client.NewConcise("http://localhost:8843/v2")
	_ = client.CheckVersion("1.2", true)
	coinID, _ := client.CoinID("SYMBOL")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	data := transaction.NewSendData().SetCoin(coinID).SetValue(transaction.BipToPip(big.NewInt(1))).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()
	subscribeClient, _ := client.Subscribe(context.Background(), fmt.Sprintf("tx.hash = '%s'", strings.ToUpper(hash[2:])))
	defer subscribeClient.CloseSend()

	res, _ := client.SendTransaction(encode)
	if res.Code != 0 {
		panic(res.Log)
	}

	_, err := subscribeClient.Recv()
	if err == io.EOF {
		return
	}
	if err == context.Canceled || err == context.DeadlineExceeded {
		return
	}
	if err != nil {
		panic(err)
	}

	response, _ := client.Transaction(hash)
	sendData := new(models.SendData)
	_ = response.Data.UnmarshalTo(sendData)
}
