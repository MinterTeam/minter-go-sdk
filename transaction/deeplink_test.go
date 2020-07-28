package transaction

import (
	"math/big"
	"testing"
)

func TestDeepLink_NewDeepLink(t *testing.T) {
	link, err := NewDeepLink(
		NewSendData().
			MustSetTo("Mx18467bbb64a8edf890201d526c35957d82be3d95").
			SetCoin(2).
			SetValue(big.NewInt(0).Mul(big.NewInt(123456789), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-8), nil))),
	)
	if err != nil {
		t.Fatal(err)
	}

	link.SetPayload([]byte("Hello World"))

	encode, err := link.Encode()
	if err != nil {
		t.Fatal(err)
	}

	data := "8QGg3wKUGEZ7u2So7fiQIB1SbDWVfYK-PZWIESIQ9HaNtACLSGVsbG8gV29ybGSAgIA"
	if encode != data {
		t.Errorf("Encode got %s, want %s", encode, data)
	}
}

func TestDeepLink_CreateLinkSend(t *testing.T) {
	link, err := NewDeepLink(
		NewSendData().
			MustSetTo("Mx7633980c000139dd3bd24a3f54e06474fa941e16").
			SetCoin(1).
			SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))),
	)
	if err != nil {
		t.Fatal(err)
	}

	link.SetPayload([]byte("custom message")).SetGasCoin(3)

	text, err := link.CreateLink("pass")
	if err != nil {
		t.Fatal(err)
	}

	data := "https://bip.to/tx/9AGg3wGUdjOYDAABOd070ko_VOBkdPqUHhaIiscjBInoAACOY3VzdG9tIG1lc3NhZ2WAgAM?p=cGFzcw"
	if text != data {
		t.Errorf("Encode got %s, want %s", text, data)
	}
}
