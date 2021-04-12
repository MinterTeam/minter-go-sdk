package wallet

import (
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
	t.Parallel()
	wallet, err := Create("", validSeed)
	if err != nil {
		t.Fatal(err)
	}

	privateKey := wallet.PrivateKey
	if privateKey != validPrivateKey {
		t.Fatalf("privateKey got %s, want %s", privateKey, validPrivateKey)
	}
}

func TestWallet_PublicKey(t *testing.T) {
	t.Parallel()
	wallet, err := Create("", validSeed)
	if err != nil {
		t.Fatal(err)
	}

	publicKey := wallet.PublicKey
	if publicKey != validPublicKey {
		t.Fatalf("publicKey got %s, want %s", publicKey, validPublicKey)
	}
}

func TestWallet_Address(t *testing.T) {
	t.Parallel()
	wallet, err := Create("", validSeed)
	if err != nil {
		t.Fatal(err)
	}

	address := wallet.Address
	if address != validAddress {
		t.Fatalf("address got %s, want %s", address, validAddress)
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()
	data, err := New()
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
	if err != nil {
		t.Fatal(err)
	}

	if pubKeyByPrivate != data.PublicKey {
		t.Errorf("pubKeyByPrivate len got %s, want %s", pubKeyByPrivate, data.PublicKey)
	}
}

func TestWalletBugAddress(t *testing.T) {
	t.Parallel()
	wallet, err := Create("real town addict extend shoot name disagree vital turn live can tip", "")
	if err != nil {
		t.Fatal(err)
	}

	if wallet.PrivateKey != "c29cd818232f40614691a58c9afcbd5518b539cb135d73b1d9e26d5f5f5dcbbe" {
		t.Fatalf("PrivateKey got %s, want %s", wallet.PrivateKey, "11c332a5107bdede093dc057b146fbb633cbe0f2c50958c9de3aee13fe6caad3")
	}

	if wallet.PublicKey != "Mp32721d081431dea637f8826fd58babe93ebc9648c76978e26ebf8b56e91292a832fbd62f569fefdc05139c4df7b283310162e43847dd19f2dc7499e3f18bcd57" {
		t.Fatalf("PublicKey got %s, want %s", wallet.PublicKey, "Mp32721d081431dea637f8826fd58babe93ebc9648c76978e26ebf8b56e91292a832fbd62f569fefdc05139c4df7b283310162e43847dd19f2dc7499e3f18bcd57")
	}

	if wallet.Address != "Mx68eb0cb118f8e9666a6d8f5ded4a3eec20fed24b" {
		t.Fatalf("Address got %s, want %s", wallet.Address, "Mx68eb0cb118f8e9666a6d8f5ded4a3eec20fed24b")
	}
}
