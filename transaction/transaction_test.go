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

	if signature.V.String() != big.NewInt(28).String() {
		t.Errorf("signature get %+v, want signature.V %d", signature, 28)
	}
}

func TestObject_PublicKey(t *testing.T) {
	transaction, err := Decode("0xf8840102018a4d4e540000000000000001aae98a4d4e5400000000000000941b685a7c1e78726c48f619c497a07ed75fe00483880de0b6b3a7640000808001b845f8431ca01f36e51600baa1d89d2bee64def9ac5d88c518cdefe45e3de66a3cf9fe410de4a01bc2228dc419a97ded0efe6848de906fbe6c659092167ef0e7dcb8d15024123a")
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := transaction.PublicKey()
	if err != nil {
		t.Fatal(err)
	}

	if publicKey != "Mp5ecb93ea4368127b6311fd73fb86df0b63c38162fd564a204611a9549059c89bd91e63be5ee1341c43e0a692fdae1da7894ab91969e13aa4cb15a5e59d2dbdee" {
		t.Errorf("PublicKey want %s,\ngot %s", "Mp5ecb93ea4368127b6311fd73fb86df0b63c38162fd564a204611a9549059c89bd91e63be5ee1341c43e0a692fdae1da7894ab91969e13aa4cb15a5e59d2dbdee", publicKey)
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
		t.Errorf("PublicKey want %s,\ngot %s", "Mx31e61a05adbd13c6b625262704bc305bf7725026", senderAddress)
	}
}
