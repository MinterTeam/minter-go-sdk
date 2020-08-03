package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
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
	// 0xf8b4010201010eb863f861a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8809489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0941b685a7c1e78726c48f619c497a07ed75fe00483808001b845f8431ca0686c5c11a26d5afd9161b7cf895c86c3d7e3c6110e29b1ad5c5df8ac52a70691a008f221e93111f923dd12e695de45fbbc81e86a0aece1a104c2438913c28f31df
}

func ExampleNewEditCandidateData_setPubKey() {
	data := transaction.NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetNewPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47").
		MustSetControlAddress("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8d4010201010eb883f881a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a439489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0941b685a7c1e78726c48f619c497a07ed75fe00483808001b845f8431ba0eba98e815d70b8cc42ca16f6f9411d05b0f3ab7f354bfb788afc3cce8d0de8c3a0505f0ecfebdbcdf433586e7cec5796aacbe61f3a6f5dc1ffaf40fe4889907aac
}
