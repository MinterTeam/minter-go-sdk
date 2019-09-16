// +build integration

package api

import (
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"github.com/foxnut/go-hdwallet"
	"math/big"
	"strings"
	"testing"
)

func TestApi_Send(t *testing.T) {
	data := transaction.NewSendData().
		SetCoin("MNT").
		SetValue(big.NewInt(1)).
		MustSetTo("Mxee81347211c72524338f9680072af90744333146")

	newTransaction, err := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	mnemonic := "perfect bid satoshi giant cigar judge tonight possible harbor render else food"
	master, err := hdwallet.NewKey(hdwallet.Mnemonic(mnemonic))
	if err != nil {
		t.Fatal(err)
	}

	wallet, err := master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
	if err != nil {
		t.Fatal(err)
	}
	address, err := wallet.GetAddress()
	if err != nil {
		t.Fatal(err)
	}

	addr := strings.Replace(strings.ToLower(address), "0x", "Mx", 1)
	wantAddress := "Mxeeee1973381ab793719fff497b9a516719fcd5a2"
	if wantAddress != addr {
		t.Fatalf("wallet address get %s, want %s", addr, wantAddress)
	}

	wantedPrKey := "ecc067573863f893f2195d550ff0d703d31e5a1255791e410ca7ff7cf5f0a7aa"
	if wallet.GetKey().PrivateHex() != wantedPrKey {
		t.Fatalf("privateKey get %s, want %s", wallet.GetKey().PrivateHex(), wantedPrKey)
	}

	api := NewApi("https://minter-node-1.testnet.minter.network:8841")
	nonce, err := api.Nonce([]byte("Mxeeee1973381ab793719fff497b9a516719fcd5a2"))
	if err != nil {
		t.Fatal(err)
	}

	key, err := transaction.ToECDSA(wallet.GetKey().Private.Serialize())
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := newTransaction.
		SetNonce(nonce).
		SetGasCoin("MNT").
		SetGasPrice(1).
		Sign(key)
	if err != nil {
		t.Fatal(err)
	}

	res, err := api.Send(signedTransaction)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", res)
}
