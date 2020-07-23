package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction data for turning candidate on.
// This transaction should be sent from address which is set in the "Declare candidacy transaction".
type SetCandidateOnData struct {
	PubKey [32]byte // Public key of a validator
}

// New data of Transaction for turning candidate on.
func NewSetCandidateOnData() *SetCandidateOnData {
	return &SetCandidateOnData{}
}

// Set public key of a validator.
func (d *SetCandidateOnData) SetPubKey(key string) (*SetCandidateOnData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// Tries to set public key of validator and panics on error.
func (d *SetCandidateOnData) MustSetPubKey(key string) *SetCandidateOnData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *SetCandidateOnData) Type() Type {
	return TypeSetCandidateOnline
}

func (d *SetCandidateOnData) Fee() Fee {
	return feeTypeSetCandidateOnline
}

func (d *SetCandidateOnData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
