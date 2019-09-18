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

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin("MNT")

	signedTx, err := transaction.Sign("a3fb55450f53dbbf4f2494280188f7f0cd51a7b51ec27ed49ed364d920e326ba")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf8a80102018a4d4e54000000000000000eb84df84ba04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b89489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0808001b845f8431ca0421470f27f78231b669c1bf1fcc56168954d64fbb7dc3ff021bab01311fab6eaa075e86365d98c87e806fcbc5c542792f569e19d8ae7af671d9ba4679acc86d35e"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
