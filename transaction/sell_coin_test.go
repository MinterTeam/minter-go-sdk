package transaction

import (
	"math/big"
	"testing"
)

const txSellCoin = "0xf865010201010295d401880de0b6b3a764000002880de0b6b3a7640000808001b845f8431ca01552ab0503f8173bef46f2336d48ef6e1fae7bb5aa8b51ec7332b720a8a2f15ca0166970c5d209bac8b5ffae32047f1e4e868c5a20f522aeebb0bc523ae16c64fa"

func TestTransactionSellCoin_Sign(t *testing.T) {
	t.Parallel()
	data := NewSellCoinData().
		SetCoinToSell(1).
		SetValueToSell(BipToPip(big.NewInt(1))).
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
	if encode != txSellCoin {
		t.Errorf("EncodeTx got %s, want %s", string(encode), txSellCoin)
	}
}
func TestDecode_sellCoin(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txSellCoin)
	if err != nil {
		t.Fatal(err)
	}

	if decode.Fee().String() != "100000000000000000" {
		t.Error("sell coin transaction fee is invalid", decode.Fee().String())
	}
}
