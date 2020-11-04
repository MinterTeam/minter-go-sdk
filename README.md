![Minter Logo](https://github.com/MinterTeam/minter-go-sdk/raw/v1.2/minter-logo.svg?sanitize=true)

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/MinterTeam/minter-go-sdk)

# The minter-go-sdk v1 is in maintenance mode, you are encouraged to migrate to [v2](https://github.com/MinterTeam/minter-go-sdk/tree/v2), which will have a stable release.
# Check out our [v2 migration guide](https://www.minter.network/ru/docs/1-2-migrate).

## About

This is a pure Go SDK for working with **Minter** blockchain

* [Minter Api](#using-minter-api)
* [Minter gRPC](#using-minter-grpc)
* [Minter SDK](#using-mintersdk)
	- [Sign transaction](#sign-transaction)
        - [Single signature](#single-signature)
        - [Multi signature](#multi-signatures)
    - [Create transaction](#create-transaction)
        - [Send](#send-transaction)
        - [SellCoin](#sell-coin-transaction)
        - [SellAllCoin](#sell-all-coin-transaction)
        - [BuyCoin](#buy-coin-transaction)
        - [CreateCoin](#create-coin-transaction)
        - [DeclareCandidacy](#declare-candidacy-transaction)
        - [Delegate](#delegate-transaction)
        - [SetCandidateOn](#set-candidate-online-transaction)
        - [SetCandidateOff](#set-candidate-offline-transaction)
        - [RedeemCheck](#redeem-check-transaction)
        - [Unbond](#unbond-transaction)
        - [Multisend](#multisend-transaction)
        - [EditCandidate](#edit-candidate-transaction)
	- [Get fee of transaction](#get-fee-of-transaction)
	- [Get hash of transaction](#get-hash-of-transaction)
	- [Decode Transaction](#decode-transaction)
	- [Minter Deep Links](#minter-deep-links)
	- [Minter Check](#minter-check)
* [Minter Wallet](#minter-wallet)

## Installing

```bash
go get github.com/MinterTeam/minter-go-sdk
```

## Using Minter API

* v1 - Will be deprecated soon.
    - [Doc](https://github.com/MinterTeam/minter-go-sdk/tree/v1.1.1#using-minterapi)
    - [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/api?tab=doc)
* v2 - RESTful API and gRPC. 
    - [Github Source Code](https://github.com/MinterTeam/node-grpc-gateway)
    - [ReDoc](https://minterteam.github.io/node-gateway-api-v2-doc/)
    - [![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https%3A%2F%2Fraw.githubusercontent.com%2FMinterTeam%2Fnode-grpc-gateway%2Fmaster%2Fapi.swagger.json)](https://minterteam.github.io/minter-api-v2-docs/)
    - [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/node-grpc-gateway/api_pb?tab=doc)

## Using Minter gRPC

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/api/grpc_client?tab=doc)

## Using MinterSDK

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
simpleSignatureData1, _ := signedTx1.SimpleSignatureData()
simpleSignatureData2, _ := signedTx2.SimpleSignatureData()
simpleSignatureData3, _ := signedTx3.SimpleSignatureData()
signedTransaction, _ := tx.Clone().Sign(msigAddress)
signedTx123, _ := signedTransaction.AddSignature(simpleSignatureData1, simpleSignatureData2, simpleSignatureData3)
encode, _ := signedTx123.Encode()
```

### Create transaction

[](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/transaction?tab=doc)

#### Send transaction

Transaction for sending arbitrary coin.

Coin - Symbol of a coin. To - Recipient address in Minter Network. Value - Amount of Coin to send.

##### Example

```go
value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)) // 1000000000000000000
address := "Mx1b685a7c1e78726c48f619c497a07ed75fe00483"
symbolMNT := "MNT"
data, _ := transaction.NewSendData().SetCoin(symbolMNT).SetValue(value).SetTo(address)
```

#### Sell coin transaction

Transaction for selling one coin (owned by sender) in favour of another coin in a system.

CoinToSell - Symbol of a coin to give. ValueToSell - Amount of CoinToSell to give. CoinToBuy - Symbol of a coin to get. MinimumValueToBuy - Minimum value of coins to get.

##### Example

```go
data := transaction.NewSellCoinData().
	SetCoinToSell("MNT").
	SetValueToSell(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
	SetCoinToBuy("TEST").
	SetMinimumValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
```

#### Sell all coin transaction

Transaction for selling one coin (owned by sender) in favour of another coin in a system.

CoinToSell - Symbol of a coin to give. ValueToSell - Amount of CoinToSell to give. CoinToBuy - Symbol of a coin to get. MinimumValueToBuy - Minimum value of coins to get.

##### Example

```go
data := transaction.NewSellAllCoinData().
	SetCoinToSell("MNT").
	SetCoinToBuy("TEST").
	SetMinimumValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
```

#### Buy coin transaction

Transaction for buy a coin paying another coin (owned by sender).

CoinToBuy - Symbol of a coin to get. ValueToBuy - Amount of CoinToBuy to get. CoinToSell - Symbol of a coin to give. MaximumValueToSell - Maximum value of coins to sell.

##### Example

```go
data := transaction.NewBuyCoinData().
	SetCoinToBuy("TEST").
	SetValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
	SetCoinToSell("MNT").
	SetMaximumValueToSell(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
```

#### Create coin transaction

Transaction for creating new coin in a system.

Name - Name of a coin. Arbitrary string up to 64 letters length. Symbol - Symbol of a coin. Must be unique, alphabetic, uppercase, 3 to 10 symbols length. InitialAmount - Amount of coins to issue. Issued coins will be available to sender account. InitialReserve - Initial reserve in BIP's. ConstantReserveRatio - CRR, uint, should be from 10 to 100.

##### Example

```go
data := transaction.NewCreateCoinData().
	SetName("SUPER TEST").
	SetSymbol("SPRTEST").
	SetInitialAmount(big.NewInt(0).Mul(big.NewInt(100), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
	SetInitialReserve(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
	SetConstantReserveRatio(10)
```

#### Declare candidacy transaction

Transaction for declaring new validator candidacy.

Address - Address of candidate in Minter Network. This address would be able to control candidate. Also all rewards will be sent to this address. PubKey - Public key of a validator. Commission - Commission (from 0 to 100) from rewards which delegators will pay to validator. Coin - Symbol of coin to stake. Stake - Amount of coins to stake.

##### Example

```go
data, _ := transaction.NewDeclareCandidacyData().
	MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
	SetCommission(10).
	SetCoin("MNT").
	SetStake(big.NewInt(0).Mul(big.NewInt(5), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
	SetAddress("Mx9f7fd953c2c69044b901426831ed03ee0bd0597a")
```

#### Delegate transaction

Transaction for delegating funds to validator.

PubKey - Public key of a validator. Coin - Symbol of coin to stake. Stake - Amount of coins to stake.

##### Example

```go
data := transaction.NewDelegateData().
	MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
	SetCoin("MNT").
	SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
```

#### Unbond transaction

Transaction for unbonding funds from validator's stake.

PubKey - Public key of a validator. Coin - Symbol of coin to stake. Stake - Amount of coins to stake.

##### Example

```go
data := transaction.NewUnbondData().
	MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
	SetCoin("MNT").
	SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
```

#### Redeem check transaction

Transaction for redeeming a check.

RawCheck - Raw check received from sender. Proof - Proof of owning a check.

Note that maximum GasPrice is limited to 1 to prevent fraud, because GasPrice is set by redeem tx sender but commission is charded from check issuer.

##### Example

```go
data := transaction.NewRedeemCheckData().
	MustSetProof("da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600").
	MustSetRawCheck("Mcf89b01830f423f8a4d4e5400000000000000843b9aca00b8419b3beac2c6ad88a8bd54d24912754bb820e58345731cb1b9bc0885ee74f9e50a58a80aa990a29c98b05541b266af99d3825bb1e5ed4e540c6e2f7c9b40af9ecc011ca00f7ba6d0aa47d74274b960fba02be03158d0374b978dcaa5f56fc7cf1754f821a019a829a3b7bba2fc290f5c96e469851a3876376d6a6a4df937327b3a5e9e8297")
```

#### Set candidate online transaction

Transaction for turning candidate on. This transaction should be sent from address which is set in the "Declare candidacy transaction".

PubKey - Public key of a validator.

##### Example

```go
data, _ := transaction.NewSetCandidateOnData().
	SetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43")
```

#### Set candidate offline transaction

Transaction for turning candidate off. This transaction should be sent from address which is set in the "Declare candidacy transaction".

PubKey - Public key of a validator.

##### Example

```go
data, _ := transaction.NewSetCandidateOffData().
	SetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43")
```

#### Create multisig address

Transaction for creating multisignature address.

```go
data := transaction.NewCreateMultisigData().
		MustAddSigData("Mxee81347211c72524338f9680072af90744333143", 1).
		MustAddSigData("Mxee81347211c72524338f9680072af90744333145", 3).
		MustAddSigData("Mxee81347211c72524338f9680072af90744333144", 5).
		SetThreshold(7)
```

Get the multisig address to use it for transaction signatures

```go
msigAddress := dataMultisig.AddressString()
signedTx, _ := tx.Sign(msigAddress, privateKey1, privateKey2, privateKey3)
```
#### Multisend transaction

Transaction for sending coins to multiple addresses.

##### Example

```go
symbolMNT := "MNT"
data := transaction.
    NewMultisendData().
    AddItem(
        *transaction.NewMultisendDataItem().
            SetCoin(symbolMNT).
            SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
            MustSetTo("Mxfe60014a6e9ac91618f5d1cab3fd58cded61ee99"),
    ).AddItem(
        *transaction.NewMultisendDataItem().
            SetCoin(symbolMNT).
            SetValue(big.NewInt(0).Mul(big.NewInt(2), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
            MustSetTo("Mxddab6281766ad86497741ff91b6b48fe85012e3c"),
    )
```

#### Edit candidate transaction

Transaction for editing existing candidate.

##### Example

```go
data := transaction.NewEditCandidateData().
    MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
    MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
    MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47")
```

### Get fee of transaction

```go
signedTransaction, _ := transaction.Sign(privateKey)
fee := signedTransaction.Fee()
```

### Get hash of transaction

```go
hash, _ := signedTransaction.Hash()
```

### Get Public Key of transaction

```go
key, _ := signedTransaction.PublicKey()
```

### Get sender address of transaction

```go
address, _ := signedTransaction.SenderAddress()
```

### Decode Transaction

```go
transactionObject, _ := transaction.Decode("0xf8840102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ca01f36e51600baa1d89d2bee64def9ac5d88c518cdefe45e3de66a3cf9fe410de4a01bc2228dc419a97ded0efe6848de906fbe6c659092167ef0e7dcb8d15024123a")
```

### Minter Deep Links

```go
link, _ := NewDeepLink(
		NewSendData().
			MustSetTo("Mx18467bbb64a8edf890201d526c35957d82be3d95").
			SetCoin("BIP").
			SetValue(big.NewInt(0).Mul(big.NewInt(12345), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(14), nil))),
	)
link.SetPayload([]byte("Hello World"))
encode, _ := link.Encode()
```

More info about [Minter Link Protocol](https://github.com/MinterTeam/minter-link-protocol)

### Minter Check

Minter Check is like an ordinary bank check. Each user of network can issue check with any amount of coins and pass it to another person. Receiver will be able to cash a check from arbitrary account.

* Create Issue Check. Nonce - unique "id" of the check. Coin Symbol - symbol of coin. Value - amount of coins. Due Block - defines last block height in which the check can be used.

```go
check := transaction.NewIssueCheck(
    480,
    TestNetChainID,
    999999,
    "MNT",
    big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
).SetPassphrase("pass")
```

* Sign Issue Check

```go
sign, _ := check.Sign("2919c43d5c712cae66f869a524d9523999998d51157dc40ac4d8d80a7602ce02")
```

* Prepare check string and convert to data

```go
data, _ := transaction.DecodeIssueCheck("Mcf8a38334383002830f423f8a4d4e5400000000000000888ac7230489e80000b841d184caa333fe636288fc68d99dea2c8af5f7db4569a0bb91e03214e7e238f89d2b21f4d2b730ef590fd8de72bd43eb5c6265664df5aa3610ef6c71538d9295ee001ba08bd966fc5a093024a243e62cdc8131969152d21ee9220bc0d95044f54e3dd485a033bc4e03da3ea8a2cd2bd149d16c022ee604298575380db8548b4fd6672a9195")
``` 

* Proof check

```go
check, _ := transaction.NewCheckAddress("Mxa7bc33954f1ce855ed1a8c768fdd32ed927def47", "pass")
proof, _ := check.Proof()
```

### Minter Wallet

```go
import "github.com/MinterTeam/minter-go-sdk/wallet"
```

* Create wallet. This method returns generated seed, private key, public key, mnemonic and Minter address.

```go
walletData, _ := wallet.Create()
```

* Generate mnemonic.

```go
mnemonic, _ := wallet.NewMnemonic()
```

* Get seed from mnemonic.

```go
seed, _ := wallet.Seed(mnemonic)
```

* Get private key from seed.

```go
prKey, _ := wallet.PrivateKeyBySeed(seed)
```

* Get public key from private key.

```go
pubKey, _ := wallet.PublicKeyByPrivateKey(validPrivateKey)
```

* Get Minter address from public key.

```go
address, _ := wallet.AddressByPublicKey(validPublicKey)
```
