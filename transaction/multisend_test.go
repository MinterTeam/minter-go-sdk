package transaction

import (
	"math/big"
	"testing"
)

func TestTransactionMultisend_Sign(t *testing.T) {
	data := NewMultisendData().AddItem(
		NewSendData().
			SetCoin(1).
			SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-1), nil))).
			MustSetTo("Mxfe60014a6e9ac91618f5d1cab3fd58cded61ee99"),
	).AddItem(
		NewSendData().
			SetCoin(1).
			SetValue(big.NewInt(0).Mul(big.NewInt(2), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-1), nil))).
			MustSetTo("Mxddab6281766ad86497741ff91b6b48fe85012e3c"),
	)

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

	validSignature := "0xf8b30102018a4d4e54000000000000000db858f856f854e98a4d4e540000000000000094fe60014a6e9ac91618f5d1cab3fd58cded61ee9988016345785d8a0000e98a4d4e540000000000000094ddab6281766ad86497741ff91b6b48fe85012e3c8802c68af0bb140000808001b845f8431ca0b15dcf2e013df1a2aea02e36a17af266d8ee129cdcb3e881d15b70c9457e7571a0226af7bdaca9d42d6774c100b22e0c7ba4ec8dd664d17986318e905613013283"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bytes != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
