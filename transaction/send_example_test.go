package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewSendData() {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mx1b685a7c1e78726c48f619c497a07ed75fe00483"
	symbolMNT := transaction.CoinID(1)
	data, _ := transaction.NewSendData().
		SetCoin(symbolMNT).
		SetValue(value).
		SetTo(address)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8700102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ba0fffc3f503ace8a5d0c87efe50cf33ad41e3475459120d9c6fd75bd796b192313a0243d643a799e844ad82382d41cee98137a1d0c5888ff13951919e5e241ab89e0
}
