package wallet

import (
	"encoding/hex"
	"strings"
	"testing"
)

const (
	validSeed       = "33fa1096997d9b0f47469463710b3a2e91971144265b281dc71f831539a3b8e3413e5969e5ffb4d3c5a37cfa0f964bcc779efe4ae37fceef048175105caad624"
	validPrivateKey = "d3520cc797f12b8a81e805ddf5a5bf8b994e347003ea25c9ccaecb5073f3fef1"
	validPublicKey  = "Mp12a4d537e4b8bc98a99537a130a5d871536762bb8505a4935edc75fef97b1205977fffb0b3dbe92b7a18be8bd2068c9d729dab87cd1d76ac4522b58328da832d"
	validAddress    = "Mx48f502a9fc324f2c707edc3a2595e72f00c3190c"
)

func TestWallet_PrivateKey(t *testing.T) {
	bytes, err := hex.DecodeString(validSeed)
	if err != nil {
		t.Fatal(err)
	}
	wallet, err := NewWallet(bytes)
	if err != nil {
		t.Fatal(err)
	}

	privateKey := wallet.PrivateKey()
	if privateKey != validPrivateKey {
		t.Fatalf("privateKey got %s, want %s", privateKey, validPrivateKey)
	}
}

func TestWallet_PublicKey(t *testing.T) {
	bytes, err := hex.DecodeString(validSeed)
	if err != nil {
		t.Fatal(err)
	}
	wallet, err := NewWallet(bytes)
	if err != nil {
		t.Fatal(err)
	}

	publicKey := wallet.PublicKey()
	if publicKey != validPublicKey {
		t.Fatalf("publicKey got %s, want %s", publicKey, validPublicKey)
	}
}

func TestWallet_Address(t *testing.T) {
	bytes, err := hex.DecodeString(validSeed)
	if err != nil {
		t.Fatal(err)
	}
	wallet, err := NewWallet(bytes)
	if err != nil {
		t.Fatal(err)
	}

	address := wallet.Address()
	if address != validAddress {
		t.Fatalf("address got %s, want %s", address, validAddress)
	}
}

func TestCreate(t *testing.T) {
	data, err := Create()
	if err != nil {
		t.Fatal(err)
	}

	if len(strings.Fields(data.Mnemonic)) != 12 {
		t.Errorf("mnemonic count got %d, want %d", len(strings.Fields(data.Mnemonic)), 12)
	}
	if len(data.Seed) != 128 {
		t.Errorf("seed len got %d, want %d", len(data.Seed), 128)
	}
	if len(data.PrivateKey) != 64 {
		t.Errorf("PrivateKey len got %d, want %d", len(data.PrivateKey), 64)
	}
	if len(data.PublicKey) != 130 {
		t.Errorf("PublicKey len got %d, want %d", len(data.PublicKey), 130)
	}
	pubKeyByPrivate, err := PublicKeyByPrivateKey(data.PrivateKey)
	if pubKeyByPrivate != data.PublicKey {
		t.Errorf("pubKeyByPrivate len got %s, want %s", pubKeyByPrivate, data.PublicKey)
	}
}
