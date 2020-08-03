package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewSellCoinData() {
	data := transaction.NewSellCoinData().
		SetCoinToSell(1).
		SetValueToSell(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetCoinToBuy(2).
		SetMinimumValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf865010201010295d401880de0b6b3a764000002880de0b6b3a7640000808001b845f8431ca01552ab0503f8173bef46f2336d48ef6e1fae7bb5aa8b51ec7332b720a8a2f15ca0166970c5d209bac8b5ffae32047f1e4e868c5a20f522aeebb0bc523ae16c64fa
}
