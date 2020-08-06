package transaction_test

import (
	"encoding/hex"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewCheck() {
	check := transaction.NewCheck(
		"480",
		transaction.TestNetChainID,
		999999,
		1,
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		1,
	).SetPassphrase("pass")

	sign, _ := check.Sign("64e27afaab363f21eec05291084367f6f1297a7b280d69d672febecda94a09ea")
	encode, _ := sign.Encode()
	fmt.Println(encode)
	// Output:
	// Mcf89a8334383002830f423f01888ac7230489e8000001b841ea3d022c8326965556f1b651b14d3124947b8683f7b3ab56fca06e0b4204757b2a11dace85d0139ce4e8fdb18369d07905e733683b8229f41bc216c784b4d714011ca017bffff4b3f431dc938239cd2727f0c1dfa61ccdc98727fa8e9baf608b3755f5a05b768c53d09c5e9517487820df439f496e16e459862e7d449360ce69a2ccc4d6
}

func ExampleCheckAddress_Proof() {
	check, _ := transaction.NewCheckAddress("Mxa7bc33954f1ce855ed1a8c768fdd32ed927def47", "pass")
	proof, _ := check.Proof()
	fmt.Println(proof)
	// Output:
	// da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600
}

func ExampleDecodeCheck() {
	data, err := transaction.DecodeCheck("Mcf89a8334383002830f423f01888ac7230489e8000001b841ea3d022c8326965556f1b651b14d3124947b8683f7b3ab56fca06e0b4204757b2a11dace85d0139ce4e8fdb18369d07905e733683b8229f41bc216c784b4d714011ca017bffff4b3f431dc938239cd2727f0c1dfa61ccdc98727fa8e9baf608b3755f5a05b768c53d09c5e9517487820df439f496e16e459862e7d449360ce69a2ccc4d6")
	if err != nil {
		return
	}

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
	// 1
	// 999999
	// 10000000000000000000
	// 3140622329586495619178957840119431069413669815577585079847850638320450546439777094725740545888163635015645563228183373665216451470665191568503369618620355585
	// 28
	// 17bffff4b3f431dc938239cd2727f0c1dfa61ccdc98727fa8e9baf608b3755f5
	// 5b768c53d09c5e9517487820df439f496e16e459862e7d449360ce69a2ccc4d6
	// Mxce931863b9c94a526d94acd8090c1c5955a6eb4b
}

func ExampleDecodeRawCheck() {
	data, err := transaction.DecodeCheckBase64("+JqDNDgwAoMPQj8BiIrHIwSJ6AAAAbhB6j0CLIMmllVW8bZRsU0xJJR7hoP3s6tW/KBuC0IEdXsqEdrOhdATnOTo/bGDadB5BeczaDuCKfQbwhbHhLTXFAEcoBe///Sz9DHck4I5zScn8MHfphzNyYcn+o6br2CLN1X1oFt2jFPQnF6VF0h4IN9Dn0luFuRZhi59RJNgzmmizMTW")
	if err != nil {
		return
	}

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
	// 1
	// 999999
	// 10000000000000000000
	// 3140622329586495619178957840119431069413669815577585079847850638320450546439777094725740545888163635015645563228183373665216451470665191568503369618620355585
	// 28
	// 17bffff4b3f431dc938239cd2727f0c1dfa61ccdc98727fa8e9baf608b3755f5
	// 5b768c53d09c5e9517487820df439f496e16e459862e7d449360ce69a2ccc4d6
	// Mxce931863b9c94a526d94acd8090c1c5955a6eb4b
}
