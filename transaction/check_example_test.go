package transaction_test

import (
	"encoding/hex"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewCheck() {
	check := transaction.NewCheck(
		480,
		transaction.TestNetChainID,
		999999,
		"MNT",
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		"MNT",
	).SetPassphrase("pass")

	sign, _ := check.Sign("64e27afaab363f21eec05291084367f6f1297a7b280d69d672febecda94a09ea")
	encode, _ := sign.Encode()
	fmt.Println(encode)
	// Output:
	// Mcf8ae8334383002830f423f8a4d4e5400000000000000888ac7230489e800008a4d4e5400000000000000b841497c5f3e6fc182fd1a791522a9ef7576710bdfbc86fdbf165476ef220e89f9ff1380f93f2d9a2f92fdab0edc1e2605cc2c69b707cd404b2cb1522b7aba4defd5001ba083c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594a047b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df
}

func ExampleCheckAddress_Proof() {
	check, _ := transaction.NewCheckAddress("Mxa7bc33954f1ce855ed1a8c768fdd32ed927def47", "pass")
	proof, _ := check.Proof()
	fmt.Println(proof)
	// Output:
	// da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600
}

func ExampleDecodeCheck() {
	data, _ := transaction.DecodeCheck("+K6DNDgwAoMPQj+KTU5UAAAAAAAAAIiKxyMEiegAAIpNTlQAAAAAAAAAuEFJfF8+b8GC/Rp5FSKp73V2cQvfvIb9vxZUdu8iDon5/xOA+T8tmi+S/asO3B4mBcwsabcHzUBLLLFSK3q6Te/VABugg8mUUWnwp7vllpc7MtyIdgh4BYCx07x7GIvts704VZSgR7LVNFlG7VSY9b7nE/hidqrARqX++CC+rud6m2+bwd8=")

	fmt.Println(string(data.Nonce))
	// Result: 480

	fmt.Println(data.ChainID)
	// Result: 2

	fmt.Println(data.Coin.String())
	// Result: MNT

	fmt.Println(data.DueBlock)
	// Result: 999999

	fmt.Println(data.Value.String())
	// Result: 10000000000000000000

	fmt.Println(data.Lock.String())
	//Result: 985283871505876742053353384809055983203325304659217336382177168476233943196356552072809135181493031263037291805582875144952622907359402361966594748392133888

	fmt.Println(data.V.Int64())
	// Result: 27

	fmt.Println(hex.EncodeToString(data.R.Bytes()))
	// Result: 83c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594

	fmt.Println(hex.EncodeToString(data.S.Bytes()))
	// Result: 47b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df

	sender, _ := data.Sender()
	fmt.Println(sender)
	// Result: Mxce931863b9c94a526d94acd8090c1c5955a6eb4b

	// Output:
	// 480
	// 2
	// MNT
	// 999999
	// 10000000000000000000
	// 985283871505876742053353384809055983203325304659217336382177168476233943196356552072809135181493031263037291805582875144952622907359402361966594748392133888
	// 27
	// 83c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594
	// 47b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df
	// Mxce931863b9c94a526d94acd8090c1c5955a6eb4b
}
