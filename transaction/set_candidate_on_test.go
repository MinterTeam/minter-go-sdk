package transaction

import (
	"encoding/hex"
	"testing"
)

func TestTransactionSetCandidateOn_Sign(t *testing.T) {
	data := NewSetCandidateOnData().
		SetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT")

	privateKey, err := hex.DecodeString("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	key, err := ToECDSA(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := transaction.Sign(key)
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf87c0102018a4d4e54000000000000000aa2e1a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ba0095aed433171fe5ac385ccd299507bdcad3dd2269794fd0d14d4f58327ddc87ea046ec7e4f8f9b477a1255485f36e0567e62283723ecc5a0bd1e5d201e53e85245"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
