package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewEditMultisigOwnersData() {
	data := transaction.NewEditMultisigOwnersData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx772fd5bd06356250e5efe572b6ae615860ee0c17", 3).
		MustAddSigData("Mx9c7f68ff71b5819c41e8f87cc99bdf6359da3d75", 5).
		SetThreshold(7)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.
		SetNonce(11).
		SetGasPrice(1).
		SetGasCoin(1).
		SetSignatureType(transaction.SignatureTypeMulti).
		Sign(transaction.MultisigAddress("Mx9c7f68ff71b5819c41e8f87cc99bdf6359da3d75", 21), "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")

	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Result: 0xf8b20b02010112b848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75808002b85ef85c94f954480762e2cb5cfed7ab85a75e9692b6a138aaf845f8431ca0ebf73833cbf3ff51a9adb013834af1bbd383e3a90ced0bd41a9988f54fd84071a02dad3b6fe991507e59cf8643e1b243c760a6af8f6e6e9264514d3c2b50cc384e

	// Output:
	// 0xf8b20b02010112b848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75808002b85ef85c94f954480762e2cb5cfed7ab85a75e9692b6a138aaf845f8431ca0ebf73833cbf3ff51a9adb013834af1bbd383e3a90ced0bd41a9988f54fd84071a02dad3b6fe991507e59cf8643e1b243c760a6af8f6e6e9264514d3c2b50cc384e
}
