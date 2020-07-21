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
		SetInitialReserve(big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
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

	validSignature := "0xf8910102018a4d4e540000000000000005b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a021e19e0c9bab24000000a893635c9adc5dea00000808001b845f8431ba07bf9c6916aabaac7fb34811b42350c0dbcfc6228cf2ce9b927254d01f9e0ec66a0039ea86546a950cd717544d9b19c30a5248cfeb0f93060145144b5bb511a4218"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestEncode_CreateCoinData(t *testing.T) {
	decode, err := Decode("0xf8910102018a4d4e540000000000000005b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a021e19e0c9bab24000000a893635c9adc5dea00000808001b845f8431ba07bf9c6916aabaac7fb34811b42350c0dbcfc6228cf2ce9b927254d01f9e0ec66a0039ea86546a950cd717544d9b19c30a5248cfeb0f93060145144b5bb511a4218")
	if err != nil {
		t.Fatal(err)
	}

	if decode.Data().(*CreateCoinData).MaxSupply.String() != "1000000000000000000000" {
		t.Errorf("MaxSupply got %s, want %s", decode.Data().(*CreateCoinData).MaxSupply.String(), "1000000000000000000000")
	}
}
