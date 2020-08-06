package transaction

import (
	"testing"
)

func TestCreateMultisigData_Sign(t *testing.T) {
	data := NewCreateMultisigData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx772fd5bd06356250e5efe572b6ae615860ee0c17", 3).
		MustAddSigData("Mx9c7f68ff71b5819c41e8f87cc99bdf6359da3d75", 5).
		SetThreshold(7)

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(11)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1).SetSignatureType(SignatureTypeSingle)

	signedTx, err := transaction.Sign("ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8990b0201010cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75808001b845f8431ba0dfc662df298edef48a1a9623a735b55b3acd32023c24a40efc90a85d37209d04a06c2bcd518c2e47ac102941776a68f3ada206260828354e76e2080396bf18f2b1"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestCreateMultisigData_SignGetAddress(t *testing.T) {
	data := NewCreateMultisigData().
		MustAddSigData("Mxee81347211c72524338f9680072af90744333143", 1).
		MustAddSigData("Mxee81347211c72524338f9680072af90744333145", 3).
		MustAddSigData("Mxee81347211c72524338f9680072af90744333144", 5).
		SetThreshold(7)

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1).SetSignatureType(SignatureTypeSingle)

	signedTx, err := transaction.Sign("bc3503cae8c8561df5eadc4a9eda21d32c252a6c94cfae55b5310bf6085c8582")
	if err != nil {
		t.Fatal(err)
	}

	address, err := signedTx.SenderAddress()
	if "Mxb43154a0bc801c4b7361bf1a535b5e08e34e401c" != address {
		t.Errorf("Address got %s, want %s", address, "Mxb43154a0bc801c4b7361bf1a535b5e08e34e401c")
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf899010201010cb848f84607c3010305f83f94ee81347211c72524338f9680072af9074433314394ee81347211c72524338f9680072af9074433314594ee81347211c72524338f9680072af90744333144808001b845f8431ca0224c6166a1f4667cb0bee9ce7ed88879285b8ffc9b4eac3f03faa1797d1f8684a0276dc68fc640924e970c3607af33988a0955e7c2dff78a16ba795da9ddffe988"
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestDecodeCreateMultisig(t *testing.T) {
	decode, err := Decode("0xf899010201010cb848f84607c3010305f83f94ee81347211c72524338f9680072af9074433314394ee81347211c72524338f9680072af9074433314594ee81347211c72524338f9680072af90744333144808001b845f8431ca0224c6166a1f4667cb0bee9ce7ed88879285b8ffc9b4eac3f03faa1797d1f8684a0276dc68fc640924e970c3607af33988a0955e7c2dff78a16ba795da9ddffe988")
	if err != nil {
		t.Fatal(err)
	}
	data := decode.Data().(*CreateMultisigData)
	if data.Threshold != 7 {
		t.Errorf("Threshold got %d, want %d", data.Threshold, 1)
	}
}
