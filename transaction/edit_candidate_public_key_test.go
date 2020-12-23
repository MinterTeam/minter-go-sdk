package transaction

import (
	"testing"
)

const txEditCandidatePublicKey = "0xf8950102010114b844f842a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43808001b845f8431ba0933744d45b17ee431460ce6ce9a707e84f7c43b79bd7eb2067bc35352035e141a07731ab3e52be64a41482409ee3d92f6662e2e6d36ad6789c69fb1d4ec75e766f"

func TestDecode_editCandidatePublicKey(t *testing.T) {
	decode, err := Decode(txEditCandidatePublicKey)
	if err != nil {
		t.Fatal(err)
	}

	if decode.Fee().String() != "10000000000000000000000000" {
		t.Error("edit candidate public key transaction fee is invalid", decode.Fee().String())
	}
}
