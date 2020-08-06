package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction data for turning candidate off.
// This transaction should be sent from address which is set in the "Declare candidacy transaction".
type SetCandidateOffData struct {
	PubKey [32]byte // Public key of a validator
}

// New data of Transaction for turning candidate off.
func NewSetCandidateOffData() *SetCandidateOffData {
	return &SetCandidateOffData{}
}

// Set public key of a validator.
func (d *SetCandidateOffData) SetPubKey(key string) (*SetCandidateOffData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// Tries to set public key of validator and panics on error.
func (d *SetCandidateOffData) MustSetPubKey(key string) *SetCandidateOffData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *SetCandidateOffData) Type() Type {
	return TypeSetCandidateOffline
}

func (d *SetCandidateOffData) Fee() Fee {
	return feeTypeSetCandidateOffline
}

func (d *SetCandidateOffData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
