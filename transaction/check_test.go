package transaction

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
)

func TestCheck_Sign(t *testing.T) {
	check := NewIssueCheck(
		480,
		TestNetChainID,
		999999,
		"MNT",
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
	).SetPassphrase("pass")

	sign, err := check.Sign("64e27afaab363f21eec05291084367f6f1297a7b280d69d672febecda94a09ea")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := sign.Encode()
	if err != nil {
		t.Fatal(err)
	}
	validCheck := "Mcf8a38334383002830f423f8a4d4e5400000000000000888ac7230489e80000b841d184caa333fe636288fc68d99dea2c8af5f7db4569a0bb91e03214e7e238f89d2b21f4d2b730ef590fd8de72bd43eb5c6265664df5aa3610ef6c71538d9295ee001ba08bd966fc5a093024a243e62cdc8131969152d21ee9220bc0d95044f54e3dd485a033bc4e03da3ea8a2cd2bd149d16c022ee604298575380db8548b4fd6672a9195"
	checkString := string(encode)
	if checkString != validCheck {
		t.Errorf("check want %s,\ngot %s", validCheck, checkString)
	}
}

func TestCheck_Sign1(t *testing.T) {
	check := NewIssueCheck(
		1,
		MainNetChainID,
		999999,
		"MNT",
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
	).SetPassphrase("pass")

	sign, err := check.Sign("2919c43d5c712cae66f869a524d9523999998d51157dc40ac4d8d80a7602ce02")
	if err != nil {
		t.Fatal(err)
	}

	checkString, err := sign.Encode()
	if err != nil {
		t.Fatal(err)
	}

	validCheck := "Mcf8a03101830f423f8a4d4e5400000000000000888ac7230489e80000b84149eba2361855724bbd3d20eb97a54ea15ad7dc28c1111b8dcf3bb15db26f874f095803cad9f8fc88b2b4eec9ba706325a7929be31b6ccfef01260791a844cb55011ba06c63ad17bfe07b82be8a0144fd4daf8b4144281fdf88f313205ceacf37fd877fa03c243ad79cab6205f4b753bd402c4cfa5d570888659090b2f923071ac52bdf75"
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

func TestDecodeIssueCheck(t *testing.T) {
	data, err := DecodeIssueCheck("Mcf8a38334383002830f423f8a4d4e5400000000000000888ac7230489e80000b841d184caa333fe636288fc68d99dea2c8af5f7db4569a0bb91e03214e7e238f89d2b21f4d2b730ef590fd8de72bd43eb5c6265664df5aa3610ef6c71538d9295ee001ba08bd966fc5a093024a243e62cdc8131969152d21ee9220bc0d95044f54e3dd485a033bc4e03da3ea8a2cd2bd149d16c022ee604298575380db8548b4fd6672a9195")
	if err != nil {
		t.Fatal(err)
	}

	if string(data.Nonce) != "480" {
		t.Errorf("Nonce want %s, got %s", string(data.Nonce), "480")
	}
	if data.ChainID != TestNetChainID {
		t.Errorf("ChainID want %d, got %d", data.ChainID, TestNetChainID)
	}
	if string(data.Coin[:3]) != "MNT" {
		t.Errorf("Coin want %s, got %s", string(data.Coin[:3]), "MNT")
	}
	if data.DueBlock != 999999 {
		t.Errorf("DueBlock want %d, got %d", data.DueBlock, 999999)
	}
	if data.Value.String() != "10"+strings.Repeat("0", 18) {
		t.Errorf("Value want %s, got %s", data.Value.String(), "10"+strings.Repeat("0", 18))
	}
	if hex.EncodeToString(data.Lock.Bytes()) != "d184caa333fe636288fc68d99dea2c8af5f7db4569a0bb91e03214e7e238f89d2b21f4d2b730ef590fd8de72bd43eb5c6265664df5aa3610ef6c71538d9295ee00" {
		t.Errorf("Lock want %s, got %s", data.Lock.String(), "d184caa333fe636288fc68d99dea2c8af5f7db4569a0bb91e03214e7e238f89d2b21f4d2b730ef590fd8de72bd43eb5c6265664df5aa3610ef6c71538d9295ee00")
	}
	if data.V.Int64() != 27 {
		t.Errorf("V want %d, got %d", data.V.Int64(), 27)
	}
	if hex.EncodeToString(data.R.Bytes()) != "8bd966fc5a093024a243e62cdc8131969152d21ee9220bc0d95044f54e3dd485" {
		t.Errorf("R want %s, got %s", hex.EncodeToString(data.R.Bytes()), "8bd966fc5a093024a243e62cdc8131969152d21ee9220bc0d95044f54e3dd485")
	}
	if hex.EncodeToString(data.S.Bytes()) != "33bc4e03da3ea8a2cd2bd149d16c022ee604298575380db8548b4fd6672a9195" {
		t.Errorf("S want %s, got %s", hex.EncodeToString(data.S.Bytes()), "33bc4e03da3ea8a2cd2bd149d16c022ee604298575380db8548b4fd6672a9195")
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
