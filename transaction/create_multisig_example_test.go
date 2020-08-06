package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewCreateMultisigData() {
	data := transaction.NewCreateMultisigData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx772fd5bd06356250e5efe572b6ae615860ee0c17", 3).
		MustAddSigData("Mx9c7f68ff71b5819c41e8f87cc99bdf6359da3d75", 5).
		SetThreshold(7)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.
		SetNonce(11).
		SetGasPrice(1).
		SetGasCoin(1).
		SetSignatureType(transaction.SignatureTypeSingle).
		Sign("ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")

	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Result: 0xf8990b0201010cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75808001b845f8431ba0dfc662df298edef48a1a9623a735b55b3acd32023c24a40efc90a85d37209d04a06c2bcd518c2e47ac102941776a68f3ada206260828354e76e2080396bf18f2b1

	// Output:
	// 0xf8990b0201010cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75808001b845f8431ba0dfc662df298edef48a1a9623a735b55b3acd32023c24a40efc90a85d37209d04a06c2bcd518c2e47ac102941776a68f3ada206260828354e76e2080396bf18f2b1
}
