package transaction

import (
	"testing"
)

func TestTransactionEditCandidate_Sign(t *testing.T) {
	data := NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)

	signedTx, err := transaction.Sign("a3fb55450f53dbbf4f2494280188f7f0cd51a7b51ec27ed49ed364d920e326ba")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8b4010201010eb863f861a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8809489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0940000000000000000000000000000000000000000808001b845f8431ba06a61c52106973267e4a156c535881796f5d38e51e44cf86aed1a4fe8c9ef1bf9a046eefdebc170e4c81533bd5ead710d8c4d596fd18bd5038e48fd56fb3984e333"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
