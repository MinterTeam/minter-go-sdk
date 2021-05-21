package transaction_test

import (
	"fmt"

	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
)

func ExampleNewPriceVoteData() {
	data := transaction.NewPriceVoteData().
		SetPrice(1)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf852010201011382c101808001b845f8431ba00e6ceba5074a56661daf2099872627973e9ee09f82519893a1fda16717b4eec1a00664a550774a27d6f6a56c58d53d39ff46719ddd53423a371339314a65857196
}
