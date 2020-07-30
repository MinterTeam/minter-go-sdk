package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

func ExampleNewEditCandidateData() {
	data := transaction.NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8b4010201010eb863f861a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8809489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0940000000000000000000000000000000000000000808001b845f8431ca083dd40a3f66f507f9167d5e89cfb031ec6d9e776b4288288cc417e0c58ac38cfa048e5916b9ccad893d1dbe6e7c8ad23afda9b922c3c263a9c93f773cd2e6c59be
}
