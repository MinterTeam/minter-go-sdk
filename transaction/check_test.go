package transaction

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
)

func TestCheck_Sign(t *testing.T) {
	check := NewCheck(
		480,
		TestNetChainID,
		999999,
		1,
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		1,
	).SetPassphrase("pass")

	sign, err := check.Sign("64e27afaab363f21eec05291084367f6f1297a7b280d69d672febecda94a09ea")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := sign.Encode()
	if err != nil {
		t.Fatal(err)
	}
	validCheck := "Mcf8ae8334383002830f423f8a4d4e5400000000000000888ac7230489e800008a4d4e5400000000000000b841497c5f3e6fc182fd1a791522a9ef7576710bdfbc86fdbf165476ef220e89f9ff1380f93f2d9a2f92fdab0edc1e2605cc2c69b707cd404b2cb1522b7aba4defd5001ba083c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594a047b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df"
	checkString := encode
	if checkString != validCheck {
		t.Errorf("check want %s,\ngot %s", validCheck, checkString)
	}
}

func TestCheck_Sign1(t *testing.T) {
	check := NewCheck(
		1,
		MainNetChainID,
		999999,
		1,
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		1,
	).SetPassphrase("pass")

	sign, err := check.Sign("2919c43d5c712cae66f869a524d9523999998d51157dc40ac4d8d80a7602ce02")
	if err != nil {
		t.Fatal(err)
	}

	checkString, err := sign.Encode()
	if err != nil {
		t.Fatal(err)
	}

	validCheck := "Mcf8ab3101830f423f8a4d4e5400000000000000888ac7230489e800008a4d4e5400000000000000b841f69950a210196529f47df938f7af84958cdb336daf304616c37ef8bebca324910910f046e2ff999a7f2ab564bd690c1102ab65a20e0f27b57a93854339b60837011ba00a07cbf311148a6b62c1d1b34a5e0c2b6931a0547ede8b9dfb37aedff4480622a023ac93f7173ca41499624f06dfdd58c4e65d1279ea526777c194ddb623d57027"
	if checkString != validCheck {
		t.Errorf("check want %s,\ngot %s", validCheck, checkString)
	}
}

func TestCheckAddress_Proof(t *testing.T) {
	check, err := NewCheckAddress("Mxa7bc33954f1ce855ed1a8c768fdd32ed927def47", "pass")
	if err != nil {
		t.Fatal(err)
	}

	proofString, err := check.Proof()
	if err != nil {
		t.Fatal(err)
	}

	validProof := "da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600"
	if proofString != validProof {
		t.Errorf("check want %s,\ngot %s", validProof, proofString)
	}
}

func TestDecodeCheck(t *testing.T) {
	data, err := DecodeCheck("+K6DNDgwAoMPQj+KTU5UAAAAAAAAAIiKxyMEiegAAIpNTlQAAAAAAAAAuEFJfF8+b8GC/Rp5FSKp73V2cQvfvIb9vxZUdu8iDon5/xOA+T8tmi+S/asO3B4mBcwsabcHzUBLLLFSK3q6Te/VABugg8mUUWnwp7vllpc7MtyIdgh4BYCx07x7GIvts704VZSgR7LVNFlG7VSY9b7nE/hidqrARqX++CC+rud6m2+bwd8=")
	if err != nil {
		t.Fatal(err)
	}

	if string(data.Nonce) != "480" {
		t.Errorf("Nonce want %s, got %s", string(data.Nonce), "480")
	}
	if data.ChainID != TestNetChainID {
		t.Errorf("ChainID want %d, got %d", data.ChainID, TestNetChainID)
	}
	if data.Coin != 1 {
		t.Errorf("CoinID want %s, got %d", data.Coin.String(), 1)
	}
	if data.DueBlock != 999999 {
		t.Errorf("DueBlock want %d, got %d", data.DueBlock, 999999)
	}
	if data.Value.String() != "10"+strings.Repeat("0", 18) {
		t.Errorf("Value want %s, got %s", data.Value.String(), "10"+strings.Repeat("0", 18))
	}
	if data.Lock.String() != "985283871505876742053353384809055983203325304659217336382177168476233943196356552072809135181493031263037291805582875144952622907359402361966594748392133888" {
		t.Errorf("Lock want %s, got %s", data.Lock.String(), "985283871505876742053353384809055983203325304659217336382177168476233943196356552072809135181493031263037291805582875144952622907359402361966594748392133888")
	}
	if data.V.Int64() != 27 {
		t.Errorf("V want %d, got %d", data.V.Int64(), 28)
	}
	if hex.EncodeToString(data.R.Bytes()) != "83c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594" {
		t.Errorf("R want %s, got %s", hex.EncodeToString(data.R.Bytes()), "83c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594")
	}
	if hex.EncodeToString(data.S.Bytes()) != "47b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df" {
		t.Errorf("S want %s, got %s", hex.EncodeToString(data.S.Bytes()), "47b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df")
	}

	sender, err := data.Sender()
	if err != nil {
		t.Fatal(err)
	}

	address := "Mxce931863b9c94a526d94acd8090c1c5955a6eb4b"
	if sender != address {
		t.Errorf("Sender want %s, got %s", address, sender)
	}
}

func TestDecodeCheck_Sender(t *testing.T) {
	data, err := DecodeCheck("+KwxAoQ7msn/ik1OVAAAAAAAAACIiscjBInoAACKTU5UAAAAAAAAALhBYQJc9o+FL0Z6Az0c4WZu4xi3s1E2qao2l3cGnDOE62o++8TmJywtVrsNJ5fRXOYbYPV15CASqIUImRGp+3rq/gEcoLrEN0b0ePUvPXLb7stC1FL5q3WbbBT1jnb/9AvfeOg3oD0R2ZYtSHDdF8aJs5/wOgyR7Bh5KomUsBWjkoWYMQBC")
	if err != nil {
		t.Fatal(err)
	}

	if data.Coin != 1 {
		t.Errorf("CoinID want %s, got %d", data.Coin.String(), 1)
	}

	sender, err := data.Sender()
	if err != nil {
		t.Fatal(err)
	}

	address := "Mx2f574419b6cba6d886341b5cd4110d2b02eafe8e"
	if sender != address {
		t.Errorf("Sender want %s, got %s", address, sender)
	}
}
