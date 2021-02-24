package transaction

import (
	"math/big"
	"testing"
)

const txSellAllCoin = "0xf85c01020101038ccb0102880de0b6b3a7640000808001b845f8431ba0db51e3ca2b75a4a617362946f5f5ee26c75900dae8f8be2400509338bafcd8d4a02e031ead1656d1321520564a34321acc8c6dda63cd4c8bd4530ec641aa4b1c7b"

func TestTransactionSellAllCoin_Sign(t *testing.T) {
	t.Parallel()
	data := NewSellAllCoinData().
		SetCoinToSell(1).
		SetCoinToBuy(2).
		SetMinimumValueToBuy(BipToPip(big.NewInt(1)))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	transaction := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1)

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != txSellAllCoin {
		t.Errorf("EncodeTx got %s, want %s", string(encode), txSellAllCoin)
	}
}

func TestDecode_sellAllCoin(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txSellAllCoin)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypeSellAllCoin {
		t.Error("sell all coin transaction type is invalid", decode.GetTransaction().Type)
	}
}
