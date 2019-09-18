package wallet

import (
	"encoding/hex"
	"testing"
)

const (
	validSeed       = "33fa1096997d9b0f47469463710b3a2e91971144265b281dc71f831539a3b8e3413e5969e5ffb4d3c5a37cfa0f964bcc779efe4ae37fceef048175105caad624"
	validPrivateKey = "d3520cc797f12b8a81e805ddf5a5bf8b994e347003ea25c9ccaecb5073f3fef1"
	validAddress    = "Mx48f502a9fc324f2c707edc3a2595e72f00c3190c"

	validPrivateKey2 = "21bdd69f4a6d0db9508bc543e26bd23378518c8c409496f9ef3e015ff17fc005"
	validPublicKey2  = "Mpb03a4a53b02fba4023ed1afd2598fa1e5fae220198109b865218598f9456a3be4a1e577e0ddb8c72c55b3a0e7102eb35c9dccc1dd4e32228a369b0a57e15bddf"
	validAddress2    = "Mx17b1240ba6d45258f836b45ae0c4fc1106f5ce59"
)

func TestSeed(t *testing.T) {
	mnemonic := "suffer draft bacon typical start retire air sniff large biology mail diagram"
	seed, err := Seed(mnemonic)
	if err != nil {
		t.Fatal(err)
	}

	if hex.EncodeToString(seed) != validSeed {
		t.Fatalf("seed got %s, want %s", hex.EncodeToString(seed), validSeed)
	}
}

func TestPrivateKey(t *testing.T) {
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

func TestAddress(t *testing.T) {
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

func TestPublicKeyByPrivateKey(t *testing.T) {
	pubKey, err := PublicKeyByPrivateKey(validPrivateKey2)
	if err != nil {
		t.Fatal(err)
	}

	if pubKey != validPublicKey2 {
		t.Fatalf("PublicKey got %s, want %s", pubKey, validPublicKey2)
	}
}

func TestAddressByPublicKey(t *testing.T) {

}
