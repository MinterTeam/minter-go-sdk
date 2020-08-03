package transaction

import (
	"math/big"
	"testing"
)

func TestTransactionCreateCoin_Sign(t *testing.T) {
	data := NewCreateCoinData().
		SetName("SUPER TEST").
		SetSymbol("SPRTEST").
		SetInitialAmount(big.NewInt(0).Mul(big.NewInt(100), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetInitialReserve(big.NewInt(0).Mul(big.NewInt(20000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		SetConstantReserveRatio(10).
		SetMaxSupply(big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8870102010105b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a043c33c19375648000000a893635c9adc5dea00000808001b845f8431ba034615f080a026ee579395aeb4c2eac974a14c091f1bb112629b2b5be0a82628da07f3347c71fa0668d01126dfae49d2b402067275878e4ffd26fd42a73cdf01950"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestEncode_CreateCoinData(t *testing.T) {
	decode, err := Decode("0xf8870102010105b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a021e19e0c9bab24000000a893635c9adc5dea00000808001b845f8431ca0c51fc88e44ab23c5da31264bc749ac513e8817701727a1e037d6d9b3708b11aaa059e7b32a726a87f31d92c7eeb0c4624fbf270de77f77c46bbd99232c2043c49a")
	if err != nil {
		t.Fatal(err)
	}

	if decode.Data().(*CreateCoinData).MaxSupply.String() != "1000000000000000000000" {
		t.Errorf("MaxSupply got %s, want %s", decode.Data().(*CreateCoinData).MaxSupply.String(), "1000000000000000000000")
	}
}
