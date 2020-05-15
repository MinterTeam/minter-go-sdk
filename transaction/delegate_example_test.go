package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewDelegateData_SignTransaction() {
	data := transaction.NewDelegateData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCoin("MNT").
		SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Result: 0xf8900102018a4d4e540000000000000007b6f5a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a438a4d4e5400000000000000888ac7230489e80000808001b845f8431ca0a77d4b04fb9f3b6601c28369a0c02b7a2bb5d2dd0efac5388705513b0c298f9ca076044f1fb751aff727a41959dc31336c1c6bb0751bfb8c1ef05efb027e718fc2

	// Output:
	// 0xf8900102018a4d4e540000000000000007b6f5a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a438a4d4e5400000000000000888ac7230489e80000808001b845f8431ca0a77d4b04fb9f3b6601c28369a0c02b7a2bb5d2dd0efac5388705513b0c298f9ca076044f1fb751aff727a41959dc31336c1c6bb0751bfb8c1ef05efb027e718fc2
}
