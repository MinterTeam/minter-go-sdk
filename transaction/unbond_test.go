package transaction

import (
	"math/big"
	"testing"
)

func TestTransactionUnbond_Sign(t *testing.T) {
	data := NewUnbondData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCoin("MNT").
		SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT")

	signedTx, err := transaction.Sign("6e1df6ec69638d152f563c5eca6c13cdb5db4055861efc11ec1cdd578afd96bf")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf88f0102018a4d4e540000000000000008b6f5a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a438a4d4e5400000000000000888ac7230489e80000808001b844f8421ca0ff5766c85847b37a276f3f9d027fb7c99745920fa395c7bd399cedd8265c5e1d9f791bcdfe4d1bc1e73ada7bf833103c828f22d83189dad2b22ad28a54aacf2a"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
