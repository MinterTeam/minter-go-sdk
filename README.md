![Minter Logo](https://github.com/MinterTeam/minter-go-sdk/raw/master/minter-logo.svg?sanitize=true)

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/MinterTeam/minter-go-sdk)

## About

This is a pure Go SDK for working with **Minter** blockchain

* [Minter API](#minter-api)
* [Using API v2](#using-api-v2)
* [Using gRPC](#using-grpc)
* [Using Transactions](#using-transactions)
	- [Sign transaction](#sign-transaction)
        - [Single signature](#single-signature)
        - [Multi signature](#multi-signatures)
* [Minter Wallet](#minter-wallet)

## Installing

```bash
go get github.com/MinterTeam/minter-go-sdk
```

## Minter API

* v1 - Deprecated.
    - [Doc](https://github.com/MinterTeam/minter-go-sdk/tree/v1.1.1#using-minterapi)
    - [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk@v1.1.6/api?tab=doc)
* v2 - RESTful API and gRPC. 
    - [Github Source Code](https://github.com/MinterTeam/node-grpc-gateway)
    - [ReDoc](https://minterteam.github.io/node-gateway-api-v2-doc/)
    - [![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2FMinterTeam%2Fnode-grpc-gateway%2Fmaster%2Fdocs%2Fapi.swagger.json)](https://minterteam.github.io/minter-api-v2-docs/)

## Using API v2

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/api/v2?tab=doc)

## Using gRPC

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/api/grpc_client?tab=doc)

## Using Transactions

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/transaction?tab=doc)

### Sign transaction

Returns a signed tx.

⚠️ After sending the transaction, to make sure that the transaction was successfully committed on the blockchain, you need to find the transaction by hash and make sure that the status code is 0.

#### Single signature

##### Example

```go
var data transaction.DataInterface
// data = ...
tx, _ := transaction.NewBuilder(TestNetChainID).NewTransaction(data)
tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(symbolMNT).SetSignatureType(transaction.SignatureTypeSingle)
signedTx, _ := tx.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
encode, _ := signedTx.Encode()
```

#### Multi signatures

##### Example

```go
var data transaction.DataInterface
var dataMultisig *transaction.CreateMultisigData
// data = ...
tx, _ := transaction.NewBuilder(TestNetChainID).NewTransaction(data)
tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(symbolMNT).SetMultiSignatureType(transaction.SignatureTypeMulti)
dataMultisig = transaction.NewCreateMultisigData().
		MustAddSigData("Mxee81347211c72524338f9680072af90744333143", 1).
		MustAddSigData("Mxee81347211c72524338f9680072af90744333145", 3).
		MustAddSigData("Mxee81347211c72524338f9680072af90744333144", 5).
		SetThreshold(7)
msigAddress := dataMultisig.AddressString()
signedTx, _ := tx.Sign(msigAddress, privateKey1, privateKey2, privateKey3)
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
msigAddress := "Mx0023aa9371e0779189ef5a7434456fc21a938945"
signedTx1, _ := tx.Clone().Sign(msigAddress, "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
signedTx2, _ := tx.Clone().Sign(msigAddress, "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7")
signedTx3, _ := tx.Clone().Sign(msigAddress, "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
SingleSignatureData1, _ := signedTx1.SingleSignatureData()
SingleSignatureData2, _ := signedTx2.SingleSignatureData()
SingleSignatureData3, _ := signedTx3.SingleSignatureData()
signedTransaction, _ := tx.Clone().Sign(msigAddress)
signedTx123, _ := signedTransaction.AddSignature(SingleSignatureData1, SingleSignatureData2, SingleSignatureData3)
encode, _ := signedTx123.Encode()
```

### Minter Wallet

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/wallet?tab=doc)
