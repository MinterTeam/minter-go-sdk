package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewBuyCoinData() {
	data := transaction.NewBuyCoinData().
		SetCoinToBuy(2).
		SetValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetCoinToSell(1).
		SetMaximumValueToSell(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf865010201010495d402880de0b6b3a764000001880de0b6b3a7640000808001b845f8431ca0ad334ececd68741f1f9b96e15a2b5d6a7fe6c378cdaab6c6e8947541e1af74dda038c829477eb261948598fd3dd039aba41aa5691f50d3ee2bb4125bc38b294725
}
