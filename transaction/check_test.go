package transaction

import (
	"math/big"
	"testing"
)

func TestCheck_Sign(t *testing.T) {
	check := NewCheck(
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
		t.Errorf("check want %s, got %s", validCheck, checkString)
	}
}
