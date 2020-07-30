package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

func ExampleNewCreateMultisigData() {
	data := transaction.NewCreateMultisigData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx772fd5bd06356250e5efe572b6ae615860ee0c17", 3).
		MustAddSigData("Mx9c7f68ff71b5819c41e8f87cc99bdf6359da3d75", 5).
		SetThreshold(7)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	msigAddress := data.Address()
	fmt.Println(msigAddress)
	// Result: Mxd43eef7b9406762aa031b82ed0b1082264a13934

	signedTx, _ := tx.SetNonce(11).SetGasPrice(1).SetGasCoin(1).SetSignatureType(transaction.SignatureTypeSingle).
		SetPayload([]byte(fmt.Sprintf("Multisig Address %s", msigAddress))).Sign("ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")

	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Result: 0xf8d50b0201010cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75b83b4d756c74697369672041646472657373204d78643433656566376239343036373632616130333162383265643062313038323236346131333933348001b845f8431ba00bbe6869cb94f85cf2ee2df274651f6810756ac8559528cff617ec82858b0208a041d798e0f87442581ac49992a11d17ffee1920b678ff5ec2ef56e66938620983

	// Output:
	// Mxd43eef7b9406762aa031b82ed0b1082264a13934
	// 0xf8d50b0201010cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75b83b4d756c74697369672041646472657373204d78643433656566376239343036373632616130333162383265643062313038323236346131333933348001b845f8431ba00bbe6869cb94f85cf2ee2df274651f6810756ac8559528cff617ec82858b0208a041d798e0f87442581ac49992a11d17ffee1920b678ff5ec2ef56e66938620983
}
