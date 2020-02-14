package wallet

import (
	"encoding/hex"
	"testing"
)

func TestMnemonicBySeed(t *testing.T) {
	mnemonic := "suffer draft bacon typical start retire air sniff large biology mail diagram"
	seed, err := Seed(mnemonic)
	if err != nil {
		t.Fatal(err)
	}

	if hex.EncodeToString(seed) != validSeed {
		t.Fatalf("seed got %s, want %s", hex.EncodeToString(seed), validSeed)
	}
}

func TestPrivateKeyBySeed(t *testing.T) {
	bytes, err := hex.DecodeString(validSeed)
	if err != nil {
		t.Fatal(err)
	}

	seed := bytes

	prKey, err := PrivateKeyBySeed(seed)
	if err != nil {
		t.Fatal(err)
	}

	if prKey != validPrivateKey {
		t.Fatalf("PublicKey got %s, want %s", prKey, validPrivateKey)
	}
}

func TestPublicKeyByPrivateKey(t *testing.T) {
	pubKey, err := PublicKeyByPrivateKey(validPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	if pubKey != validPublicKey {
		t.Fatalf("PublicKey got %s, want %s", pubKey, validPublicKey)
	}
}

func TestAddressByPublicKey(t *testing.T) {
	address, err := AddressByPublicKey(validPublicKey)
	if err != nil {
		t.Fatal(err)
	}

	if address != validAddress {
		t.Fatalf("Address got %s, want %s", address, validAddress)
	}
}

func TestBugAddress(t *testing.T) {
	seed, err := Seed("real town addict extend shoot name disagree vital turn live can tip")
	if err != nil {
		panic(err)
	}
	if hex.EncodeToString(seed) != "57fb1e450b8afb95c62afbcd49e4100d6790e0822b8905608679180ac34ca0bd45bf7ccc6c5f5218236d0eb93afc78bd117b9f02a6b7df258ea182dfaef5aad7" {
		t.Fatalf("Seed got %s, want %s", hex.EncodeToString(seed), "57fb1e450b8afb95c62afbcd49e4100d6790e0822b8905608679180ac34ca0bd45bf7ccc6c5f5218236d0eb93afc78bd117b9f02a6b7df258ea182dfaef5aad7")
	}
	privateKeyBySeed, err := PrivateKeyBySeed(seed)
	if err != nil {
		t.Fatal(err)
	}
	if privateKeyBySeed != "11c332a5107bdede093dc057b146fbb633cbe0f2c50958c9de3aee13fe6caad3" {
		t.Fatalf("PrivateKey got %s, want %s", privateKeyBySeed, "11c332a5107bdede093dc057b146fbb633cbe0f2c50958c9de3aee13fe6caad3")
	}

	publicKeyByPrivateKey, err := PublicKeyByPrivateKey(privateKeyBySeed)
	if err != nil {
		t.Fatal(err)
	}
	if publicKeyByPrivateKey != "Mp32721d081431dea637f8826fd58babe93ebc9648c76978e26ebf8b56e91292a832fbd62f569fefdc05139c4df7b283310162e43847dd19f2dc7499e3f18bcd57" {
		t.Fatalf("PublicKey got %s, want %s", publicKeyByPrivateKey, "Mp32721d081431dea637f8826fd58babe93ebc9648c76978e26ebf8b56e91292a832fbd62f569fefdc05139c4df7b283310162e43847dd19f2dc7499e3f18bcd57")
	}

	addressByPublicKey, err := AddressByPublicKey(publicKeyByPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	if addressByPublicKey != "Mx68eb0cb118f8e9666a6d8f5ded4a3eec20fed24b" {
		t.Fatalf("Address got %s, want %s", addressByPublicKey, "Mx68eb0cb118f8e9666a6d8f5ded4a3eec20fed24b")
	}
}
