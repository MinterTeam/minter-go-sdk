package transaction

import (
	"testing"
)

const txSetCandidateOn = "0xf872010201010aa2e1a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ba0efff777e61a78141ceeab311776dfd0bfc6745f125c688db86ccfa350d3d3b84a074419c32dd0d1d2ebdc1c5bfdffb238d2ef88a618e28a2ce2410880264d3b3cc"

func TestTransactionSetCandidateOn_Sign(t *testing.T) {
	t.Parallel()
	data := NewSetCandidateOnData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
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
	if encode != txSetCandidateOn {
		t.Errorf("EncodeTx got %s, want %s", string(encode), txSetCandidateOn)
	}
}

func TestDecode_setCandidateOn(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txSetCandidateOn)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypeSetCandidateOnline {
		t.Error("set candidate on transaction type is invalid", decode.GetTransaction().Type)
	}
}
