package transaction

import (
	"math/big"
	"testing"
)

func TestTransactionDeclareCandidacy_Sign(t *testing.T) {
	data, err := NewDeclareCandidacyData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCommission(10).
		SetCoin(1).
		SetStake(big.NewInt(0).Mul(big.NewInt(5), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetAddress("Mx9f7fd953c2c69044b901426831ed03ee0bd0597a")
	if err != nil {
		t.Fatal(err)
	}

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)

	signedTx, err := transaction.Sign("6e1df6ec69638d152f563c5eca6c13cdb5db4055861efc11ec1cdd578afd96bf")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8940102010106b843f841949f7fd953c2c69044b901426831ed03ee0bd0597aa00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a430a01884563918244f40000808001b845f8431ca0d22ec4cdeb146488e39eadcefa982d4bd01a8d44a9af353030db9d1f732a04e5a031a5354af73a47834e3c3b8b9b01b1cd22991b212e598179070a585ea1fe2af2"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", bytes, validSignature)
	}
}
