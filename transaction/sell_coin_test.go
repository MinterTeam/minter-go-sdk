package transaction

import (
	"math/big"
	"testing"
)

func TestTransactionSellCoin_Sign(t *testing.T) {
	data := NewSellCoinData().
		SetCoinToSell(1).
		SetValueToSell(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetCoinToBuy(2).
		SetMinimumValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf865010201010295d401880de0b6b3a764000002880de0b6b3a7640000808001b845f8431ca01552ab0503f8173bef46f2336d48ef6e1fae7bb5aa8b51ec7332b720a8a2f15ca0166970c5d209bac8b5ffae32047f1e4e868c5a20f522aeebb0bc523ae16c64fa"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bytes != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
