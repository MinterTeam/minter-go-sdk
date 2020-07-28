package transaction

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestTransactionSend_Sign(t *testing.T) {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mx1b685a7c1e78726c48f619c497a07ed75fe00483"
	data, err := NewSendData().
		SetCoin(1).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	if data.Coin != 1 {
		t.Errorf("SendData.CoinID got %s, want %d", data.Coin, 1)
	}

	addressBytes, err := hex.DecodeString(address[2:])
	if string(data.To[:]) != string(addressBytes) {
		t.Errorf("SendData.To got %s, want %s", string(data.To[:]), string(addressBytes))
	}

	if data.Value.String() != value.String() {
		t.Errorf("SendData.Value got %s, want %s", data.Value.String(), value.String())
	}
	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)
	transaction := tx.(*Object)

	if transaction.Nonce != nonce {
		t.Errorf("Nonce got %d, want %d", transaction.Nonce, nonce)
	}

	if transaction.ChainID != TestNetChainID {
		t.Errorf("ChainID got %d, want %d", transaction.ChainID, TestNetChainID)
	}

	if transaction.GasPrice != gasPrice {
		t.Errorf("GasPrice got %d, want %d", transaction.GasPrice, gasPrice)
	}

	gasCoinBytes := CoinID(1) // MNT
	if transaction.GasCoin != gasCoinBytes {
		t.Errorf("GasCoin got %s, want %s", transaction.GasCoin, gasCoinBytes)
	}

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8700102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ba0fffc3f503ace8a5d0c87efe50cf33ad41e3475459120d9c6fd75bd796b192313a0243d643a799e844ad82382d41cee98137a1d0c5888ff13951919e5e241ab89e0"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
