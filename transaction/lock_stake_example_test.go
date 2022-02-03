package transaction_test

import (
	"fmt"

	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewLockStakeData() {
	data := transaction.NewLockStakeData()

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf851010201012581c0808001b845f8431ca0220af5b0dcd1f4c15fa7dda2e0ecb6e198e28f644a7f516590e4203403f8ea52a00d06d23e4a76f5c1a7581bdab75f6f136bf7df61285e9142aa2c5b7ff10fc3f2
}
