package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewDeepLink() {
	link, _ := transaction.NewDeepLink(
		transaction.NewSendData().
			MustSetTo("Mx18467bbb64a8edf890201d526c35957d82be3d95").
			SetCoin("BIP").
			SetValue(big.NewInt(0).Mul(big.NewInt(123456789), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-8), nil))),
	)

	link.SetPayload([]byte("Hello World"))
	encode, _ := link.Encode()
	fmt.Println(encode)
	// Output:
	// -DsBqumKQklQAAAAAAAAAJQYRnu7ZKjt-JAgHVJsNZV9gr49lYgRIhD0do20AItIZWxsbyBXb3JsZICAgA
}

func ExampleNewDeepLink_createLink() {
	link, _ := transaction.NewDeepLink(
		transaction.NewSendData().
			MustSetTo("Mx7633980c000139dd3bd24a3f54e06474fa941e16").
			SetCoin("MNT").
			SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))),
	)

	link.SetPayload([]byte("custom message")).SetGasCoin("ASD")

	data, _ := link.CreateLink("pass")
	fmt.Println(data)
	// Output:
	// https://bip.to/tx/-EgBqumKTU5UAAAAAAAAAJR2M5gMAAE53TvSSj9U4GR0-pQeFoiKxyMEiegAAI5jdXN0b20gbWVzc2FnZYCAikFTRAAAAAAAAAA?p=cGFzcw
}
