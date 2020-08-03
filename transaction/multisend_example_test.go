package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewMultisendData() {
	coinID := transaction.CoinID(1)
	data := transaction.NewMultisendData().AddItem(
		transaction.NewSendData().
			SetCoin(coinID).
			SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-1), nil))).
			MustSetTo("Mxfe60014a6e9ac91618f5d1cab3fd58cded61ee99"),
	).AddItem(
		transaction.NewSendData().
			SetCoin(coinID).
			SetValue(big.NewInt(0).Mul(big.NewInt(2), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-1), nil))).
			MustSetTo("Mxddab6281766ad86497741ff91b6b48fe85012e3c"),
	)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(coinID).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf895010201010db844f842f840df0194fe60014a6e9ac91618f5d1cab3fd58cded61ee9988016345785d8a0000df0194ddab6281766ad86497741ff91b6b48fe85012e3c8802c68af0bb140000808001b845f8431ba0718f10b4468989919adabd215f5a6e83bd70eb1358d725541c2661122f66350ba05ab9e5e28107612f89ce56f4d7846edcbf272e8929eaf0c7c945e2530f40b667

}
