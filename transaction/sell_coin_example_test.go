package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewSellCoinData_SignTransaction() {
	data := transaction.NewSellCoinData().
		SetCoinToSell("MNT").
		SetValueToSell(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetCoinToBuy("TEST").
		SetMinimumValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))
	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8830102018a4d4e540000000000000002a9e88a4d4e5400000000000000880de0b6b3a76400008a54455354000000000000880de0b6b3a7640000808001b845f8431ba0e34be907a18acb5a1aed263ef419f32f5adc6e772b92f949906b497bba557df3a0291d7704980994f7a6f5950ca84720746b5928f21c3cfc5a5fbca2a9f4d35db0
}
