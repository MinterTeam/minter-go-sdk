package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleObject_Sign_simple() {
	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(
		transaction.NewSendData().
			SetCoin(1).
			SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
			MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
	)

	signedTransaction, _ := tx.
		SetGasPrice(1).
		SetGasCoin(1).
		SetNonce(1).
		Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")

	senderAddress, _ := signedTransaction.SenderAddress()
	fmt.Println(senderAddress)

	fee := signedTransaction.Fee().String()
	fmt.Println(fee)

	hash, _ := signedTransaction.Hash()
	fmt.Println(hash)

	encode, _ := signedTransaction.Encode()
	fmt.Println(encode)

	// Output:
	// Mx622e1e0e788f4b1258f7e2a196f738c6a360c3de
	// 10000000000000000
	// Mtec2166cced36276426360a79934fbf49f29f9e48e9d1f06ef4afc4f557aa3767
	// 0xf8700102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ba0fffc3f503ace8a5d0c87efe50cf33ad41e3475459120d9c6fd75bd796b192313a0243d643a799e844ad82382d41cee98137a1d0c5888ff13951919e5e241ab89e0

}

func ExampleBuilder_NewTransaction_signMultiSignature1() {
	coinID := transaction.CoinID(1)
	data, _ := transaction.NewSendData().
		SetCoin(coinID).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(coinID).SetSignatureType(transaction.SignatureTypeMulti).Sign(
		"Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2",
		"ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
		"b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
		"4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba",
	)

	encode, _ := signedTx.Encode()
	fmt.Println(encode)
	// Result: 0xf901130102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca07dd407fa5d2a161581d03cdeb7c94fcd5ade47d376af75f2c92d1483f821fe2ca00d16b6cdbceaadcd0fd72bd39ee17841871da333a571535fccfbcf6285881c2af8431ba07c2d063126024a1e19363e7e254312ca9ab37795b06102da25bd1c0dec81a934a043b7bec83db41c594ac7a8d416fca2f83f0e65ada1221fe659ba4dbe1f3c921af8431ba09318e56a242c39c10ce87ab51d10322cc62cf23885867bc89a24e8c3fa8483e9a04c82c1224d1b4efa7fba06623da2896745ce444d35ed77973759e6404b66bb95

	// Output:
	// 0xf901130102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca07dd407fa5d2a161581d03cdeb7c94fcd5ade47d376af75f2c92d1483f821fe2ca00d16b6cdbceaadcd0fd72bd39ee17841871da333a571535fccfbcf6285881c2af8431ba07c2d063126024a1e19363e7e254312ca9ab37795b06102da25bd1c0dec81a934a043b7bec83db41c594ac7a8d416fca2f83f0e65ada1221fe659ba4dbe1f3c921af8431ba09318e56a242c39c10ce87ab51d10322cc62cf23885867bc89a24e8c3fa8483e9a04c82c1224d1b4efa7fba06623da2896745ce444d35ed77973759e6404b66bb95
}

func ExampleBuilder_NewTransaction_signMultiSignature2() {
	coinID := transaction.CoinID(1)
	data, _ := transaction.NewSendData().
		SetCoin(coinID).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	tx.SetNonce(1).SetGasPrice(1).SetGasCoin(coinID).SetSignatureType(transaction.SignatureTypeMulti)

	msigAddress := "Mx0023aa9371e0779189ef5a7434456fc21a938945"
	signedTx1, _ := tx.Sign(
		msigAddress,
		"ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
	)
	signedTx2, _ := signedTx1.Sign(
		msigAddress,
		"b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
	)
	signedTx3, _ := signedTx2.Sign(
		msigAddress,
		"4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba",
	)

	encode, _ := signedTx3.Encode()
	fmt.Println(encode)
	// Output:
	// 0xf901130102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ca07dd407fa5d2a161581d03cdeb7c94fcd5ade47d376af75f2c92d1483f821fe2ca00d16b6cdbceaadcd0fd72bd39ee17841871da333a571535fccfbcf6285881c2af8431ba07c2d063126024a1e19363e7e254312ca9ab37795b06102da25bd1c0dec81a934a043b7bec83db41c594ac7a8d416fca2f83f0e65ada1221fe659ba4dbe1f3c921af8431ba09318e56a242c39c10ce87ab51d10322cc62cf23885867bc89a24e8c3fa8483e9a04c82c1224d1b4efa7fba06623da2896745ce444d35ed77973759e6404b66bb95

}

func ExampleSignatureMulti_Single() {
	coinID := transaction.CoinID(1)
	data, _ := transaction.NewSendData().
		SetCoin(coinID).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	tx.SetNonce(1).SetGasPrice(1).SetGasCoin(coinID).SetSignatureType(transaction.SignatureTypeMulti)

	msigAddress := "Mx0023aa9371e0779189ef5a7434456fc21a938945"
	signedTx1, _ := tx.Clone().Sign(msigAddress, "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
	signedTx2, _ := tx.Clone().Sign(msigAddress, "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7")
	signedTx3, _ := tx.Clone().Sign(msigAddress, "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
	simpleSignatureData1, _ := signedTx1.SingleSignatureData()
	simpleSignatureData2, _ := signedTx2.SingleSignatureData()
	simpleSignatureData3, _ := signedTx3.SingleSignatureData()
	signedTransaction, _ := tx.Clone().Sign(msigAddress)
	signedTx123, _ := signedTransaction.AddSignature(simpleSignatureData1, simpleSignatureData2, simpleSignatureData3)

	encode, _ := signedTx123.Encode()
	fmt.Println(encode)
	// Output:
	// 0xf901130102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ca07dd407fa5d2a161581d03cdeb7c94fcd5ade47d376af75f2c92d1483f821fe2ca00d16b6cdbceaadcd0fd72bd39ee17841871da333a571535fccfbcf6285881c2af8431ba07c2d063126024a1e19363e7e254312ca9ab37795b06102da25bd1c0dec81a934a043b7bec83db41c594ac7a8d416fca2f83f0e65ada1221fe659ba4dbe1f3c921af8431ba09318e56a242c39c10ce87ab51d10322cc62cf23885867bc89a24e8c3fa8483e9a04c82c1224d1b4efa7fba06623da2896745ce444d35ed77973759e6404b66bb95

}

func ExampleObject_Signers_multi() {
	decode, _ := transaction.Decode("0xf901130102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ca07dd407fa5d2a161581d03cdeb7c94fcd5ade47d376af75f2c92d1483f821fe2ca00d16b6cdbceaadcd0fd72bd39ee17841871da333a571535fccfbcf6285881c2af8431ba07c2d063126024a1e19363e7e254312ca9ab37795b06102da25bd1c0dec81a934a043b7bec83db41c594ac7a8d416fca2f83f0e65ada1221fe659ba4dbe1f3c921af8431ba09318e56a242c39c10ce87ab51d10322cc62cf23885867bc89a24e8c3fa8483e9a04c82c1224d1b4efa7fba06623da2896745ce444d35ed77973759e6404b66bb95")
	signers, _ := decode.Signers()
	for _, signer := range signers {
		fmt.Println(signer)
	}

	// Output:
	// Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c
	// Mx6bf192730d01a19739b5030cdb6a60c992712a59
	// Mx823bb524d5702addbe13086082f7f0310e07d176
}

func ExampleObject_Signers_single() {
	decode, _ := transaction.Decode("0xf8700102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ba0fffc3f503ace8a5d0c87efe50cf33ad41e3475459120d9c6fd75bd796b192313a0243d643a799e844ad82382d41cee98137a1d0c5888ff13951919e5e241ab89e0")
	signers, _ := decode.Signers()
	for _, signer := range signers {
		fmt.Println(signer)
	}

	address, _ := decode.SenderAddress()
	fmt.Println(address)

	// Output:
	// Mx622e1e0e788f4b1258f7e2a196f738c6a360c3de
	// Mx622e1e0e788f4b1258f7e2a196f738c6a360c3de
}
