package transaction

import (
	"testing"
)

const txEditCandidate = "0xf8b3010201010eb862f860a04ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b89489e5dc185e6bab772ac8e00cf3fb3f4cb0931c4794e731fcddd37bb6e72286597d22516c8ba3ddffa0941b685a7c1e78726c48f619c497a07ed75fe00483808001b845f8431ca0e88140aadd6cdc38d5ff59e2be43d1d7dfe118b85faa435e04f27d29e3e3f7caa014db705d5e6be34931515744a42080be80a60e39a85d9ebbbb2e977b83cf78c6"

func TestTransactionEditCandidate_Sign(t *testing.T) {
	t.Parallel()
	data := NewEditCandidateData().
		MustSetPubKey("Mp4ae1ee73e6136c85b0ca933a9a1347758a334885f10b3238398a67ac2eb153b8").
		MustSetOwnerAddress("Mxe731fcddd37bb6e72286597d22516c8ba3ddffa0").
		MustSetRewardAddress("Mx89e5dc185e6bab772ac8e00cf3fb3f4cb0931c47").
		MustSetControlAddress("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	transaction := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(1)

	signedTx, err := transaction.Sign("a3fb55450f53dbbf4f2494280188f7f0cd51a7b51ec27ed49ed364d920e326ba")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if encode != txEditCandidate {
		t.Errorf("EncodeTx got %s, want %s", encode, txEditCandidate)
	}
}

func TestDecode_editCandidate(t *testing.T) {
	t.Parallel()
	decode, err := Decode(txEditCandidate)
	if err != nil {
		t.Fatal(err)
	}

	if decode.GetTransaction().Type != TypeEditCandidate {
		t.Error("edit candidate transaction type is invalid", decode.GetTransaction().Type)
	}
}
