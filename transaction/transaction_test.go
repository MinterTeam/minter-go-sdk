package transaction

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestTransaction_Hash(t *testing.T) {
	signedTransaction, err := Decode("0xf872010201010aa2e1a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ba0efff777e61a78141ceeab311776dfd0bfc6745f125c688db86ccfa350d3d3b84a074419c32dd0d1d2ebdc1c5bfdffb238d2ef88a618e28a2ce2410880264d3b3cc")
	if err != nil {
		t.Fatal(err)
	}

	hash, err := signedTransaction.Hash()
	if err != nil {
		t.Fatal(err)
	}

	validHash := "Mtd018688eceaed7b7f09b4a0b7b8fadd208f06d1f040eef1048d7beea2370790d"
	if hash != validHash {
		t.Errorf("Hash got %s, want %s", hash, validHash)
	}
}

func TestTransaction_Hash2(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewSendData().
		SetCoin(1).
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
	)
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.
		SetGasPrice(1).
		SetGasCoin(1).
		SetNonce(1).
		Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	hash, err := signedTransaction.Hash()
	if err != nil {
		t.Fatal(err)
	}

	validHash := "Mtec2166cced36276426360a79934fbf49f29f9e48e9d1f06ef4afc4f557aa3767"
	if hash != validHash {
		t.Errorf("Hash got %s, want %s", hash, validHash)
	}
}

func TestTransaction_Encode(t *testing.T) {
	transaction, err := Decode("0xf865010201010495d402880de0b6b3a764000001880de0b6b3a7640000808001b845f8431ca0ad334ececd68741f1f9b96e15a2b5d6a7fe6c378cdaab6c6e8947541e1af74dda038c829477eb261948598fd3dd039aba41aa5691f50d3ee2bb4125bc38b294725")
	if err != nil {
		t.Fatal(err)
	}

	if transaction.ChainID != TestNetChainID {
		t.Errorf("ChainID got %d, want %d", transaction.ChainID, TestNetChainID)
	}

	if transaction.Type != TypeBuyCoin {
		t.Errorf("Type got %d, want %d", transaction.Type, TypeSend)
	}

	signature, err := transaction.Signature()
	if err != nil {
		t.Fatal(err)
	}

	if signature.(*Signature).V.String() != big.NewInt(28).String() {
		t.Errorf("signature get %+v, want signature.V %d", signature, 28)
	}
}

func TestObject_SenderAddress(t *testing.T) {
	transaction, err := Decode("0xf865010201010495d402880de0b6b3a764000001880de0b6b3a7640000808001b845f8431ca0ad334ececd68741f1f9b96e15a2b5d6a7fe6c378cdaab6c6e8947541e1af74dda038c829477eb261948598fd3dd039aba41aa5691f50d3ee2bb4125bc38b294725")
	if err != nil {
		t.Fatal(err)
	}

	senderAddress, err := transaction.SenderAddress()
	if err != nil {
		t.Fatal(err)
	}

	if senderAddress != "Mx622e1e0e788f4b1258f7e2a196f738c6a360c3de" {
		t.Errorf("SenderAddress want %s,\ngot %s", "Mx622e1e0e788f4b1258f7e2a196f738c6a360c3de", senderAddress)
	}
}

func TestObject_Fee_Send(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewSendData())
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	if signedTransaction.Fee().String() != "10000000000000000" {
		t.Errorf("Fee want %s, got %s", "10000000000000000", signedTransaction.Fee().String())
	}
}

func TestObject_Fee_Payload(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewSendData())
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.SetPayload([]byte("asd")).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	if signedTransaction.Fee().String() != "16000000000000000" {
		t.Errorf("Fee want %s, got %s", "16000000000000000", signedTransaction.Fee().String())
	}
}

func TestObject_Fee_PayloadUTF8(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewSendData())
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.SetPayload([]byte("asé")).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	if signedTransaction.Fee().String() != "18000000000000000" {
		t.Errorf("Fee want %s, got %s", "18000000000000000", signedTransaction.Fee().String())
	}
}

func TestCreateCoinData_Fee_3symbol(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewCreateCoinData().SetSymbol("ABC"))
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	if signedTransaction.Fee().String() != "1000000000000000000000000" {
		t.Errorf("Fee want %s, got %s", "1000000000000000000000000", signedTransaction.Fee().String())
	}
}

func TestObject_Fee_Multisend(t *testing.T) {
	data := NewMultisendData().
		AddItem(NewSendData()).
		AddItem(NewSendData()).
		AddItem(NewSendData()).
		AddItem(NewSendData()).
		AddItem(NewSendData())

	multisendTransaction, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := multisendTransaction.SetGasPrice(1).Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	fee := signedTransaction.Fee().String()
	feeValid := big.NewInt(0).Mul(big.NewInt(30), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-3), nil)).String()
	if fee != feeValid {
		t.Errorf("Fee want %s, got %s", fee, feeValid)
	}
}

func TestMultisigSig(t *testing.T) {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mx1b685a7c1e78726c48f619c497a07ed75fe00483"
	coinID := CoinID(1)
	data, err := NewSendData().
		SetCoin(coinID).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(coinID).SetSignatureType(SignatureTypeMulti)

	signedTx, err := tx.Sign("Mxade8c935e6e33b3ffd775cfa7612d89f3cde21b4",
		"ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
		"b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
		"4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901130102010101a0df01941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e694ade8c935e6e33b3ffd775cfa7612d89f3cde21b4f8cff8431ca07dd407fa5d2a161581d03cdeb7c94fcd5ade47d376af75f2c92d1483f821fe2ca00d16b6cdbceaadcd0fd72bd39ee17841871da333a571535fccfbcf6285881c2af8431ba07c2d063126024a1e19363e7e254312ca9ab37795b06102da25bd1c0dec81a934a043b7bec83db41c594ac7a8d416fca2f83f0e65ada1221fe659ba4dbe1f3c921af8431ba09318e56a242c39c10ce87ab51d10322cc62cf23885867bc89a24e8c3fa8483e9a04c82c1224d1b4efa7fba06623da2896745ce444d35ed77973759e6404b66bb95"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestMultisigSig1(t *testing.T) {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mxd82558ea00eb81d35f2654953598f5d51737d31d"
	coinID := CoinID(1)
	data, err := NewSendData().
		SetCoin(coinID).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(coinID).SetSignatureType(SignatureTypeMulti)

	signedTx, err := tx.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2",
		"b354c3d1d456d5a1ddd65ca05fd710117701ec69d82dac1858986049a0385af9",
		"38b7dfb77426247aed6081f769ed8f62aaec2ee2b38336110ac4f7484478dccb",
		"94c0915734f92dd66acfdc48f82b1d0b208efd544fe763386160ec30c968b4af")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901130102010101a0df0194d82558ea00eb81d35f2654953598f5d51737d31d880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca02eb820086487f83a5ffd99e9fc4e32e5684c915d7b1b02e76204adb8efd452eea0107f3e6cbd2d42649128e85fed63b1e633fc6718f93d50ee7f07df580c891491f8431ba05be1f763cd25c360e14cd154887facd9170ceec0483d49c95805c75c59da0faca03e0446bac2f12dadb22f43b95cb9ce4038582f779f55c17187680c7920ef672af8431ba0716e2989d1b0e97844b271fb05a1e743722e7164c7aeb15ec76482633ff0bc2aa0210628724dbf267ea82228d784eb1d72cf9c8f4353804af1f2f56d94ca20fdbc"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestMultisigSigTwoTimeSig(t *testing.T) {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mxd82558ea00eb81d35f2654953598f5d51737d31d"
	coinID := CoinID(1)
	data, err := NewSendData().
		SetCoin(coinID).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(coinID).SetSignatureType(SignatureTypeMulti)

	signedTx1, err := tx.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2",
		"b354c3d1d456d5a1ddd65ca05fd710117701ec69d82dac1858986049a0385af9")
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := signedTx1.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2",
		"38b7dfb77426247aed6081f769ed8f62aaec2ee2b38336110ac4f7484478dccb",
		"94c0915734f92dd66acfdc48f82b1d0b208efd544fe763386160ec30c968b4af")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901130102010101a0df0194d82558ea00eb81d35f2654953598f5d51737d31d880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca02eb820086487f83a5ffd99e9fc4e32e5684c915d7b1b02e76204adb8efd452eea0107f3e6cbd2d42649128e85fed63b1e633fc6718f93d50ee7f07df580c891491f8431ba05be1f763cd25c360e14cd154887facd9170ceec0483d49c95805c75c59da0faca03e0446bac2f12dadb22f43b95cb9ce4038582f779f55c17187680c7920ef672af8431ba0716e2989d1b0e97844b271fb05a1e743722e7164c7aeb15ec76482633ff0bc2aa0210628724dbf267ea82228d784eb1d72cf9c8f4353804af1f2f56d94ca20fdbc"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestMultisigAddSignatures(t *testing.T) {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mxd82558ea00eb81d35f2654953598f5d51737d31d"
	coinID := CoinID(1)
	data, err := NewSendData().
		SetCoin(coinID).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(coinID).SetMultiSignatureType()
	transaction0 := *(tx.(*Object).Transaction)
	tx0 := Object{Transaction: &transaction0}

	transaction1 := *(tx.(*Object).Transaction)
	tx1 := Object{Transaction: &transaction1}
	signedTx1, err := tx1.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2", "b354c3d1d456d5a1ddd65ca05fd710117701ec69d82dac1858986049a0385af9")
	if err != nil {
		t.Fatal(err)
	}

	transaction2 := *(tx.(*Object).Transaction)
	tx2 := Object{Transaction: &transaction2}
	signedTx2, err := tx2.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2", "38b7dfb77426247aed6081f769ed8f62aaec2ee2b38336110ac4f7484478dccb")
	if err != nil {
		t.Fatal(err)
	}

	transaction3 := *(tx.(*Object).Transaction)
	tx3 := Object{Transaction: &transaction3}
	signedTx3, err := tx3.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2", "94c0915734f92dd66acfdc48f82b1d0b208efd544fe763386160ec30c968b4af")
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := tx0.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2")
	if err != nil {
		t.Fatal(err)
	}
	simpleSignatureData1, err := signedTx1.SingleSignatureData()
	if err != nil {
		t.Fatal(err)
	}
	simpleSignatureData2, err := signedTx2.SingleSignatureData()
	if err != nil {
		t.Fatal(err)
	}
	simpleSignatureData3, err := signedTx3.SingleSignatureData()
	if err != nil {
		t.Fatal(err)
	}

	signedTx0, err := signedTransaction.AddSignature(simpleSignatureData1, simpleSignatureData2, simpleSignatureData3)
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901130102010101a0df0194d82558ea00eb81d35f2654953598f5d51737d31d880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca02eb820086487f83a5ffd99e9fc4e32e5684c915d7b1b02e76204adb8efd452eea0107f3e6cbd2d42649128e85fed63b1e633fc6718f93d50ee7f07df580c891491f8431ba05be1f763cd25c360e14cd154887facd9170ceec0483d49c95805c75c59da0faca03e0446bac2f12dadb22f43b95cb9ce4038582f779f55c17187680c7920ef672af8431ba0716e2989d1b0e97844b271fb05a1e743722e7164c7aeb15ec76482633ff0bc2aa0210628724dbf267ea82228d784eb1d72cf9c8f4353804af1f2f56d94ca20fdbc"
	encode, err := signedTx0.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestMultisigAddSignatures2(t *testing.T) {
	value := big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))
	address := "Mxd82558ea00eb81d35f2654953598f5d51737d31d"
	coinID := CoinID(1)
	data, err := NewSendData().
		SetCoin(coinID).
		SetValue(value).
		SetTo(address)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(coinID).SetMultiSignatureType()

	signedTransaction, err := tx.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2")
	if err != nil {
		t.Fatal(err)
	}

	sig1, err := hex.DecodeString("f8431ca0a116e33d2fea86a213577fc9dae16a7e4cadb375499f378b33cddd1d4113b6c1a021ee1e9eb61bbd24233a0967e1c745ab23001cf8816bb217d01ed4595c6cb2cd")
	if err != nil {
		t.Fatal(err)
	}
	sig2, err := hex.DecodeString("f8431ca0f7f9c7a6734ab2db210356161f2d012aa9936ee506d88d8d0cba15ad6c84f8a7a04b71b87cbbe7905942de839211daa984325a15bdeca6eea75e5d0f28f9aaeef8")
	if err != nil {
		t.Fatal(err)
	}
	sig3, err := hex.DecodeString("f8431ba0d8c640d7605034eefc8870a6a3d1c22e2f589a9319288342632b1c4e6ce35128a055fe3f93f31044033fe7b07963d547ac50bccaac38a057ce61665374c72fb454")
	if err != nil {
		t.Fatal(err)
	}
	signedTx0, err := signedTransaction.AddSignature(sig1, sig2, sig3)
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901130102010101a0df0194d82558ea00eb81d35f2654953598f5d51737d31d880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca0a116e33d2fea86a213577fc9dae16a7e4cadb375499f378b33cddd1d4113b6c1a021ee1e9eb61bbd24233a0967e1c745ab23001cf8816bb217d01ed4595c6cb2cdf8431ca0f7f9c7a6734ab2db210356161f2d012aa9936ee506d88d8d0cba15ad6c84f8a7a04b71b87cbbe7905942de839211daa984325a15bdeca6eea75e5d0f28f9aaeef8f8431ba0d8c640d7605034eefc8870a6a3d1c22e2f589a9319288342632b1c4e6ce35128a055fe3f93f31044033fe7b07963d547ac50bccaac38a057ce61665374c72fb454"
	encode, err := signedTx0.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestDecodeMulti(t *testing.T) {
	decode, err := Decode("0xf899010201010cb848f84607c3010305f83f94ee81347211c72524338f9680072af9074433314394ee81347211c72524338f9680072af9074433314594ee81347211c72524338f9680072af90744333144808001b845f8431ca0224c6166a1f4667cb0bee9ce7ed88879285b8ffc9b4eac3f03faa1797d1f8684a0276dc68fc640924e970c3607af33988a0955e7c2dff78a16ba795da9ddffe988")
	if err != nil {
		t.Fatal(err)
	}
	senderAddress, err := decode.SenderAddress()
	validSenderAddress := "Mxb43154a0bc801c4b7361bf1a535b5e08e34e401c"
	if senderAddress != validSenderAddress {
		t.Fatalf("SenderAddress got %s, want %s", senderAddress, validSenderAddress)
	}
}
