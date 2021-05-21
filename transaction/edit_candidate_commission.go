package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/v2/internal/rlp"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
)

type EditCandidateCommissionData struct {
	PubKey     PublicKey
	Commission uint32
}

// SetCommission sets commission of a validator.
func (d *EditCandidateCommissionData) SetCommission(commission uint32) *EditCandidateCommissionData {
	d.Commission = commission
	return d
}

// SetPubKey sets public key of a validator.
func (d *EditCandidateCommissionData) SetPubKey(key string) (*EditCandidateCommissionData, error) {
	pk, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey PublicKey
	copy(pubKey[:], pk)
	d.PubKey = pubKey
	return d, nil
}

// MustSetPubKey tries to set public key of validator and panics on error.
func (d *EditCandidateCommissionData) MustSetPubKey(key string) *EditCandidateCommissionData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

// Type returns Data type of the transaction.
func (d *EditCandidateCommissionData) Type() Type {
	return TypeEditCommissionCandidate
}

// Encode returns the byte representation of a transaction Data.
func (d *EditCandidateCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
