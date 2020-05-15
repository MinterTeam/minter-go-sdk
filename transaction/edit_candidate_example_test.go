package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

func ExampleNewEditCandidateData_SignTransaction() {
	data := transaction.NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8a80102018a4d4e54000000000000000eb84df84ba04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b89489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0808001b845f8431ba01408280f2cb444eb9b150026992eeaf0cf10fc64ac5dad77ec8bd87ff7ffb02ca03947d22bd5c50f5ddba4ccf2c62dcf3aa99a78913c01069aa300c04fa6a99018
}
