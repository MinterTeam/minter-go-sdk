package transaction

import (
	"math/big"
	"testing"
)

const txSend = "0xf8700102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ba0fffc3f503ace8a5d0c87efe50cf33ad41e3475459120d9c6fd75bd796b192313a0243d643a799e844ad82382d41cee98137a1d0c5888ff13951919e5e241ab89e0"

func TestTransactionSend_Sign(t *testing.T) {
	t.Parallel()
	value := BipToPip(big.NewInt(1))
	address := "Mx1b685a7c1e78726c48f619c497a07ed75fe00483"
	data, err := NewSendData().
		SetCoin(1).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	if data.Coin != 1 {
		t.Errorf("SendData.CoinID got %s, want %d", data.Coin.String(), 1)
	}

	if data.To.String() != address {
		t.Errorf("SendData.To got %s, want %s", data.To.String(), address)
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
	transaction := tx.(*object)

	if transaction.Nonce != nonce {
		t.Errorf("Nonce got %d, want %d", transaction.Nonce, nonce)
	}

	if transaction.ChainID != TestNetChainID {
		t.Errorf("ChainID got %d, want %d", transaction.ChainID, TestNetChainID)
	}

	if transaction.GasPrice != gasPrice {
		t.Errorf("GasPrice got %d, want %d", transaction.GasPrice, gasPrice)
	}

	if transaction.GasCoin != 1 {
		t.Errorf("GasCoin got %s, want %d", transaction.GasCoin.String(), 1)
	}

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != txSend {
		t.Errorf("EncodeTx got %s, want %s", encode, txSend)
	}
}

func TestDecode_send(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txSend)
	if err != nil {
		t.Fatal(err)
	}

	if decode.Fee().String() != "10000000000000000" {
		t.Error("send transaction fee is invalid", decode.Fee().String())
	}
}
