package transaction

import (
	"testing"
)

func TestCreateMultisigData_Sign(t *testing.T) {
	data := NewCreateMultisigData().
		MustAddSigData("Mx5bee1f257334734d850d300a097e3278d27146c6", 1).
		MustAddSigData("Mx1575064e41cf4408e8c96aef681b770f102024db", 3).
		MustAddSigData("Mxe19a63c522e042fe1919ed34ee7a59fc0f8ef305", 5).
		SetThreshold(7)

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT")

	signedTx, err := transaction.Sign("cbbe8a949c85edec1d4c073ac549f0d256abf6e0c0d3877a997f12240f1d08e7", "670862b4d0490d9b3cb797bcd9f910b8d7e47dfe62d24fee8a9e3587ed9d04c9", "a38ebac3f6130a34a1bd28476f733346b74c1dccf3f29bcc927826f844724960")
	if err != nil {
		t.Fatal(err)
	}

	//validSignature := ""
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encode)
	//if (bytes) != validSignature {
	//	t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	//}
}
