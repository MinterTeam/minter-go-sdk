package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewRecreateCoinData() {
	data := transaction.NewRecreateCoinData().
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
	// 0xf87c0102010110aceb8a5350525445535400000089056bc75e2d631000008a043c33c19375648000000a893635c9adc5dea00000808001b845f8431ca0f9b7c10e95dab3c124b61973bfef3be774bfe2abb06b4292a75ba5d1d25c6b88a046952529c7e20dd112ae5b9f94b27de0ec39fa93eb59ceaad68761d36b6bdc4f
}
