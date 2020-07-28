package transaction

import (
	"testing"
)

func TestTransactionSetCandidateOn_Sign(t *testing.T) {
	data := NewSetCandidateOnData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)

	signedTx, err := transaction.Sign("05ddcd4e6f7d248ed1388f0091fe345bf9bf4fc2390384e26005e7675c98b3c1")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf872010201010aa2e1a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ba0efff777e61a78141ceeab311776dfd0bfc6745f125c688db86ccfa350d3d3b84a074419c32dd0d1d2ebdc1c5bfdffb238d2ef88a618e28a2ce2410880264d3b3cc"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
