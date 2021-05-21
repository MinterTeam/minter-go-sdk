package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
)

// SetCandidateOffData is a Data of Transaction for turning candidate off.
// This transaction should be sent from ControlAddress or OwnerAddress which is set in the "Declare candidacy transaction".
type SetCandidateOffData struct {
	PubKey PublicKey // Public key of a validator
}

// NewSetCandidateOffData returns new SetCandidateOffData of Transaction for turning candidate off.
func NewSetCandidateOffData() *SetCandidateOffData {
	return &SetCandidateOffData{}
}

// SetPubKey sets public key of a validator.
func (d *SetCandidateOffData) SetPubKey(key string) (*SetCandidateOffData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key of validator and panics on error.
func (d *SetCandidateOffData) MustSetPubKey(key string) *SetCandidateOffData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Type returns Data type of the transaction.
func (d *SetCandidateOffData) Type() Type {
	return TypeSetCandidateOffline
}

// Encode returns the byte representation of a transaction Data.
func (d *SetCandidateOffData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
