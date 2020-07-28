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

	validSignature := "0xf895010201010db844f842f840df0194fe60014a6e9ac91618f5d1cab3fd58cded61ee9988016345785d8a0000df0194ddab6281766ad86497741ff91b6b48fe85012e3c8802c68af0bb140000808001b845f8431ba0718f10b4468989919adabd215f5a6e83bd70eb1358d725541c2661122f66350ba05ab9e5e28107612f89ce56f4d7846edcbf272e8929eaf0c7c945e2530f40b667"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if bytes != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
