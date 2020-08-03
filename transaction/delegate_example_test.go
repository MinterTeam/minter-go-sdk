package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewDelegateData() {
	data := transaction.NewDelegateData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCoin(1).
		SetValue(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Result: 0xf87c0102010107aceba00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a4301888ac7230489e80000808001b845f8431ca016bed99d271578829200f5c7c8f2f04dbf242b1d2c4ed4676c6f3480ade1f6a3a0677cac19e1a36367358cdeeed4a66aa135615b41206f502b9284239539084b5a

	// Output:
	// 0xf87c0102010107aceba00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a4301888ac7230489e80000808001b845f8431ca016bed99d271578829200f5c7c8f2f04dbf242b1d2c4ed4676c6f3480ade1f6a3a0677cac19e1a36367358cdeeed4a66aa135615b41206f502b9284239539084b5a
}
