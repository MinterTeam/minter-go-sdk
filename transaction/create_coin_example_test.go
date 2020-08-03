package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewCreateCoinData() {
	data := transaction.NewCreateCoinData().
		SetName("SUPER TEST").
		SetSymbol("SPRTEST").
		SetInitialAmount(big.NewInt(0).Mul(big.NewInt(100), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetInitialReserve(big.NewInt(0).Mul(big.NewInt(20000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetConstantReserveRatio(10).
		SetMaxSupply(big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)

	// Output:
	// 0xf8870102010105b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a043c33c19375648000000a893635c9adc5dea00000808001b845f8431ba034615f080a026ee579395aeb4c2eac974a14c091f1bb112629b2b5be0a82628da07f3347c71fa0668d01126dfae49d2b402067275878e4ffd26fd42a73cdf01950
}
