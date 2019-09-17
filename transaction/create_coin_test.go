package transaction

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestTransactionCreateCoin_Sign(t *testing.T) {
	data := NewCreateCoinData().
		SetName("SUPER TEST").
		SetSymbol("SPRTEST").
		SetInitialAmount(big.NewInt(100)).
		SetInitialReserve(big.NewInt(10)).
		SetConstantReserveRatio(10)

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

	validSignature := "0xf8850102018a4d4e540000000000000005abea8a535550455220544553548a5350525445535400000089056bc75e2d63100000888ac7230489e800000a808001b845f8431ca0a0b58787e19d8ef3cbd887936617af5cf069a25a568f838c3d04daf5ad2f6f8ea07660c13ab5017edb87f5b52be4574c8a33a893bac178adec9c262a1408e4f1fe"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
