package transaction_test

import (
	"fmt"

	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewEditCandidatePublicKeyData() {
	data := transaction.NewEditCandidatePublicKeyData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetNewPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8950102010114b844f842a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ba0933744d45b17ee431460ce6ce9a707e84f7c43b79bd7eb2067bc35352035e141a07731ab3e52be64a41482409ee3d92f6662e2e6d36ad6789c69fb1d4ec75e766f
}
