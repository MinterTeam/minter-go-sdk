package transaction

import (
	"math/big"
	"testing"
)

func TestTransaction_Hash(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewSendData().
		SetCoin("MNT").
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
	)
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.
		SetGasPrice(1).
		SetGasCoin("MNT").
		SetNonce(1).
		Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	hash, err := signedTransaction.Hash()
	if err != nil {
		t.Fatal(err)
	}

	validHash := "Mt13b73500c171006613fa8e82cc8b29857af1d63a"
	if hash != validHash {
		t.Errorf("Hash got %s, want %s", hash, validHash)
	}
}
