package transaction

import (
	"math/big"
	"testing"
)

const txDeclareCandidacy = "0xf8940102010106b843f841949f7fd953c2c69044b901426831ed03ee0bd0597aa00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a430a01884563918244f40000808001b845f8431ca0d22ec4cdeb146488e39eadcefa982d4bd01a8d44a9af353030db9d1f732a04e5a031a5354af73a47834e3c3b8b9b01b1cd22991b212e598179070a585ea1fe2af2"

func TestTransactionDeclareCandidacy_Sign(t *testing.T) {
	t.Parallel()
	data := NewDeclareCandidacyData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCommission(10).
		SetCoin(1).
		SetStake(BipToPip(big.NewInt(5))).
		MustSetAddress("Mx9f7fd953c2c69044b901426831ed03ee0bd0597a")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	transaction := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1)

	signedTx, err := transaction.Sign("6e1df6ec69638d152f563c5eca6c13cdb5db4055861efc11ec1cdd578afd96bf")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != txDeclareCandidacy {
		t.Errorf("EncodeTx got %s, want %s", encode, txDeclareCandidacy)
	}
}

func TestDecode_declareCandidacy(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txDeclareCandidacy)
	if err != nil {
		t.Fatal(err)
	}

	if decode.Fee().String() != "10000000000000000000" {
		t.Error("declare candidacy transaction fee is invalid", decode.Fee().String())
	}
}
