package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// SetCandidateOnData is a Data of Transaction for turning candidate on.
// This transaction should be sent from ControlAddress or OwnerAddress which is set in the "Declare candidacy transaction".
type SetCandidateOnData struct {
	PubKey PublicKey // Public key of a validator
}

// NewSetCandidateOnData returns new SetCandidateOnData of Transaction for turning candidate on.
func NewSetCandidateOnData() *SetCandidateOnData {
	return &SetCandidateOnData{}
}

// SetPubKey ets public key of a validator.
func (d *SetCandidateOnData) SetPubKey(key string) (*SetCandidateOnData, error) {
	pubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	copy(d.PubKey[:], pubKey)
	return d, nil
}

// MustSetPubKey tries to set public key of validator and panics on error.
func (d *SetCandidateOnData) MustSetPubKey(key string) *SetCandidateOnData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Type returns Data type of the transaction.
func (d *SetCandidateOnData) Type() Type {
	return TypeSetCandidateOnline
}

// Encode returns the byte representation of a transaction Data.
func (d *SetCandidateOnData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
