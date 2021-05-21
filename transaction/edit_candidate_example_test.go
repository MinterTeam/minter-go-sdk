package transaction_test

import (
	"fmt"

	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewEditCandidateData() {
	data := transaction.NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47").
		MustSetControlAddress("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8b3010201010eb862f860a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b89489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0941b685a7c1e78726c48f619c497a07ed75fe00483808001b845f8431ca07e6ea59db677f322752f0c62b867b8fd35db69124849c5c6eb037850031f3781a04f14493da1f218205935835b14d9ccca77614d51547a3196a94778ff6f69ee7d
}
