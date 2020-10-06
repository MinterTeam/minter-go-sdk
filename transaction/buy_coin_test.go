package transaction

import (
	"math/big"
	"testing"
)

func TestTransactionBuyCoin_Sign(t *testing.T) {
	data := NewBuyCoinData().
		SetCoinToBuy(2).
		SetValueToBuy(BipToPip(big.NewInt(1))).
		SetCoinToSell(1).
		SetMaximumValueToSell(BipToPip(big.NewInt(1)))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	transaction := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1)

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf865010201010495d402880de0b6b3a764000001880de0b6b3a7640000808001b845f8431ca0ad334ececd68741f1f9b96e15a2b5d6a7fe6c378cdaab6c6e8947541e1af74dda038c829477eb261948598fd3dd039aba41aa5691f50d3ee2bb4125bc38b294725"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}
