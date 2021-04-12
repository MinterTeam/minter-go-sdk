package transaction

import (
	"math/big"
	"testing"
)

const txDelegate = "0xf87c0102010107aceba00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a4301888ac7230489e80000808001b845f8431ca06de9cf33b536456b2197b121031c215cb72f9498a187df85a988d737464e5dc3a059789b60988e64f6cfe65d72b942293413387e2956a852d89562546bd425b694"

func TestTransactionDelegate_Sign(t *testing.T) {
	t.Parallel()
	data := NewDelegateData().
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
	if encode != txDelegate {
		t.Errorf("EncodeTx got %s, want %s", encode, txDelegate)
	}
}

func TestDecode_delegate(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txDelegate)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypeDelegate {
		t.Error("delegate transaction type is invalid", decode.GetTransaction().Type)
	}
}
