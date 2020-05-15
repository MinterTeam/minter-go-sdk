package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleTransaction_Sign() {
	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(
		transaction.NewSendData().
			SetCoin("MNT").
			SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
			MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
	)

	signedTransaction, _ := tx.
		SetGasPrice(1).
		SetGasCoin("MNT").
		SetNonce(1).
		Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")

	senderAddress, _ := signedTransaction.SenderAddress()
	fmt.Println(senderAddress)

	fee := signedTransaction.Fee().String()
	fmt.Println(fee)

	hash, _ := signedTransaction.Hash()
	fmt.Println(hash)

	// Output:
	// Mx622e1e0e788f4b1258f7e2a196f738c6a360c3de
	// 10000000000000000
	// Mt13b73500c171006613fa8e82cc8b29857af1d63a34ca2cada95024bacca1670c

}

func ExampleTransaction_Sign_SignatureTypeMulti_1() {
	createMultisigData := transaction.NewCreateMultisigData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx6bf192730d01a19739b5030cdb6a60c992712a59", 3).
		MustAddSigData("Mx823bb524d5702addbe13086082f7f0310e07d176", 5).
		SetThreshold(7)

	multisigAddress := createMultisigData.AddressString()
	fmt.Println(multisigAddress)
	// Result: Mx0023aa9371e0779189ef5a7434456fc21a938945

	symbolMNT := "MNT"
	data, _ := transaction.NewSendData().
		SetCoin(symbolMNT).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(symbolMNT).SetSignatureType(transaction.SignatureTypeMulti).Sign(
		multisigAddress,
		"ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
		"b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
		"4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba",
	)

	encode, _ := signedTx.Encode()
	fmt.Println(encode)
	// Result: 0xf901270102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ba0014aaffef58c3def74bbb828d7cba907df59b50a68749b8d90aa0d7520571be3a04397def13aa5a38b666d5ecf590af7fdec18663bfa448d517d6671fbe25cdde6f8431ba07bd81f68708141c01ed3bac914cc04dc07831989cb86c4b0e992ad9677bfa33aa03b0d936c268e080bbb85a70cfa6c48a88f023d9e06fa4ecfb9e3cb6659bc767af8431ba0767922509d65315ddf728da8cf5450fa8ba410680f7046405a1eeb7cf22f521aa01222a82c41f7ef51e5b6a64414078185393578f8a5373ac5f5a19ee512b9317b

	// Output:
	// Mx0023aa9371e0779189ef5a7434456fc21a938945
	// 0xf901270102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ba0014aaffef58c3def74bbb828d7cba907df59b50a68749b8d90aa0d7520571be3a04397def13aa5a38b666d5ecf590af7fdec18663bfa448d517d6671fbe25cdde6f8431ba07bd81f68708141c01ed3bac914cc04dc07831989cb86c4b0e992ad9677bfa33aa03b0d936c268e080bbb85a70cfa6c48a88f023d9e06fa4ecfb9e3cb6659bc767af8431ba0767922509d65315ddf728da8cf5450fa8ba410680f7046405a1eeb7cf22f521aa01222a82c41f7ef51e5b6a64414078185393578f8a5373ac5f5a19ee512b9317b
}

func ExampleTransaction_Sign_SignatureTypeMulti_2() {
	symbolMNT := "MNT"
	data, _ := transaction.NewSendData().
		SetCoin(symbolMNT).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	tx.SetNonce(1).SetGasPrice(1).SetGasCoin(symbolMNT).SetSignatureType(transaction.SignatureTypeMulti)

	signedTx1, _ := tx.Sign(
		"Mx0023aa9371e0779189ef5a7434456fc21a938945",
		"ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
	)
	signedTx2, _ := signedTx1.Sign(
		"Mx0023aa9371e0779189ef5a7434456fc21a938945",
		"b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
	)
	signedTx3, _ := signedTx2.Sign(
		"Mx0023aa9371e0779189ef5a7434456fc21a938945",
		"4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba",
	)

	encode, _ := signedTx3.Encode()
	fmt.Println(encode)
	// Output:
	// 0xf901270102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ba0014aaffef58c3def74bbb828d7cba907df59b50a68749b8d90aa0d7520571be3a04397def13aa5a38b666d5ecf590af7fdec18663bfa448d517d6671fbe25cdde6f8431ba07bd81f68708141c01ed3bac914cc04dc07831989cb86c4b0e992ad9677bfa33aa03b0d936c268e080bbb85a70cfa6c48a88f023d9e06fa4ecfb9e3cb6659bc767af8431ba0767922509d65315ddf728da8cf5450fa8ba410680f7046405a1eeb7cf22f521aa01222a82c41f7ef51e5b6a64414078185393578f8a5373ac5f5a19ee512b9317b

}

func ExampleTransaction_Sign_SignatureTypeMulti_3() {
	symbolMNT := "MNT"
	data, _ := transaction.NewSendData().
		SetCoin(symbolMNT).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	tx.SetNonce(1).SetGasPrice(1).SetGasCoin(symbolMNT).SetSignatureType(transaction.SignatureTypeMulti)

	msigAddress := "Mx0023aa9371e0779189ef5a7434456fc21a938945"
	signedTx1, _ := tx.Clone().Sign(msigAddress, "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
	signedTx2, _ := tx.Clone().Sign(msigAddress, "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7")
	signedTx3, _ := tx.Clone().Sign(msigAddress, "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
	simpleSignatureData1, _ := signedTx1.SimpleSignatureData()
	simpleSignatureData2, _ := signedTx2.SimpleSignatureData()
	simpleSignatureData3, _ := signedTx3.SimpleSignatureData()
	signedTransaction, _ := tx.Clone().Sign(msigAddress)
	signedTx123, _ := signedTransaction.AddSignature(simpleSignatureData1, simpleSignatureData2, simpleSignatureData3)

	encode, _ := signedTx123.Encode()
	fmt.Println(encode)
	// Output:
	// 0xf901270102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ba0014aaffef58c3def74bbb828d7cba907df59b50a68749b8d90aa0d7520571be3a04397def13aa5a38b666d5ecf590af7fdec18663bfa448d517d6671fbe25cdde6f8431ba07bd81f68708141c01ed3bac914cc04dc07831989cb86c4b0e992ad9677bfa33aa03b0d936c268e080bbb85a70cfa6c48a88f023d9e06fa4ecfb9e3cb6659bc767af8431ba0767922509d65315ddf728da8cf5450fa8ba410680f7046405a1eeb7cf22f521aa01222a82c41f7ef51e5b6a64414078185393578f8a5373ac5f5a19ee512b9317b

}
