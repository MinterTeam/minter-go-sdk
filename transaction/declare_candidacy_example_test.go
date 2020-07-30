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
		SetCoin(1).
		SetStake(big.NewInt(0).Mul(big.NewInt(5), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetAddress("Mx9f7fd953c2c69044b901426831ed03ee0bd0597a")

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf8940102010106b843f841949f7fd953c2c69044b901426831ed03ee0bd0597aa00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a430a01884563918244f40000808001b845f8431ba05c5650c040fbce0d4a923a5b0b7fdd4bf52156736a50e7bc1b175a4cbed68c60a07dd76883673bd41fa1293dc38313ec5288dd23c67281f851a67e83fcb917faa7
}
