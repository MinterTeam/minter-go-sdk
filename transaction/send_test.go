package transaction

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSendTransaction_Sets(t *testing.T) {
	tx := NewSendTransaction()

	//tx.Data.SetTo()

	nonce := uint64(15)
	tx.SetNonce(nonce)
	if tx.Nonce != nonce {
		t.Errorf("Nonce got %d, want %d", tx.Nonce, nonce)
	}

	tx.SetChainID(TestNetChainID)
	if tx.ChainID != TestNetChainID {
		t.Errorf("ChainID got %d, want %d", tx.ChainID, TestNetChainID)
	}
	tx.SetGasPrice(1)
	gasCoinBytes := [10]byte{'\x4d', '\x4e', '\x54', '\x0a'} //MNT
	tx.SetGasCoin(string(gasCoinBytes[:]))
	if !reflect.DeepEqual(tx.GasCoin, gasCoinBytes) {
		t.Errorf("GasCoin got %s, want %s", tx.GasCoin, gasCoinBytes)
	}

	tx.SetData(NewSendData("MNT", "Mxeeee1973381ab793719fff497b9a516719fcd5a2", big.NewInt(10)))
}
