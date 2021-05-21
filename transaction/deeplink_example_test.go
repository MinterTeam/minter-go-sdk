package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewDeepLink() {
	link, _ := transaction.NewDeepLink(
		transaction.NewSendData().
			MustSetTo("Mx18467bbb64a8edf890201d526c35957d82be3d95").
			SetCoin(1).
			SetValue(big.NewInt(0).Mul(big.NewInt(123456789), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-8), nil))),
	)
	link.SetPayload([]byte("Hello World"))
	encode, _ := link.Encode()
	fmt.Println(encode)
	// Output:
	// 8QGg3wGUGEZ7u2So7fiQIB1SbDWVfYK-PZWIESIQ9HaNtACLSGVsbG8gV29ybGTAwMA
}

func ExampleDeepLink_CreateLink() {
	link, _ := transaction.NewDeepLink(
		transaction.NewSendData().
			MustSetTo("Mx7633980c000139dd3bd24a3f54e06474fa941e16").
			SetCoin(1).
			SetValue(transaction.BipToPip(big.NewInt(10))),
	)

	link.SetPayload([]byte("custom message")).SetGasCoin(3)

	data, _ := link.CreateLink("pass")
	fmt.Println(data)
	// Output:
	// https://bip.to/tx/9AGg3wGUdjOYDAABOd070ko_VOBkdPqUHhaIiscjBInoAACOY3VzdG9tIG1lc3NhZ2XAwAM?p=cGFzcw
}

func ExampleDeepLink_CreateLink_customHost() {
	link, _ := transaction.NewDeepLink(
		transaction.NewSendData().
			MustSetTo("Mx7633980c000139dd3bd24a3f54e06474fa941e16").
			SetCoin(1).
			SetValue(transaction.BipToPip(big.NewInt(10))),
	)

	link.MustSetUrl("https://testnet.bip.to").SetPayload([]byte("custom message")).SetGasCoin(3)

	data, _ := link.CreateLink("pass")
	fmt.Println(data)
	// Output:
	// https://testnet.bip.to/tx/9AGg3wGUdjOYDAABOd070ko_VOBkdPqUHhaIiscjBInoAACOY3VzdG9tIG1lc3NhZ2XAwAM?p=cGFzcw
}
