![Minter Logo](https://github.com/MinterTeam/minter-go-sdk/raw/v2/minter-logo.svg?sanitize=true)

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/MinterTeam/minter-go-sdk/v2)

## Overview

This is a pure Go SDK for working with **Minter** blockchain

## Table of contents

* [Installing](#installing)
* [Minter API](#minter-api)
    - [Using API v2](#using-api-v2)
    - [Using gRPC](#using-grpc)
* [Using Transactions](#using-transactions)
    - [Sign transaction](#sign-transaction)
        - [Single signature](#single-signature)
        - [Multi signature](#multi-signatures)
* [Minter Wallet](#minter-wallet)

### Installing

```bash
go get github.com/MinterTeam/minter-go-sdk/v2@latest
```

### Minter API

Minter blockchain nodes have built-in API with grpc and http + websocket interfaces. This package will help you with
forming requests, parsing responses, as well as when working with other API entities.

#### Using API v2

Package _http_client_ package implements the API v2 methods usage interface.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/api/http_client?tab=doc)

```go
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

func main() {
	client, _ := http_client.New("http://localhost:8843/v2")
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
		time.Sleep(5*time.Second)
		response, _ := client.Transaction(hash)
		_, _ = client.Marshal(response)
		sendData := new(models.SendData)
		_ = response.Data.UnmarshalTo(sendData)
		_, _ = client.Marshal(sendData)
	}
}
```

#### Using gRPC

Package _grpc_client_ package implements the gRPC methods usage interface.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client?tab=doc)

```go
package grpc_client_test

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api"
	"github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"math/big"
	"time"
)

func main() {
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
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1*time.Second)),
		grpc_retry.WithMax(4),
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
		time.Sleep(5*time.Second)
		response, _ := client.Transaction(hash)
		_, _ = client.Marshal(response)
		sendData := new(api_pb.SendData)
		_ = response.Data.UnmarshalTo(sendData)
		_, _ = client.Marshal(sendData)
	}
}
```
### Using Transactions

Package _transaction_ is a guide and assistant in working with fields and types of transactions, creating signatures,
encoding and decoding with RLP.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/transaction?tab=doc)

#### Sign transaction

Returns a signed tx.

⚠️ After sending the transaction, to make sure that the transaction was successfully committed on the blockchain, you
need to find the transaction by hash and make sure that the status code is 0.

##### Single signature

###### Example

```go
tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(
  transaction.NewSendData().
    SetCoin(0).
    SetValue(transaction.BipToPip(big.NewInt(1))).
    MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
)

signedTransaction, _ := tx.SetNonce(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")

encode, _ := signedTransaction.Encode()
```

##### Multi signatures

###### Example

```go
tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(
  transaction.NewSendData().
  SetCoin(0).
  SetValue(transaction.BipToPip(big.NewInt(1))).
  MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")
)

signedTx, _ := tx.SetNonce(1).SetMultiSignatureType().Sign(
  multisigAddress,
  "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
  "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
  "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba",
)

encode, _ := signedTx.Encode()
```

You can transfer the transaction to the remaining addresses

```go
signedTx1, _ := tx.Sign(msigAddress, privateKey1)
encode, _ := signedTx.Encode()
// transfer encode transaction
signedTx1, _ = transaction.Decode(encode)
// and continue its signature by the remaining participants
signedTx12, _ := signedTx1.Sign(msigAddress, privateKey2)
signedTx123, _ := signedTx12.Sign(msigAddress, privateKey3)
encode, _ := signedTx123.Encode()
```

You can collect all signatures in one place without revealing the private key

```go
signedTx1, _ := tx.Clone().Sign(msigAddress, "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
signedTx2, _ := tx.Clone().Sign(msigAddress, "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7")
signedTx3, _ := tx.Clone().Sign(msigAddress, "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
simpleSignatureData1, _ := signedTx1.SingleSignatureData()
simpleSignatureData2, _ := signedTx2.SingleSignatureData()
simpleSignatureData3, _ := signedTx3.SingleSignatureData()
signedTransaction, _ := tx.Clone().Sign(msigAddress)
signedTx123, _ := signedTransaction.AddSignature(simpleSignatureData1, simpleSignatureData2, simpleSignatureData3)

encode, _ := signedTx123.Encode()
```

### Minter Wallet

Package _wallet_ is a guide and assistant in working with addresses, mnemonic and initial phrases, private and public
keys.

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/wallet?tab=doc)
