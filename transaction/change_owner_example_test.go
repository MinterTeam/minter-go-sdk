package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

func ExampleNewChangeOwnerData() {
	data := transaction.NewChangeOwnerData().
		SetSymbol("SPRTEST").
		MustSetNewOwner("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8710102010111a1e08a535052544553540000009489e5dc185e6bab772ac8e00cf3fb3f4cb0931c47808001b845f8431ba03a08817c5b87aaa4eed6ed0c0a86270bbeb0a9f309453ef996546cbee383e8a0a01d34934dba2d1cb07740156d5af8a9e37b68524fd57e4eca1218d12854350f97
}
