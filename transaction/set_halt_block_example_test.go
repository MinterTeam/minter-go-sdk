package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewSetHaltBlockData() {
	data := transaction.NewSetHaltBlockData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetHeight(123456)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf875010201010fa6e5a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a438301e240808001b844f8421ca03732a48f4c52c2ec63741ccab959d77c8618b34cbf95f6b8bd1ae57ac87888939fb1551f624c8b7519231716cd6ba9166d14fe27a1e55863ae22f7a3f6d88bfa
}
