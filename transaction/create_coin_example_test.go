package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewCreateCoinData() {
	data := transaction.NewCreateCoinData().
		SetName("SUPER TEST").
		SetSymbol("SPRTEST").
		SetInitialAmount(big.NewInt(0).Mul(big.NewInt(100), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetInitialReserve(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetConstantReserveRatio(10).
		SetMaxSupply(big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)

	// Output:
	// 0xf8850102010105b5f48a535550455220544553548a5350525445535400000089056bc75e2d63100000888ac7230489e800000a893635c9adc5dea00000808001b845f8431ba0e6fc1f864b0a7bd0ac385c11e059aaab1c8507b4a7b05233e226663ba8fc22bda044f4ca23b492d038e7e5d0075293b23d9aa585cab5740f75aa8f8bd1aaae6eca
}
