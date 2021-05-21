package transaction

import (
	"testing"
)

const txSetCandidateOff = "0xf872010201010ba2e1a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ca07c573ba835b7501fae8af3c5d2a414933c9070a217681dc359f8f44a55eeecfaa010d3fd6aa1078d5a2fca66b5f4b456deacfc49eb4be055ba531a4247cb516839"

func TestTransactionSetCandidateOff_Sign(t *testing.T) {
	t.Parallel()
	pubKey := "Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43"
	data := NewSetCandidateOffData().
		MustSetPubKey(pubKey)

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	if data.PubKey.String() != pubKey {
		t.Errorf("PubKey got %s, want %s", data.PubKey.String(), pubKey)
	}

	transaction := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1)

	signedTx, err := transaction.Sign("05ddcd4e6f7d248ed1388f0091fe345bf9bf4fc2390384e26005e7675c98b3c1")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != txSetCandidateOff {
		t.Errorf("EncodeTx got %s, want %s", encode, txSetCandidateOff)
	}
}

func TestDecode_setCandidateOff(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txSetCandidateOff)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypeSetCandidateOffline {
		t.Error("set candidate off transaction type is invalid", decode.GetTransaction().Type)
	}
}
