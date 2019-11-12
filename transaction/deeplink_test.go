package transaction

import (
	"math/big"
	"testing"
)

func TestDeepLink_NewDeepLink(t *testing.T) {
	link, err := NewDeepLink(
		NewSendData().
			MustSetTo("Mx18467bbb64a8edf890201d526c35957d82be3d95").
			SetCoin("BIP").
			SetValue(big.NewInt(0).Mul(big.NewInt(12345), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(14), nil))),
	)
	if err != nil {
		t.Fatal(err)
	}

	link.SetPayload([]byte("Hello World"))

	encode, err := link.Encode()
	if err != nil {
		t.Fatal(err)
	}

	data := "f83b01aae98a424950000000000000009418467bbb64a8edf890201d526c35957d82be3d9588112210f4768db4008b48656c6c6f20576f726c64808080"
	if encode == data {
		t.Errorf("Encode got %s, want %s", encode, data)
	}
}
