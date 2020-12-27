package transaction

import (
	"testing"
)

const txEditCoinOwner = "0xf8710102010111a1e08a535052544553540000009489e5dc185e6bab772ac8e00cf3fb3f4cb0931c47808001b845f8431ba03a08817c5b87aaa4eed6ed0c0a86270bbeb0a9f309453ef996546cbee383e8a0a01d34934dba2d1cb07740156d5af8a9e37b68524fd57e4eca1218d12854350f97"

func TestDecode_editCoinOwner(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txEditCoinOwner)
	if err != nil {
		t.Fatal(err)
	}

	if decode.Fee().String() != "10000000000000000000000" {
		t.Error("edit coin owner transaction fee is invalid", decode.Fee().String())
	}
}
