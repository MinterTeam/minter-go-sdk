package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
	"math/big"
)

func ExampleNewDeclareCandidacyData() {
	data, _ := transaction.NewDeclareCandidacyData().
		MustSetPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		SetCommission(10).
		SetCoin("MNT").
		SetStake(big.NewInt(0).Mul(big.NewInt(5), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetAddress("Mx9f7fd953c2c69044b901426831ed03ee0bd0597a")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin("MNT").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8a80102018a4d4e540000000000000006b84df84b949f7fd953c2c69044b901426831ed03ee0bd0597aa00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a430a8a4d4e5400000000000000884563918244f40000808001b845f8431ca06994a0bc24bb1a492db8d037d2f046f1436a6c166e59540db8de6230cc581e5ea03e5cdf8db2c7ef486438ef5fad3ccbdf5aa778a96de7bc22c2dc8b4e32a4a531
}
