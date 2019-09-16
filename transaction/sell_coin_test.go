package transaction

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestTransactionSellCoin_Sign(t *testing.T) {
	data := NewSellCoinData().
		SetCoinToSell("MNT").
		SetValueToSell(big.NewInt(1)).
		SetCoinToBuy("TEST").
		SetMinimumValueToBuy(big.NewInt(1))

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

	key, err := ToECDSA(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := transaction.Sign(key)
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8830102018a4d4e540000000000000002a9e88a4d4e5400000000000000880de0b6b3a76400008a54455354000000000000880de0b6b3a7640000808001b845f8431ba0e34be907a18acb5a1aed263ef419f32f5adc6e772b92f949906b497bba557df3a0291d7704980994f7a6f5950ca84720746b5928f21c3cfc5a5fbca2a9f4d35db0"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
