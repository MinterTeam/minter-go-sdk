package transaction

import (
	"math/big"
	"testing"
)

func TestTransaction_Hash(t *testing.T) {
	transaction, err := NewBuilder(TestNetChainID).NewTransaction(NewSendData().
		SetCoin("MNT").
		SetValue(big.NewInt(0).Mul(big.NewInt(1), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil))).
		MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
	)
	if err != nil {
		t.Fatal(err)
	}

	signedTransaction, err := transaction.
		SetGasPrice(1).
		SetGasCoin("MNT").
		SetNonce(1).
		Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")
	if err != nil {
		t.Fatal(err)
	}

	hash, err := signedTransaction.Hash()
	if err != nil {
		t.Fatal(err)
	}

	validHash := "Mt13b73500c171006613fa8e82cc8b29857af1d63a"
	if hash != validHash {
		t.Errorf("Hash got %s, want %s", hash, validHash)
	}
}

func TestTransaction_Encode(t *testing.T) {
	transaction, err := Decode("0xf8840102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ca01f36e51600baa1d89d2bee64def9ac5d88c518cdefe45e3de66a3cf9fe410de4a01bc2228dc419a97ded0efe6848de906fbe6c659092167ef0e7dcb8d15024123a")
	if err != nil {
		t.Fatal(err)
	}

	objectTransaction, ok := transaction.(*object)
	if !ok {
		t.Fatal("error ")
	}

	if objectTransaction.ChainID != TestNetChainID {
		t.Errorf("ChainID got %d, want %d", objectTransaction.ChainID, TestNetChainID)
	}

	if objectTransaction.Type != TypeSend {
		t.Errorf("Type got %d, want %d", objectTransaction.Type, TypeSend)
	}

	signature, err := objectTransaction.Signature()
	if err != nil {
		t.Fatal(err)
	}

	if signature.(*Signature).V.String() != big.NewInt(28).String() {
		t.Errorf("signature get %+v, want signature.V %d", signature, 28)
	}
}

func TestObject_SenderAddress(t *testing.T) {
	transaction, err := Decode("0xf8840102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ca01f36e51600baa1d89d2bee64def9ac5d88c518cdefe45e3de66a3cf9fe410de4a01bc2228dc419a97ded0efe6848de906fbe6c659092167ef0e7dcb8d15024123a")
	if err != nil {
		t.Fatal(err)
	}

	senderAddress, err := transaction.SenderAddress()
	if err != nil {
		t.Fatal(err)
	}

	if senderAddress != "Mx31e61a05adbd13c6b625262704bc305bf7725026" {
		t.Errorf("SenderAddress want %s,\ngot %s", "Mx31e61a05adbd13c6b625262704bc305bf7725026", senderAddress)
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

func TestObject_Fee_Coin3Symbol(t *testing.T) {
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
		AddItem(*NewMultisendDataItem()).
		AddItem(*NewMultisendDataItem()).
		AddItem(*NewMultisendDataItem()).
		AddItem(*NewMultisendDataItem()).
		AddItem(*NewMultisendDataItem())

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
	symbolMNT := "MNT"
	data, err := NewSendData().
		SetCoin(symbolMNT).
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

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(symbolMNT).SetSignatureType(SignatureTypeMulti)

	createMultisigData := NewCreateMultisigData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx6bf192730d01a19739b5030cdb6a60c992712a59", 3).
		MustAddSigData("Mx823bb524d5702addbe13086082f7f0310e07d176", 5).
		SetThreshold(7)

	multisigAddress := createMultisigData.MultisigAddressString()
	validAddr := "Mx0023aa9371e0779189ef5a7434456fc21a938945"
	if multisigAddress != validAddr {
		t.Fatalf("Address got %s, want %s", multisigAddress, validAddr)
	}

	signedTx, err := tx.Sign(multisigAddress,
		"ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
		"b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
		"4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901270102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808002b8e8f8e6940023aa9371e0779189ef5a7434456fc21a938945f8cff8431ba0014aaffef58c3def74bbb828d7cba907df59b50a68749b8d90aa0d7520571be3a04397def13aa5a38b666d5ecf590af7fdec18663bfa448d517d6671fbe25cdde6f8431ba07bd81f68708141c01ed3bac914cc04dc07831989cb86c4b0e992ad9677bfa33aa03b0d936c268e080bbb85a70cfa6c48a88f023d9e06fa4ecfb9e3cb6659bc767af8431ba0767922509d65315ddf728da8cf5450fa8ba410680f7046405a1eeb7cf22f521aa01222a82c41f7ef51e5b6a64414078185393578f8a5373ac5f5a19ee512b9317b"
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
	symbolMNT := "MNT"
	data, err := NewSendData().
		SetCoin(symbolMNT).
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

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(symbolMNT).SetSignatureType(SignatureTypeMulti)

	signedTx, err := tx.Sign("Mxdb4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2",
		"b354c3d1d456d5a1ddd65ca05fd710117701ec69d82dac1858986049a0385af9",
		"38b7dfb77426247aed6081f769ed8f62aaec2ee2b38336110ac4f7484478dccb",
		"94c0915734f92dd66acfdc48f82b1d0b208efd544fe763386160ec30c968b4af")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901270102018a4d4e540000000000000001aae98a4d4e540000000000000094d82558ea00eb81d35f2654953598f5d51737d31d880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca0a116e33d2fea86a213577fc9dae16a7e4cadb375499f378b33cddd1d4113b6c1a021ee1e9eb61bbd24233a0967e1c745ab23001cf8816bb217d01ed4595c6cb2cdf8431ca0f7f9c7a6734ab2db210356161f2d012aa9936ee506d88d8d0cba15ad6c84f8a7a04b71b87cbbe7905942de839211daa984325a15bdeca6eea75e5d0f28f9aaeef8f8431ba0d8c640d7605034eefc8870a6a3d1c22e2f589a9319288342632b1c4e6ce35128a055fe3f93f31044033fe7b07963d547ac50bccaac38a057ce61665374c72fb454"
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
	symbolMNT := "MNT"
	data, err := NewSendData().
		SetCoin(symbolMNT).
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

	tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(symbolMNT).SetSignatureType(SignatureTypeMulti)

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

	validSignature := "0xf901270102018a4d4e540000000000000001aae98a4d4e540000000000000094d82558ea00eb81d35f2654953598f5d51737d31d880de0b6b3a7640000808002b8e8f8e694db4f4b6942cb927e8d7e3a1f602d0f1fb43b5bd2f8cff8431ca0a116e33d2fea86a213577fc9dae16a7e4cadb375499f378b33cddd1d4113b6c1a021ee1e9eb61bbd24233a0967e1c745ab23001cf8816bb217d01ed4595c6cb2cdf8431ca0f7f9c7a6734ab2db210356161f2d012aa9936ee506d88d8d0cba15ad6c84f8a7a04b71b87cbbe7905942de839211daa984325a15bdeca6eea75e5d0f28f9aaeef8f8431ba0d8c640d7605034eefc8870a6a3d1c22e2f589a9319288342632b1c4e6ce35128a055fe3f93f31044033fe7b07963d547ac50bccaac38a057ce61665374c72fb454"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}

func TestDecodeMulti(t *testing.T) {
	decode, err := Decode("0xf901b30a02018a4d4e54000000000000000cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c943211a16aeda4d92de05d1e7e981b6db63571b032940c1e7b65e79a7df40494e5df99e9b39122791b5cb9010e616530383962333265346530393736636136383838636231303233313438626431613966316363323863356434343265353265353836373534666634386436332c20313664623465363332326566313335303136643733616561636530646438333939383133316662386535316562373933633636376430313635666539343762352c20383637393933313438393066656232623034626163646662346464356565373661393432643435643864623630356266613739376634643864653261333964632c205b323436203234322035342037322032333220323334203920383320393520313038203231362032343620383020323038203433203338203231362031313220313832203131355d8001b845f8431ca0554a56ba708535b1b5ebc5c49daa17b0b97d4c9c18d8860be83dbfab954e8b06a045125d4c9d572404a4cf266ed4f9d9c0452548c72905d346679b930001f829b9")
	if err != nil {
		t.Fatal(err)
	}
	senderAddress, err := decode.SenderAddress()
	validSenderAddress := "Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c"
	if senderAddress != validSenderAddress {
		t.Errorf("SenderAddress got %s, want %s", senderAddress, validSenderAddress)
	}
	multisigAddress := decode.Data().(*CreateMultisigData).MultisigAddressString()
	validMultisigAddress := "Mxf6f23648e8ea09535f6cd8f650d02b26d870b673"
	if multisigAddress != validMultisigAddress {
		t.Errorf("SenderAddress got %s, want %s", multisigAddress, validMultisigAddress)
	}
}
