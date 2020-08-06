package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"math/big"
)

func ExampleNewRedeemCheckData() {
	check := transaction.NewCheck("02c0c0a4-8023-4654-ac5d-3c39ba1bd19c", transaction.MainNetChainID, 500, 0, big.NewInt(9223372036854775807), 0)
	checkSigned, _ := check.SetPassphrase("secret").Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	checkSignedEncode, _ := checkSigned.Encode()

	pr, _ := transaction.NewCheckAddress("Mx31e61A05adBD13c6B625262704bc305Bf7725026", "secret")
	proof, _ := pr.Proof()

	data := transaction.NewRedeemCheckData().
		MustSetProof(proof).
		MustSetRawCheck(checkSignedEncode)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)
	signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Output:
	// 0xf901560102010109b90104f90101b8bcf8baa430326330633061342d383032332d343635342d616335642d336333396261316264313963018201f480887fffffffffffffff80b84164a1ab995f8d8883fa6e7d41e57774f5bcef72aa29cdf61875a4276250f610796d2a50e8dcadfe346fc8d3a555ff97fe3e6f3899879c1d852da6a444609bc1bc001ca0751931e85b03d4ecec72a7b1b2a9b1f92b7a7d260662d9822019cb4b1654f9d0a05e79fdcf59b85fd4637173d6a219255bf8e22898efa12227930c3a761e6225b5b8414aa2806ff70943c8a405011c76a416bf328d642b3ad656ef8cd7d26560facdf22ae4ef789f497bc8bec49b8e56d2bf3c2b2904262858032048d35776a2909e7b00808001b845f8431ca0ceb258c5216b4e95b8bb9121ce97d1133d6003b3de6680af3017e85ee9744c4aa04767cdc002ad8c8d5719c26518e65b222d4f137fc6861d70adba0fbee0bd6688

}
