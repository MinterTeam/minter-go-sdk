package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewUnbondData_SignTransaction() {
	data := transaction.NewUnbondData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCoin("MNT").
		SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8900102018a4d4e540000000000000008b6f5a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a438a4d4e5400000000000000888ac7230489e80000808001b845f8431ca01d103930ab6ba067bb0ea4ac67d520a7eb6b2cc547860404a194777a7772659aa04b3968aa5a8455fe1ddc09b6a850624b27ce9acafac8962a90b438ab8d963439
}
