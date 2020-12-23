package transaction

import (
	"math/big"
	"testing"
)

const txCreateCoin = "0xf8870102010105b7f68a535550455220544553548a5350525445535400000089056bc75e2d631000008a043c33c19375648000000a893635c9adc5dea00000808001b845f8431ba034615f080a026ee579395aeb4c2eac974a14c091f1bb112629b2b5be0a82628da07f3347c71fa0668d01126dfae49d2b402067275878e4ffd26fd42a73cdf01950"

func TestTransactionCreateCoin_Sign(t *testing.T) {
	data := NewCreateCoinData().
		SetName("SUPER TEST").
		SetSymbol("SPRTEST").
		SetInitialAmount(BipToPip(big.NewInt(100))).
		SetInitialReserve(BipToPip(big.NewInt(20000))).
		SetConstantReserveRatio(10).
		SetMaxSupply(BipToPip(big.NewInt(1000)))

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	transaction := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1)

	signedTx, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != txCreateCoin {
		t.Errorf("EncodeTx got %s, want %s", encode, txCreateCoin)
	}
}

func TestDecode_createCoin(t *testing.T) {
	decode, err := Decode(txCreateCoin)
	if err != nil {
		t.Fatal(err)
	}
	if decode.GetTransaction().Type != TypeCreateCoin {
		t.Error("create coin transaction type is invalid", decode.GetTransaction().Type)
	}
	if decode.Fee().String() != "10000000000000000000000" {
		t.Error("create coin transaction fee is invalid", decode.Fee().String())
	}
}

func TestCreateCoinData_Fee(t *testing.T) {
	data := NewCreateCoinData().SetSymbol("AAA")
	if data.Fee() != 1000000000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAAB")
	if data.Fee() != 100000000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAABC")
	if data.Fee() != 10000000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAABCD")
	if data.Fee() != 1000000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAABCDE")
	if data.Fee() != 100000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAABCDEF")
	if data.Fee() != 100000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAABCDEFG")
	if data.Fee() != 100000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
	data = NewCreateCoinData().SetSymbol("AAABCDEFGH")
	if data.Fee() != 100000 {
		t.Error("create coin data fee is invalid", data.Fee())
	}
}
