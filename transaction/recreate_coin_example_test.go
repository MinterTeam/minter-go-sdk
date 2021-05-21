package transaction_test

import (
	"fmt"
	"math/big"

	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewRecreateCoinData() {
	data := transaction.NewRecreateCoinData().
		SetSymbol("SPRTEST").
		SetInitialAmount(transaction.BipToPip(big.NewInt(100))).
		SetInitialReserve(transaction.BipToPip(big.NewInt(20000))).
		SetConstantReserveRatio(10).
		SetMaxSupply(transaction.BipToPip(big.NewInt(1000)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf87d0102010110adec808a5350525445535400000089056bc75e2d631000008a043c33c19375648000000a893635c9adc5dea00000808001b845f8431ca04743e4b01fc1c8305bbe9e84f483fb4a7411c419f9ec73124e4e75579a6fd5e0a06d241ed5b6a8c1b9154e7e1cba57de520fc6b5681b2aaa28f578b4e5c071c36b
}
