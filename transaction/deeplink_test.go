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

	data := "-DsBqumKQklQAAAAAAAAAJQYRnu7ZKjt-JAgHVJsNZV9gr49lYgRIhD0do20AItIZWxsbyBXb3JsZICAgA"
	if encode != data {
		t.Errorf("Encode got %s, want %s", encode, data)
	}
}

func TestDeepLink_CreateLinkSend(t *testing.T) {
	link, err := NewDeepLink(
		NewSendData().
			MustSetTo("Mx7633980c000139dd3bd24a3f54e06474fa941e16").
			SetCoin("MNT").
			SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))),
	)
	if err != nil {
		t.Fatal(err)
	}

	link.SetPayload([]byte("custom message")).SetGasCoin("ASD")

	text, err := link.CreateLink("pass")
	if err != nil {
		t.Fatal(err)
	}

	data := "https://bip.to/tx/-EgBqumKTU5UAAAAAAAAAJR2M5gMAAE53TvSSj9U4GR0-pQeFoiKxyMEiegAAI5jdXN0b20gbWVzc2FnZYCAikFTRAAAAAAAAAA?p=cGFzcw"
	if text != data {
		t.Errorf("Encode got %s, want %s", text, data)
	}
}
