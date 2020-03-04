package transaction

import (
	"fmt"
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

	msigAddress := data.MultisigAddressString()
	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT").SetSignatureType(SignatureTypeSingle).
		SetPayload([]byte(fmt.Sprintf("%v, %v, %v, %v", "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63", "9d78895fa954b2b07fb3f29d2ae9f5eb0dc0e925a68ef8362e40c47ba4adb30c", "7e4089c7b683f1b8d1832a8e977cf79aa459bf170ff196354112747124bbd072", msigAddress)))

	signedTx, err := transaction.Sign("ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901940b02018a4d4e54000000000000000cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75b8f0616530383962333265346530393736636136383838636231303233313438626431613966316363323863356434343265353265353836373534666634386436332c20396437383839356661393534623262303766623366323964326165396635656230646330653932356136386566383336326534306334376261346164623330632c20376534303839633762363833663162386431383332613865393737636637396161343539626631373066663139363335343131323734373132346262643037322c204d78643433656566376239343036373632616130333162383265643062313038323236346131333933348001b845f8431ba03032390eb7457b987b223128b956e514c9847fda62c01b62672712719cca4edfa013d566aef3ba4729a2075ef0db8217e7ae390699be6bc959541d8d4b372219cf"
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

	msigAddress := data.MultisigAddressString()
	addr := "Mx4fe800483f59a36eec2b6f218778f9c5fceb38c0"
	if addr != msigAddress {
		t.Errorf("Address got %s, want %s", msigAddress, addr)
	}

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT").SetSignatureType(SignatureTypeSingle)

	signedTx, err := transaction.Sign("bc3503cae8c8561df5eadc4a9eda21d32c252a6c94cfae55b5310bf6085c8582")
	if err != nil {
		t.Fatal(err)
	}

	address, err := signedTx.SenderAddress()
	if "Mx3e4d56e776ff42c023b1ec99a7486b592a654981" != address {
		t.Errorf("Address got %s, want %s", address, "Mx3e4d56e776ff42c023b1ec99a7486b592a654981")
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8a30102018a4d4e54000000000000000cb848f84607c3010305f83f94ee81347211c72524338f9680072af9074433314394ee81347211c72524338f9680072af9074433314594ee81347211c72524338f9680072af90744333144808001b845f8431ca094eb41d39e6782f5539615cc66da7073d4283893f0b3ee2b2f36aee1eaeb7c57a037f90ffdb45eb9b6f4cf301b48e73a6a81df8182e605b656a52057537d264ab4"
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestDecodeCreateMultisig(t *testing.T) {
	decode, err := Decode("0xf901940b02018a4d4e54000000000000000cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75b8f0616530383962333265346530393736636136383838636231303233313438626431613966316363323863356434343265353265353836373534666634386436332c20396437383839356661393534623262303766623366323964326165396635656230646330653932356136386566383336326534306334376261346164623330632c20376534303839633762363833663162386431383332613865393737636637396161343539626631373066663139363335343131323734373132346262643037322c204d78643433656566376239343036373632616130333162383265643062313038323236346131333933348001b845f8431ba03032390eb7457b987b223128b956e514c9847fda62c01b62672712719cca4edfa013d566aef3ba4729a2075ef0db8217e7ae390699be6bc959541d8d4b372219cf")
	if err != nil {
		t.Fatal(err)
	}
	address := decode.Data().(*CreateMultisigData).MultisigAddressString()
	validAddress := "Mxd43eef7b9406762aa031b82ed0b1082264a13934"
	if address != validAddress {
		t.Errorf("Address got %s, want %s", address, validAddress)
	}
}
