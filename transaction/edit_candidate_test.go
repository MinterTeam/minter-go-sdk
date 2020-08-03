package transaction

import (
	"testing"
)

func TestTransactionEditCandidate_Sign(t *testing.T) {
	data := NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47").
		MustSetControlAddress("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

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

	validSignature := "0xf8b4010201010eb863f861a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8809489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0941b685a7c1e78726c48f619c497a07ed75fe00483808001b845f8431ba0f24ee252163f0ffa243ea47031ae34dc4ee9d94e063ae0be201f3b88b49fbedda056006f085890d2d37ec214883f0c0c62f2ce6ed3e354f42655e63d3745d6e072"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}

func TestTransactionEditCandidate_withNewPublicKey_Sign(t *testing.T) {
	data := NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetNewPubKey("Mp0eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a43").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47").
		MustSetControlAddress("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

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

	validSignature := "0xf8d4010201010eb883f881a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8a00eb98ea04ae466d8d38f490db3c99b3996a90e24243952ce9822c6dc1e2c1a439489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0941b685a7c1e78726c48f619c497a07ed75fe00483808001b845f8431ba01b544e2bf840bca9ec046de93166476aef177dbee2541c6831722128b0b2e0f4a0794b7fd0a732667cdb19b2d61817b08671d574eb1a220dddb08b18f2543f02fb"
	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != validSignature {
		t.Errorf("EncodeTx got %s, want %s", encode, validSignature)
	}
}
