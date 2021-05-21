package http_client_test

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client"
	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
)

func Example() {
	client, _ := http_client.New("http://localhost:8843/v2")
	_ = client.CheckVersion("2.2", true)
	coinID, _ := client.CoinID("SYMBOL")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	data := transaction.NewSendData().SetCoin(coinID).SetValue(transaction.BipToPip(big.NewInt(1))).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(1).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()
	subscribeClient, _ := client.Subscribe(fmt.Sprintf("tx.hash = '%s'", strings.ToUpper(hash[2:])))
	defer subscribeClient.CloseSend()

	res, err := client.SendTransaction(encode)
	if err != nil {
		_, _, _ = http_client.ErrorBody(err)
	}
	if res.Code != 0 {
		panic(res.Log)
	}

	_, err = subscribeClient.Recv()
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
	_, _ = http_client.Marshal(response)
	sendData := new(models.SendData)
	_ = response.Data.UnmarshalTo(sendData)
}
