package transaction

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestTransactionBuyCoin_Sign(t *testing.T) {
	data := NewBuyCoinData().
		SetCoinToBuy("TEST").
		SetValueToBuy(big.NewInt(1)).
		SetCoinToSell("MNT").
		SetMaximumValueToSell(big.NewInt(1))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT")

	privateKey, err := hex.DecodeString("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := transaction.Sign(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8830102018a4d4e540000000000000004a9e88a54455354000000000000880de0b6b3a76400008a4d4e5400000000000000880de0b6b3a7640000808001b845f8431ca04ee095a20ca58062a5758e2a6d3941857daa8943b5873c57f111190ca88dbc56a01148bf2fcc721ca353105e4f4a3419bec471d7ae08173f443a28c3ae6d27018a"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
