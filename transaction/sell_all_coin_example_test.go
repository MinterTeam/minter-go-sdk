package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewSellAllCoinData_SignTransaction() {
	data := transaction.NewSellAllCoinData().
		SetCoinToSell("MNT").
		SetCoinToBuy("TEST").
		SetMinimumValueToBuy(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf87a0102018a4d4e540000000000000003a0df8a4d4e54000000000000008a54455354000000000000880de0b6b3a7640000808001b845f8431ca0b10794a196b6ad2f94e6162613ca9538429dd49ca493594ba9d99f80d2499765a03c1d78e9e04f57336691e8812a16faccb00bf92ac817ab61cd9bf001e9380d47
}
