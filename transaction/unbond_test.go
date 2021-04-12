package transaction

import (
	"math/big"
	"testing"
)

const txUnbond = "0xf87c0102010108aceba00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a4301888ac7230489e80000808001b845f8431ca0f5b9273c522c6b948523ae922594389619fd5c21846361bec6c72ee2c45b9a21a00dbeed5293f74a0a7924f2a3459f57270358d8621e092f66da38d0dbab9055e1"

func TestTransactionUnbond_Sign(t *testing.T) {
	t.Parallel()
	data := NewUnbondData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCoin(1).
		SetValue(BipToPip(big.NewInt(10)))

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
	if encode != txUnbond {
		t.Errorf("EncodeTx got %s, want %s", encode, txUnbond)
	}
}

func TestDecode_unbond(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txUnbond)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypeUnbond {
		t.Error("unbond transaction type is invalid", decode.GetTransaction().Type)
	}
}
